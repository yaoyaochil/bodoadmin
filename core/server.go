package core

import (
	"fmt"
	"time"

	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/initialize"
	"github.com/yaoyaochil/bodo-admin-server/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.BODO_CONFIG.System.UseMultipoint || global.BODO_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.BODO_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.BODO_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.BODO_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	波动科技，未来会更好
	药要吃私人播客:https://moonwife.top
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.BODO_LOG.Error(s.ListenAndServe().Error())
}
