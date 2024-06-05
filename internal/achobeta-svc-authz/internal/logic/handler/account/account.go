package account

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"context"
)

// CreateAccount 创建账号
// 方法内部对密码进行加密, 外层调用无需关心加密逻辑
func CreateAccount(ctx context.Context, ue *entity.Account) error {
	ue.ID = uint(utils.GetSnowflakeID())
	ue.Password = hashPassword(ue.Password)
	result := config.GetMysql().Create(&ue)
	if result.Error != nil {
		return result.Error
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

func QueryAccount(ctx context.Context, params *entity.Account) (*entity.Account, error) {
	account := &entity.Account{}
	tlog.CtxInfof(ctx, "query account, params:[%+v\n]", params)
	result := config.GetMysql().Debug().Where(params).First(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}
