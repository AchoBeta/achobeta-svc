package database

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"database/sql"

	"gorm.io/gorm"
)

// Database Cache is an interface
type Database interface {
	Get() *gorm.DB
	Create(model any) (*sql.Rows, error)
	Update(model any) (*sql.Rows, error)
	Delete(model any) (*sql.Rows, error)
	Transaction(func(trc *gorm.DB) error) error
}

type impl struct {
	mysql_ *gorm.DB
}

func New() Database {
	// 应该由 go lib 提供统一的 New 方法，用于初始化 Redis
	return &impl{
		mysql_: config.GetDB(),
	}
}

func (i *impl) Get() *gorm.DB {
	return i.mysql_
}

func (i *impl) Create(model any) (*sql.Rows, error) {
	return i.mysql_.Debug().Create(model).Rows()
}

func (i *impl) Update(model any) (*sql.Rows, error) {
	return i.mysql_.Debug().Updates(model).Rows()
}

func (i *impl) Delete(model any) (*sql.Rows, error) {
	return i.mysql_.Debug().Delete(model).Rows()
}
func (i *impl) Transaction(fn func(trc *gorm.DB) error) error {
	return i.mysql_.Debug().Transaction(fn)
}
