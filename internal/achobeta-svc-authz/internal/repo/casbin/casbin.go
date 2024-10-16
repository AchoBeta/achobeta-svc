package casbin

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"flag"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	jwt "github.com/dgrijalva/jwt-go"
)

type Casbin interface {
	Check(sub, dom, lvl, obj, act string) bool
	CreateToken(sub, obj, dom string) (string, error)
	AddPolicy(sub *entity.CasbinRule) error
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
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(config.GetDB(), &entity.CasbinRule{})
	if err != nil {
		tlog.Errorf("load db config error: %s", err.Error())
		panic(err)
	}
	return adapter
}

func (c *impl) Check(sub, dom, lvl, obj, act string) bool {
	ok, _ := c.e.Enforce(sub, dom, lvl, obj, act)
	return ok
}

func (c *impl) AddPolicy(sub *entity.CasbinRule) error {
	// 添加策略, 开启事务
	if err := c.e.GetAdapter().(*gormadapter.Adapter).Transaction(c.e, func(e casbin.IEnforcer) error {
		_, err := e.AddRoleForUser(sub.V0, sub.V1, sub.V2, sub.V3)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
