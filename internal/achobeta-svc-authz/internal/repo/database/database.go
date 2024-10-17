package database

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"context"
	"database/sql"

	"gorm.io/gorm"
)

// Database Cache is an interface
type Database interface {
	Get() *gorm.DB
	Create(ctx context.Context, model any) (*sql.Rows, error)
	Update(ctx context.Context, model any) (*sql.Rows, error)
	Delete(ctx context.Context, model any) (*sql.Rows, error)
	Transaction(ctx context.Context, fn func(trc *gorm.DB) error) error
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

func (i *impl) Create(ctx context.Context, model any) (*sql.Rows, error) {
	return i.mysql_.Debug().WithContext(ctx).Create(model).Rows()
}

func (i *impl) Update(ctx context.Context, model any) (*sql.Rows, error) {
	return i.mysql_.Debug().WithContext(ctx).Updates(model).Rows()
}

func (i *impl) Delete(ctx context.Context, model any) (*sql.Rows, error) {
	return i.mysql_.Debug().WithContext(ctx).Delete(model).Rows()
}
func (i *impl) Transaction(ctx context.Context, fn func(trc *gorm.DB) error) error {
	return i.mysql_.Debug().WithContext(ctx).Transaction(fn)
}
