package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/api/v1/system"
	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/service"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
	"go.uber.org/zap"
)

type WxLoginApi struct{}

// MsgService 微信工具
var MsgService = service.ServiceGroupApp.WxMsgService

// GetLoginPic 获取登录二维码
func (p *WxLoginApi) GetLoginPic(c *gin.Context) {
	loginFlag := c.Query("loginFlag")
	pic, err := service.ServiceGroupApp.GetLoginPic(loginFlag)
	if err != nil {
		global.BODO_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(pic, c)
	}
}

type TextMessage struct {
	Msg    string
	ToUser string
}

// LoginOrCreate 登录或者创建用户
func (p *WxLoginApi) LoginOrCreate(c *gin.Context) {
	loginFlag := c.Query("loginFlag")
	var ID uint
	// 从map中获取登录标识
	mapInfo, ok := Wxglobal.LoginCodeMap.Load(loginFlag)
	if !ok {
		response.FailWithMessage("请重新获取二维码", c)
		return
	}
	fmt.Println(mapInfo)
	if mapInfo.(Wxglobal.CodeInfo).CanLogin {
		user, err := Wxglobal.GlobalConfig.OfficialAccountApp.User.Get(Wxglobal.GlobalConfig.NullCtx, mapInfo.(Wxglobal.CodeInfo).OpenId, "zh_CN")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		token, _ := c.GetQuery("state")
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, _ := j.ParseToken(token)
		if claims == nil {
			ID = 0
		} else {
			ID = claims.ID
		}

		// 转换数据
		var userInfo model.WXUserInfo
		userInfo = model.WXUserInfo{
			BODOUserId:     ID,
			Subscribe:      user.Subscribe,
			Openid:         user.OpenID,
			Language:       user.Language,
			SubscribeTime:  user.SubscribeTime,
			Unionid:        user.UnionID,
			Remark:         user.Remark,
			Groupid:        user.GroupID,
			SubscribeScene: user.SubscribeScene,
			QrScene:        user.QrScene,
			QrSceneStr:     user.QrSceneStr,
		}

		//

		if resUser, err := service.ServiceGroupApp.LoginOrCreate(userInfo, ID); err != nil {
			global.BODO_LOG.Error("操作失败！", zap.Error(err))
			response.FailWithMessage("操作失败！", c)
			return
		} else {
			if ID != 0 {
				response.OkWithDetailed(gin.H{"scan": true}, "绑定成功", c)
				err := MsgService.SendMsg(model.WxMsg{
					ToUser:  userInfo.Openid,
					Content: "账号" + claims.Username + "成功绑定",
				})
				if err != nil {
					return
				}
			} else {
				fmt.Println(claims)
				err := MsgService.SendMsg(model.WxMsg{
					ToUser:  userInfo.Openid,
					Content: "账号 " + resUser.Username + " 登陆成功！",
				})
				if err != nil {
					return
				}
				var baseApi = new(system.BaseApi)
				baseApi.TokenNext(c, *resUser)
			}
			return
		}
	} else {
		response.OkWithDetailed(gin.H{"scan": false}, "未扫码", c)
		return
	}
}

// ClearWx 清除微信绑定
func (p *WxLoginApi) ClearWx(c *gin.Context) {
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	if err := service.ServiceGroupApp.ClearWx(id); err != nil {
		global.BODO_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("解绑成功", c)
	}
}

// CheckBind 检查是否绑定
func (p *WxLoginApi) CheckBind(c *gin.Context) {
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	isBind, err := service.ServiceGroupApp.CheckBindWx(id)
	if err != nil {
		global.BODO_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(isBind, c)
	}
}
