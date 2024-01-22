package service

import (
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/templateMessage/request"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/wx_request"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
	"sync"
)

type WxMsgService struct{}

type TextMsg struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// SendMsg 发送文本消息
func (s *WxMsgService) SendMsg(msg model.WxMsg) (err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", Wxglobal.GlobalConfig.AccessToken)
	method := "POST"

	payload := fmt.Sprintf(`{
		"touser": "%s",
		"msgtype": "text",
		"text": {
			"content": "%s"
		}
	}`, msg.ToUser, msg.Content)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	response, err := utils.SendRequest(url, method, payload, headers)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

// SendWarningTemplateMsg 发送预警模板消息
func (s *WxMsgService) SendWarningTemplateMsg(msg wx_request.TemplateMessageRequest) (response wx_request.TemplateMessageResponse, err error) {
	var wg sync.WaitGroup
	ch := make(chan string)
	for _, v := range msg.ToUser {
		wg.Add(1)
		openid := v
		go func() {
			defer wg.Done()
			response, err := Wxglobal.GlobalConfig.OfficialAccountApp.TemplateMessage.Send(Wxglobal.GlobalConfig.NullCtx, &request.RequestTemlateMessage{
				ToUser:     openid,
				TemplateID: "gh8651al9njKciGjROFpvmzifvO_XJ-i7A2lHk5gMQU",
				URL:        msg.Url,
				Data: &power.HashMap{
					"short_thing1": &power.HashMap{
						"value": msg.Type,
						"color": "#173177",
					},
					"short_thing2": &power.HashMap{
						"value": msg.Name,
						"color": "#173177",
					},
					"time3": &power.HashMap{
						"value": msg.Time,
						"color": "#173177",
					},
					"thing5": &power.HashMap{
						"value": msg.Reason,
						"color": "#173177",
					},
					"thing7": &power.HashMap{
						"value": msg.ProjectName,
						"color": "#173177",
					},
				},
			})
			if err != nil {
				return
			}
			if response.ErrCode != 0 {
				return
			}
			ch <- openid
		}()
	}

	// 使用另一个goroutine来等待所有任务完成
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 计算成功数量和失败数量
	for result := range ch {
		response.SuccessCount++
		response.FailedCount = len(msg.ToUser) - response.SuccessCount
		response.OkSendOpenId = append(response.OkSendOpenId, result)
	}
	return
}

// SendCommonTemplateMsg 发送通用模板消息
func (s *WxMsgService) SendCommonTemplateMsg(msg wx_request.TemplateCommonMessageRequest) (response wx_request.TemplateMessageResponse, err error) {
	var wg sync.WaitGroup
	ch := make(chan string)
	for _, v := range msg.ToUser {
		wg.Add(1)
		openid := v
		go func() {
			defer wg.Done()
			response, err := Wxglobal.GlobalConfig.OfficialAccountApp.TemplateMessage.Send(Wxglobal.GlobalConfig.NullCtx, &request.RequestTemlateMessage{
				ToUser:     openid,
				TemplateID: msg.TemplateId,
				URL:        msg.Url,
				Data:       msg.Data,
			})
			if err != nil {
				return
			}
			if response.ErrCode != 0 {
				return
			}
			ch <- openid
		}()
	}

	// 使用另一个goroutine来等待所有任务完成
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 计算成功数量和失败数量
	for result := range ch {
		response.SuccessCount++
		response.FailedCount = len(msg.ToUser) - response.SuccessCount
		response.OkSendOpenId = append(response.OkSendOpenId, result)
	}
	return
}
