package config

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
)

type WeComConfig struct {
	AppID              string // 公众号AppID
	Secret             string // 公众号Secret
	Token              string // 公众号Token
	AESKey             string // 公众号AESKey
	OfficialAccountApp *officialAccount.OfficialAccount
	NullCtx            context.Context
	AuthorityID        uint   // 权限ID
	AccessToken        string // AccessToken
}
