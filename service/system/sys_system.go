package system

import (
	"github.com/yaoyaochil/bodo-admin-server/server/config"
	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
	"go.uber.org/zap"
)

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: conf Config.Server, err error

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetSystemConfig() (conf config.Server, err error) {
	return global.BODO_CONFIG, nil
}

// @description   set system Config,
//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system Model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.BODO_VP.Set(k, v)
	}
	err = global.BODO_VP.WriteConfig()
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.BODO_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.BODO_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.BODO_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
