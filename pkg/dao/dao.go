package dao

import (
	"blackhole-blog/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

var (
	User    = userDao{}
	Article = articleDao{}
)

func Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.Config.Database.User,
		setting.Config.Database.Password,
		setting.Config.Database.Host,
		setting.Config.Database.Port,
		setting.Config.Database.DBName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(getLogMode(setting.Config.Database.LogMode)),
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "bh_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
}

func getLogMode(str string) logger.LogLevel {
	switch str {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
