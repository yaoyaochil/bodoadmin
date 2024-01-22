package seazen

import (
	"github.com/gin-gonic/gin"
	BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/router"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/service"
	"go.uber.org/zap"
)

type PluginSeazen struct {
}

var job = MockJob{}

type MockJob struct{}

func (job MockJob) Run() {
	err := service.ServiceGroupApp.WarnMsgSendService.SendWarningTemplateMsg()
	if err != nil {
		BODO_Global.BODO_LOG.Error("数据推送失败!", zap.Error(err))
	}
}

func CreateSeazenPlug() *PluginSeazen {
	err := BODO_Global.BODO_DB.AutoMigrate(
		&model.SendMsgUserList{},
		&model.QrCodeList{},
		&model.SendMsgType{},
		&model.SeazenConfig{},
		&model.WarnMsgInfoList{},
	)
	if err != nil {
		panic(err)
	}

	var seazenConfig model.SeazenConfig
	err = BODO_Global.BODO_DB.First(&seazenConfig).Error
	if err != nil {
		panic(err)
	}
	global.GlobalConfig.Cookie = seazenConfig.Cookie
	// 每天凌晨1点执行 DoDo身份组同步 DoDo成员同步 DoDo成员对应身份组同步
	_, err = BODO_Global.BODO_Timer.AddTaskByJob("sendWarnMsg", "0 0 9 * * *", job)
	if err != nil {
		BODO_Global.BODO_LOG.Error("定时任务注册失败", zap.Error(err))
	}
	return &PluginSeazen{}
}

func (*PluginSeazen) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitSeazenRouter(group)
	router.RouterGroupApp.InitGagRouter(group)
	router.RouterGroupApp.InitMsgRouter(group)
}

func (*PluginSeazen) RouterPath() string {
	return "seazen"
}
