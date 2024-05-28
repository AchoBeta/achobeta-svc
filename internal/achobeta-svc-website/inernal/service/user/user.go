package user

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-website/config"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"context"
)

func CreateUser(ctx context.Context, ue *entity.User) error {
	tlog.CtxInfof(ctx, "create user, username:[%s], email:[%s], phone:[%s]", ue.Username, ue.Email, ue.Phone)
	result := config.GetMysql().Create(&ue)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func QueryUser(ctx context.Context, params *entity.User) (*entity.User, error) {
	user := &entity.User{}
	tlog.CtxInfof(ctx, "query user, params:[%+v\n]", params)
	result := config.GetMysql().Debug().Where(params).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
