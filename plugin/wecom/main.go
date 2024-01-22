package wecom

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wecom/ComGlobal"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wecom/Router"
)

//go:generate go get -u github.com/ArtisanCloud/PowerWeChat/v3
type PluginWeCom struct {
}

func CreateWeComPlug(CorpID, Secret, ContactSecret string, AgentID int) *PluginWeCom {
	ComGlobal.GlobalConfig.CorpID = CorpID
	ComGlobal.GlobalConfig.AgentID = AgentID
	ComGlobal.GlobalConfig.Secret = Secret
	ComGlobal.GlobalConfig.ContactSecret = ContactSecret

	WeComApp, err := work.NewWork(&work.UserConfig{
		CorpID:      ComGlobal.GlobalConfig.CorpID,  // 企业微信的app id，所有企业微信共用一个。
		AgentID:     ComGlobal.GlobalConfig.AgentID, // 内部应用的app id
		Secret:      ComGlobal.GlobalConfig.Secret,  // 内部应用的app secret
		Token:       "ayX7Im",                       // 企业微信后台，开发与服务，企业微信开发，配置项，token
		AESKey:      "EWRf9UO5B4bNUTHVIEAVzsCNwGGAsJj1DEYnn8dsUYv",
		CallbackURL: "https://moonwife.top/system_api/wecom/callback",
		HttpDebug:   true,
		OAuth: work.OAuth{
			Callback: "https://moonwife.top/system_api/wecom/callback", //
			Scopes:   nil,
		},
		// 可选，不传默认走程序内存
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
	})

	if err != nil {
		panic(err)
	}

	WeComContactApp, err := work.NewWork(&work.UserConfig{
		CorpID:      ComGlobal.GlobalConfig.CorpID,        // 企业微信的app id，所有企业微信共用一个。
		AgentID:     ComGlobal.GlobalConfig.AgentID,       // 内部应用的app id
		Secret:      ComGlobal.GlobalConfig.ContactSecret, // 内部应用的app secret
		Token:       "ayX7Im",                             // 企业微信后台，开发与服务，企业微信开发，配置项，token
		AESKey:      "EWRf9UO5B4bNUTHVIEAVzsCNwGGAsJj1DEYnn8dsUYv",
		CallbackURL: "https://moonwife.top/system_api/wecom/callback",
		HttpDebug:   true,
		OAuth: work.OAuth{
			Callback: "https://moonwife.top/system_api/wecom/callback", //
			Scopes:   nil,
		},
		// 可选，不传默认走程序内存
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
	})

	if err != nil {
		panic(err)
	}

	ComGlobal.GlobalConfig.WeComApp = WeComApp
	ComGlobal.GlobalConfig.WeComContactApp = WeComContactApp

	return &PluginWeCom{}
}

func (*PluginWeCom) Register(group *gin.RouterGroup) {
	Router.RouterGroupApp.InitWeComRouter(group)
	Router.RouterGroupApp.InitDepartmentRouter(group)
}

func (*PluginWeCom) RouterPath() string {
	return "wecom"
}
