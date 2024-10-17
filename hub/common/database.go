package common

import (
	"fmt"
	"go-web-mini/config"
	"go-web-mini/model"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 全局mysql数据库变量
var DB *gorm.DB

// 初始化mysql数据库
func InitDB() {
	var dialector gorm.Dialector
	switch config.Conf.Database.Type {
	case "sqlite":
		dialector = sqlite.Open(config.Conf.Database.Sqlite.Path)
	case "postgres":
		dbString := fmt.Sprintf(
			"host=%s dbname=%s user=%s",
			config.Conf.Database.Postgres.Host,
			config.Conf.Database.Postgres.Name,
			config.Conf.Database.Postgres.User,
		)
		if sslEnabled, err := strconv.ParseBool(config.Conf.Database.Postgres.Ssl); err == nil {
			if !sslEnabled {
				dbString += " sslmode=disable"
			}
		} else {
			dbString += fmt.Sprintf(" sslmode=%s", config.Conf.Database.Postgres.Ssl)
		}

		if config.Conf.Database.Postgres.Port != 0 {
			dbString += fmt.Sprintf(" port=%d", config.Conf.Database.Postgres.Port)
		}

		if config.Conf.Database.Postgres.Pass != "" {
			dbString += fmt.Sprintf(" password=%s", config.Conf.Database.Postgres.Pass)
		}

		dialector = postgres.Open(dbString)
	default:
		panic("Database driver type error")
	}

	gConfig := gorm.Config{}
	if config.Conf.Database.Gorm.TablePrefix != "" {
		gConfig.NamingStrategy = schema.NamingStrategy{
			TablePrefix: config.Conf.Database.Gorm.TablePrefix + "_",
		}
	}

	//Log.Info("数据库连接DSN: ", showDsn)
	db, err := gorm.Open(dialector, &gConfig)
	if err != nil {
		Log.Panicf("初始化数据库异常: %v", err)
		panic(fmt.Errorf("初始化数据库异常: %v", err))
	}

	// // 开启日志
	if config.Conf.Database.Debug {
		db.Debug()
	}
	// // 全局DB赋值
	DB = db
	// 自动迁移表结构
	if err = dbAutoMigrate(); err != nil {
		panic(fmt.Errorf("初始化数据库异常: %v", err))
	}
	Log.Infof("初始化数据库完成!")
}

// 自动迁移表结构
func dbAutoMigrate() error {
	return DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Menu{},
		&model.Api{},
		&model.OperationLog{},
	)
}
