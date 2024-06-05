package casbin

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-website/config"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var e *casbin.Enforcer

func Init(modelFile string) {
	var err error
	db := loadDBConfig(config.GetMysql())
	e, err = casbin.NewEnforcer(modelFile, db)
	if err != nil {
		tlog.Errorf("init casbin error: %s", err.Error())
		panic(err)
	}
	tlog.Infof("init casbin success!")
	e.LoadPolicy()
}

func loadDBConfig(db *gorm.DB) *gormadapter.Adapter {
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &entity.CasbinRule{})
	if err != nil {
		tlog.Errorf("load db config error: %s", err.Error())
		panic(err)
	}
	return adapter
}

func Check(sub, dom, obj, act string) bool {
	ok, _ := e.Enforce(sub, dom, obj, act)
	return ok
}

func ModifyPolicy(ptype string, sub, dom, obj, act string) error {
	// 修改策略, 开启事务
	if err := e.GetAdapter().(*gormadapter.Adapter).Transaction(e, func(e casbin.IEnforcer) error {
		_, err := e.AddPolicy(sub, dom, obj, act)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func AddPolicy(sub, dom, obj, act string) error {
	// 添加策略, 开启事务
	if err := e.GetAdapter().(*gormadapter.Adapter).Transaction(e, func(e casbin.IEnforcer) error {
		_, err := e.AddPolicy(sub, dom, obj, act)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
