package database

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Cache is an interface
type Database interface {
	Get() *gorm.DB
}

type impl struct {
	mysql_ *gorm.DB
}

func New() Database {
	// 应该由 go lib 提供统一的 New 方法，用于初始化 Redis
	return &impl{
		mysql_: initDatabase(),
	}
}

func initDatabase() *gorm.DB {
	// 初始化数据库连接
	var err error
	dm := config.Get().Database.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dm.Username, dm.Password, dm.Host, dm.Port, dm.Database)
	mysql_, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: dsn,
		},
	), &gorm.Config{
		DisableAutomaticPing: true,
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,          // Don't include params in the SQL log
				Colorful:                  false,         // Disable color
			}),
	})
	if err != nil {
		tlog.Errorf("connect mysql error: %s", err.Error())
		panic(err)
	}
	tlog.Infof("connect mysql success!")
	return mysql_
}

func (i *impl) Get() *gorm.DB {
	return i.mysql_
}
