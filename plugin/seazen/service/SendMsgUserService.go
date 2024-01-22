package service

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/util"
	models3 "github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/service"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/wx_request"
)

type SendMsgUserService struct {
}

var WechatService = service.ServiceGroupApp.WxMsgService

func (s SendMsgUserService) AddUser(msg models3.WXTextMsg) (err error) {
	if msg.EventKey == "gag" {
		var atuser model.SendMsgUserList
		err = BODO_Global.BODO_DB.Where("open_id = ?", msg.FromUserName).First(&atuser).Error
		if err == nil {
			_, err = WechatService.SendCommonTemplateMsg(wx_request.TemplateCommonMessageRequest{
				ToUser: []string{msg.FromUserName},
				// TemplateId: "Wv_uKi0pWiggoIjIvw4YNOWxyDwTu-hERLiNmos9b9w",
				TemplateId: "cMbBKTWUDNy04n8mW3F9HTqzNlC68ye9vpHLE5yMgKo",
				Url:        "",
				Data: &power.HashMap{
					"thing2": &power.HashMap{
						"value": "您已订阅过该服务！",
						"color": "#173177",
					}, // 事项
					"time3": &power.HashMap{
						"value": util.GetToday(),
						"color": "#173177",
					},
					"thing4": &power.HashMap{
						"value": "销采报警模块",
						"color": "#173177",
					},
				},
			})
			if err != nil {
				return err
			}
			return
		}

		user := model.SendMsgUserList{
			OpenID:    msg.FromUserName,
			IsEnabled: false,
		}
		err = BODO_Global.BODO_DB.Create(&user).Error
		if err != nil {
			_, err = WechatService.SendCommonTemplateMsg(wx_request.TemplateCommonMessageRequest{
				ToUser: []string{msg.FromUserName},
				// TemplateId: "Wv_uKi0pWiggoIjIvw4YNOWxyDwTu-hERLiNmos9b9w",
				TemplateId: "cMbBKTWUDNy04n8mW3F9HTqzNlC68ye9vpHLE5yMgKo",
				Url:        "",
				Data: &power.HashMap{
					"thing2": &power.HashMap{
						"value": "订阅服务失败！",
						"color": "#173177",
					}, // 事项
					"time3": &power.HashMap{
						"value": util.GetToday(),
						"color": "#173177",
					},
					"thing4": &power.HashMap{
						"value": "销采报警模块",
						"color": "#173177",
					},
				},
			})
			if err != nil {
				return err
			}
			return
		}

		_, err = WechatService.SendCommonTemplateMsg(wx_request.TemplateCommonMessageRequest{
			ToUser: []string{msg.FromUserName},
			// TemplateId: "Wv_uKi0pWiggoIjIvw4YNOWxyDwTu-hERLiNmos9b9w",
			TemplateId: "cMbBKTWUDNy04n8mW3F9HTqzNlC68ye9vpHLE5yMgKo",
			Url:        "",
			Data: &power.HashMap{
				"thing2": &power.HashMap{
					"value": "订阅服务成功！",
					"color": "#173177",
				},
				"time3": &power.HashMap{
					"value": util.GetToday(),
					"color": "#173177",
				},
				"thing4": &power.HashMap{
					"value": "销采报警模块",
					"color": "#173177",
				},
			},
		})
		return
	}
	return
}

// GetUserList 获取用户列表
func (s SendMsgUserService) GetUserList() (userList []model.SendMsgUserList, err error) {
	err = BODO_Global.BODO_DB.Where("is_enabled = ?", true).Find(&userList).Error
	return
}

// CheckUser 审核用户
func (s SendMsgUserService) CheckUser(id uint) (err error) {
	var user model.SendMsgUserList
	err = BODO_Global.BODO_DB.Where("id = ?", id).First(&user).Update("is_enabled", true).Error
	if err != nil {
		return
	}
	_, err = WechatService.SendCommonTemplateMsg(wx_request.TemplateCommonMessageRequest{
		ToUser: []string{user.OpenID},
		// TemplateId: "Wv_uKi0pWiggoIjIvw4YNOWxyDwTu-hERLiNmos9b9w",
		TemplateId: "cMbBKTWUDNy04n8mW3F9HTqzNlC68ye9vpHLE5yMgKo",
		Url:        "",
		Data: &power.HashMap{
			"thing2": "审核已通过！", // 事项
			"time3":  util.GetToday(),
			"thing4": "销采报警模块",
		},
	})
	return
}
