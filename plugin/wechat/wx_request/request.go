package wx_request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type TemplateMessageRequest struct {
	ToUser      []string `json:"touser"`
	Url         string   `json:"url"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Time        string   `json:"time"`
	Reason      string   `json:"reason"`
	ProjectName string   `json:"project_name"`
}

type TemplateMessageResponse struct {
	SuccessCount int      `json:"success_count"`
	FailedCount  int      `json:"failed_count"`
	OkSendOpenId []string `json:"ok_send_openid"`
}

type TemplateCommonMessageRequest struct {
	ToUser     []string       `json:"touser"`
	TemplateId string         `json:"template_id"`
	Url        string         `json:"url"`
	Data       *power.HashMap `json:"data"`
}
