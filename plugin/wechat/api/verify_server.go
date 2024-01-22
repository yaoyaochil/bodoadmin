package api

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	models2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/server/handlers/models"
	"github.com/gin-gonic/gin"
	global2 "github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/service"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
	models3 "github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/model"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WechatApi struct{}

var SeazenService = service.ServiceGroupApp.SendMsgUserService

// VerifyServer 验证服务器
func (s *WechatApi) VerifyServer(c *gin.Context) {
	openid := c.Query("openid")
	var err error
	rs, err := Wxglobal.GlobalConfig.OfficialAccountApp.Server.VerifyURL(c.Request)
	if err != nil {
		global2.BODO_LOG.Error("验证服务器错误！")
		return
	}
	_, err = Wxglobal.GlobalConfig.OfficialAccountApp.Server.Notify(c.Request, func(event contract.EventInterface) interface{} {
		msg := models3.WXTextMsg{}
		// 这里需要获取到事件类型，然后把对应的结构体传递进去进一步解析
		// 所有包含的结构体请参考： https://github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/server/handlers/models
		switch event.GetMsgType() {
		case models2.CALLBACK_MSG_TYPE_TEXT:
			msg := models.MessageText{}
			err = event.ReadMessage(&msg)
			if err != nil {
				global2.BODO_LOG.Error("读取消息错误！")
				return err
			}
			break
		}

		switch event.GetEvent() {
		case Wxglobal.SCAN:
			err = event.ReadMessage(&msg)
			Wxglobal.LoginCodeMap.Store(msg.EventKey, Wxglobal.CodeInfo{
				CanLogin:   true,
				OpenId:     openid,
				CreateTime: time.Now().Unix(),
			})
			Wxglobal.BindMap.Store(msg.EventKey, Wxglobal.CodeInfo{
				CanLogin:   true,
				OpenId:     openid,
				CreateTime: time.Now().Unix(),
			})
			// 注册消息处理插件
			err = SeazenService.AddUser(msg)
			if err != nil {
				break
			}
			break
		case Wxglobal.SUBSCRIBE:
			if strings.Index(msg.EventKey, "_") > -1 {
				arg := strings.Split(msg.EventKey, "_")
				if arg[0] == "qrscene" {
					Wxglobal.LoginCodeMap.Store(arg[1], Wxglobal.CodeInfo{
						CanLogin:   true,
						OpenId:     openid,
						CreateTime: time.Now().Unix(),
					})
				}
			}
			break
		}

		// 假设用户给应用发送消息，这里可以直接回复消息文本，

		// 这里回复success告诉微信我收到了，后续需要回复用户信息可以主动调发消息接口
		return kernel.SUCCESS_EMPTY_RESPONSE
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	// 返回给微信的消息
	text, err := ioutil.ReadAll(rs.Body)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	c.String(http.StatusOK, string(text))
}
