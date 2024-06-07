package casbin

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"flag"
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	jwt "github.com/dgrijalva/jwt-go"
)

type Casbin interface {
	Check(sub, dom, obj, act string) bool
	ModifyPolicy(ptype string, sub, dom, obj, act string) error
	AddPolicy(sub, dom, obj, act string) error
	CreateToken(sub, dom, obj, act string) (string, error)
	VerifyToken(token string) (jwt.MapClaims, error)
}
type impl struct {
	e *casbin.Enforcer
}

var modelFile = flag.String("casbin", "", "specify casbin model path [casbin_model.conf]")

func New() Casbin {
	return &impl{
		e: initCasbin(),
	}

}
func initCasbin() *casbin.Enforcer {
	flag.Parse()
	var err error
	db := loadDBConfig()
	e, err := casbin.NewEnforcer(*modelFile, db)
	if err != nil {
		tlog.Errorf("init casbin error: %s", err.Error())
		panic(err)
	}
	tlog.Infof("init casbin success!")
	if err = e.LoadPolicy(); err != nil {
		tlog.Errorf("load policy error: %s", err.Error())
		panic(err)
	}
	return e
}

func loadDBConfig() *gormadapter.Adapter {
	dm := config.Get().Database.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dm.Username, dm.Password, dm.Host, dm.Port, dm.Database)
	adapter, err := gormadapter.NewAdapter("mysql", dsn, true)
	// adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &entity.CasbinRule{})
	if err != nil {
		tlog.Errorf("load db config error: %s", err.Error())
		panic(err)
	}
	return adapter
}

func (c *impl) Check(sub, dom, obj, act string) bool {
	ok, _ := c.e.Enforce(sub, dom, obj, act)
	return ok
}

func (c *impl) ModifyPolicy(ptype string, sub, dom, obj, act string) error {
	// 修改策略, 开启事务
	if err := c.e.GetAdapter().(*gormadapter.Adapter).Transaction(c.e, func(e casbin.IEnforcer) error {
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

func (c *impl) AddPolicy(sub, dom, obj, act string) error {
	// 添加策略, 开启事务
	if err := c.e.GetAdapter().(*gormadapter.Adapter).Transaction(c.e, func(e casbin.IEnforcer) error {
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
