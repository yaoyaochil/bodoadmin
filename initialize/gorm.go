package initialize

import (
	"os"

	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/example"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.BODO_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},

		// 示例模块表
		example.ExaFile{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},

		// 自动化模块表
		// Code generated by github.com/yaoyaochil/bodo-admin-server/server Begin; DO NOT EDIT.

		// Code generated by github.com/yaoyaochil/bodo-admin-server/server End; DO NOT EDIT.
	)
	if err != nil {
		global.BODO_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.BODO_LOG.Info("register table success")
}
