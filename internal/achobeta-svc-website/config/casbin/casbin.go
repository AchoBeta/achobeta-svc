package casbin

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"

	"github.com/casbin/casbin/v2"
)

var e *casbin.Enforcer

func Init(modelFile string) {
	var err error
	e, err = casbin.NewEnforcer(modelFile)
	if err != nil {
		tlog.Errorf("init casbin error: %s", err.Error())
		panic(err)
	}
}

func Check(sub, dom, obj, act string) bool {
	ok, _ := e.Enforce(sub, dom, obj, act)
	return ok
}
