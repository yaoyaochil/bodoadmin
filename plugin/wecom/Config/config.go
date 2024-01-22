package Config

import "github.com/ArtisanCloud/PowerWeChat/v3/src/work"

type WeComConfig struct {
	CorpID          string     // 企业微信CorpID
	AgentID         int        // 企业微信应用AgentID
	Secret          string     // 企业微信应用Secret
	ContactSecret   string     // 企业微信通讯录Secret
	WeComApp        *work.Work // 企业微信实例
	WeComContactApp *work.Work // 企业微信通讯实例
}
