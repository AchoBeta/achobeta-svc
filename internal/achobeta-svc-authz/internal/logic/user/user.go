package user

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/cache"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/database"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"

	"context"
	"fmt"
)

type Logic struct {
	database database.Database
	cache    cache.Cache
}

func New(db database.Database, c cache.Cache) *Logic {
	return &Logic{
		database: db,
		cache:    c,
	}
}

// CreateUser  创建用户
func (u *Logic) CreateUser(ctx context.Context, user *entity.User) (uint, error) {
	if _, err := u.database.Create(ctx, user); err != nil {
		tlog.CtxErrorf(ctx, "create user error: %v", err)
		return 0, err
	}
	return user.ID, nil
}

func (u *Logic) Modify(ctx context.Context, user *entity.User) error {
	if user == nil {
		return fmt.Errorf("modify user error, user is nil")
	}
	if _, err := u.database.Update(ctx, user); err != nil {
		return err
	}
	tlog.Infof("modify user info success, id: %d", user.ID)
	return nil
}

func (u *Logic) Query(ctx context.Context, user *entity.User) error {
	if err := u.database.Get().Where(user).First(user).Error; err != nil {
		tlog.CtxErrorf(ctx, "query user error: %v", err)
		return err
	}
	return nil
}
