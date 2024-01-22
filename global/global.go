package global

import (
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/yaoyaochil/bodo-admin-server/server/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/yaoyaochil/bodo-admin-server/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	BODO_DB     *gorm.DB
	BODO_DBList map[string]*gorm.DB
	BODO_REDIS  *redis.Client
	BODO_CONFIG config.Server
	BODO_VP     *viper.Viper
	// BODO_LOG    *oplogging.Logger
	BODO_LOG                 *zap.Logger
	BODO_Timer               timer.Timer = timer.NewTimerTask()
	BODO_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return BODO_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := BODO_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
