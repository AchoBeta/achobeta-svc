package account

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/cache"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/casbin"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/database"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"context"
	"fmt"
)

type Permission struct {
	database database.Database
	cache    cache.Cache
	casbin   casbin.Casbin
}

func New(db database.Database, c cache.Cache, cas casbin.Casbin) *Permission {
	return &Permission{
		database: db,
		cache:    c,
		casbin:   cas,
	}
}

// CreateAccount 创建账号
// 方法内部对密码进行加密, 外层调用无需关心加密逻辑
func (p *Permission) CreateAccount(ctx context.Context, ue *entity.Account) error {
	ue.ID = uint(utils.GetSnowflakeID())
	ue.Password = hashPassword(ue.Password)
	_, err := p.database.Create(&ue)
	if err != nil {
		return err
	}
	tlog.CtxInfof(ctx, "create account, username:[%s], email:[%s], phone:[%s]", ue.Username, ue.Email, ue.Phone)
	return nil
}

func hashPassword(pwd string) string {
	hashedPwd, err := utils.HashPassword(pwd)
	if err != nil {
		tlog.Errorf("hash password error: %v", err)
		return pwd
	}
	return string(hashedPwd)
}

func (p *Permission) QueryAccount(ctx context.Context, params *entity.Account) (*entity.Account, error) {
	account := &entity.Account{}
	tlog.CtxInfof(ctx, "query account, params:[%+v\n]", params)
	result := p.database.Get().Where(params).First(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func (p *Permission) CheckToken(ctx context.Context, token string) (bool, error) {
	claims, err := p.casbin.VerifyToken(token)
	if err != nil {
		tlog.CtxErrorf(ctx, "verify token error: %v", err)
		return false, err
	}
	isVaild := p.casbin.Check(claims["userId"].(string), claims["domain"].(string), claims["object"].(string), claims["action"].(string))
	return isVaild, nil
}

func (p *Permission) Login(ctx context.Context, req *entity.LoginRequest) (string, error) {
	if req.Type == entity.LoginTypeUsername {
		account := &entity.Account{}
		p.database.Get().Where("username = ?", req.LoginKey).First(&account)
		if utils.ComparePasswords(account.Password, req.LoginPwd) {
			cb := &entity.CasbinRule{}
			p.database.Get().Where("ptype = ?", "p").Where("v0 = ?", account.UserId).Find(&cb)
			// ptype, v0(userid), v1(domain), v2(object), v3(action)
			token, err := p.casbin.CreateToken(cb.V0, cb.V1, cb.V2, cb.V3)
			if err != nil {
				tlog.CtxErrorf(ctx, "create token error: %v", err)
				return "", err
			}
			return token, nil
		}
		return "", fmt.Errorf("login failed")
	}
	return "", fmt.Errorf("login type not support")
}
