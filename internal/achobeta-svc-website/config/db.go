package config

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"fmt"
    "github.com/go-redis/redis"

    "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"log"
	"os"
	"time"
)

var (
	mysql_ *gorm.DB
	redis_ *redis.Client
)

func InitDatabase() {
	// 初始化数据库连接
	var err error
	dm := Get().Database.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dm.Username, dm.Password, dm.Host, dm.Port, dm.Database)
	mysql_, err = gorm.Open(mysql.New(
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
}

func InitRedis() {
	rc := Get().Database.Redis
	redis_ =redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",rc.Host,rc.Port),
		Password: rc.Password,
		DB: rc.Database,
		})
	err := redis_.Ping().Err()
	if err != nil {
		tlog.Errorf("connect redis error: %s", err.Error())
		panic(err)
	}
	tlog.Infof("connect redis success!")
}

func GetMysql() *gorm.DB {
	return mysql_
}

func GetRedis() *redis.Client {
	return redis_
}
