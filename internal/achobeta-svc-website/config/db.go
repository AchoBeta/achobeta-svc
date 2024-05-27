package config

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	mysql_ *gorm.DB
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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ab_",
			SingularTable: true,
		},
	})
	if err != nil {
		tlog.Errorf("connect mysql error: %s", err.Error())
		panic(err)
	}
	tlog.Infof("connect mysql success!")
}

func GetMysql() *gorm.DB {
	return mysql_
}
