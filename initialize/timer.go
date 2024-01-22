package initialize

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"github.com/yaoyaochil/bodo-admin-server/server/config"
	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
)

func Timer() {
	if global.BODO_CONFIG.Timer.Start {
		for i := range global.BODO_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.BODO_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.BODO_Timer.AddTaskByFunc("ClearDB", global.BODO_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.BODO_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.BODO_CONFIG.Timer.Detail[i])
		}
	}
}
