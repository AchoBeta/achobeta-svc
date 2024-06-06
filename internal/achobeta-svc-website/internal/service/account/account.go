package account

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"achobeta-svc/internal/achobeta-svc-website/config"
	"achobeta-svc/internal/achobeta-svc-website/internal/entity"
	"context"
)

func CreateAccount(ctx context.Context, ue *entity.Account) error {
	ue.ID = uint(utils.GetSnowflakeID())
	result := config.GetMysql().Create(&ue)
	if result.Error != nil {
		return result.Error
	}
	tlog.CtxInfof(ctx, "create account, username:[%s], email:[%s], phone:[%s]", ue.Username, ue.Email, ue.Phone)
	return nil
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
