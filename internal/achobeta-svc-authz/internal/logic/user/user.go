package user

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/cache"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/database"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"

	"context"
	"fmt"
)

type UserLogic struct {
	database database.Database
	cache    cache.Cache
}

func New(db database.Database, c cache.Cache) *UserLogic {
	return &UserLogic{
		database: db,
		cache:    c,
	}
}

// Create 创建用户
func (u *UserLogic) CreateUser(ctx context.Context, user *entity.User) (uint, error) {
	if _, err := u.database.Create(user); err != nil {
		tlog.CtxErrorf(ctx, "create user error: %v", err)
		return 0, err
	}
	return user.ID, nil
}

func (u *UserLogic) Modify(ctx context.Context, user *entity.User) error {
	if user == nil {
		return fmt.Errorf("modify user error, user is nil")
	}
	if _, err := u.database.Update(user); err != nil {
		return err
	}
	tlog.Infof("modify user info success, id: %d", user.ID)
	return nil
}

func (u *UserLogic) Query(ctx context.Context, user *entity.User) error {
	if err := u.database.Get().Where(user).First(user).Error; err != nil {
		tlog.CtxErrorf(ctx, "query user error: %v", err)
		return err
	}
	return nil
}
