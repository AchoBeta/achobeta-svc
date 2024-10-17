package account

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/cache"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/casbin"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/database"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	if err := p.database.Transaction(ctx, func(trc *gorm.DB) error {
		ue.ID = uint(utils.GetSnowflakeID())
		// 创建账号, 设置一个normal的角色
		if err := trc.Create(&entity.CasbinRule{
			PType: "g",
			V0:    fmt.Sprintf("%d", ue.ID),
			V1:    "ab-normal",
			V2:    "achobeta",
		}).Error; err != nil {
			return err
		}
		ue.Password = hashPassword(ue.Password)
		if err := trc.Create(&ue).Error; err != nil {
			tlog.CtxErrorf(ctx, "create account error: %v", err)
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return status.Errorf(codes.AlreadyExists, "username or email or phone already exists")
			}
			return status.Errorf(codes.Internal, "create account error: %v", err)
		}
		return nil
	}); err != nil {
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

func (p *Permission) CheckToken(ctx context.Context, token string, role int32, act string) (bool, error) {
	claims, err := p.casbin.VerifyToken(token)
	if err != nil {
		tlog.CtxErrorf(ctx, "verify token error: %v", err)
		return false, err
	}
	tlog.CtxInfof(ctx, "%+v\n, role: %d\n, act: %s", claims, role, act)

	isValid := p.casbin.Check(claims["userId"].(string), claims["domain"].(string),
		fmt.Sprintf("v%d", role), claims["object"].(string), act)
	return isValid, nil
}

func (p *Permission) Login(ctx context.Context, req *entity.LoginRequest) (string, error) {
	if req.Type == entity.LoginTypeUsername {
		account := &entity.Account{}
		p.database.Get().Where("username = ?", req.LoginKey).First(&account)
		if utils.ComparePasswords(account.Password, req.LoginPwd) {
			cb := &entity.CasbinRule{}
			if row := p.database.Get().Where("ptype = ?", "g").Where(
				"v0 = ?", account.ID).Find(&cb).RowsAffected; row == 0 {
				return "", fmt.Errorf("no records found for Casbin, please check the data")
			}
			// ptype, v0(userid), v1(object), v2(domain)
			token, err := p.casbin.CreateToken(cb.V0, "data", cb.V2)
			_ = p.cache.Set(ctx, token, strconv.Itoa(int(account.ID)), 30*time.Minute)
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
