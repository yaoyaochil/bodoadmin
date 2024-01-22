package wechat

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/router"
	"time"
)

//go:generate go get -u github.com/ArtisanCloud/PowerWeChat/v3
type PluginWeCom struct {
}

func CreateMiniPlug(AppID string, Secret string, Token string, AESKey string, AuthorityID uint) *PluginWeCom {
	Wxglobal.GlobalConfig.AppID = AppID
	Wxglobal.GlobalConfig.Secret = Secret
	Wxglobal.GlobalConfig.Token = Token
	Wxglobal.GlobalConfig.AESKey = AESKey
	Wxglobal.GlobalConfig.AuthorityID = AuthorityID
	err := global.BODO_DB.AutoMigrate(&model.WXUserInfo{})
	if err != nil {
		return nil
	}

	// 初始化微信公众号
	OfficialAccountApp, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:  Wxglobal.GlobalConfig.AppID,  // 公众号、小程序的appid
		Secret: Wxglobal.GlobalConfig.Secret, // 公众号、小程序的secret
		Token:  Wxglobal.GlobalConfig.Token,  // 公众号的token
		AESKey: Wxglobal.GlobalConfig.AESKey, // 公众号的EncodingAESKey

		Log: officialAccount.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		//Cache RedisCache
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     global.BODO_CONFIG.Redis.Addr,
			Password: global.BODO_CONFIG.Redis.Password,
			DB:       0,
		}),

		HttpDebug: true,
		Debug:     true,
	})
	if err != nil {
		panic(err)
	}
	// 保存到全局变量
	Wxglobal.GlobalConfig.OfficialAccountApp = OfficialAccountApp
	Wxglobal.GlobalConfig.NullCtx = context.Context(context.Background())

	token, err := Wxglobal.GlobalConfig.OfficialAccountApp.AccessToken.GetToken(true)
	if err != nil {
		return nil
	}
	Wxglobal.GlobalConfig.AccessToken = token.AccessToken
	go func() {
		for {
			time.Sleep(300 * time.Second)
			Wxglobal.LoginCodeMap.Range(func(key, value interface{}) bool {
				if time.Now().Unix()-value.(Wxglobal.CodeInfo).CreateTime > 300 {
					Wxglobal.LoginCodeMap.Delete(key)
				}
				return false
			})
			token, err = Wxglobal.GlobalConfig.OfficialAccountApp.AccessToken.GetToken(true)
			if err != nil {
				return
			}
			Wxglobal.GlobalConfig.AccessToken = token.AccessToken
		}
	}()
	return &PluginWeCom{}
}

func (*PluginWeCom) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitWechatRouter(group)
	router.RouterGroupApp.InitWxLoginRouter(group)
	router.RouterGroupApp.InitSendMsgRouter(group)
	router.RouterGroupApp.InitQrRouter(group)
	router.RouterGroupApp.InitMediaRouter(group)
	router.RouterGroupApp.InitMenuRouter(group)
}

func (*PluginWeCom) RouterPath() string {
	return "wechat"
}
