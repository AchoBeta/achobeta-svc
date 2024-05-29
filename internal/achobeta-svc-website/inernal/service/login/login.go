package login

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
    "achobeta-svc/internal/achobeta-svc-website/config"
    "achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/user"
	"context"
    "encoding/base64"
    "encoding/json"

    "time"
)

func Login(ctx context.Context, req *entity.LoginRequest) (string, error) {
	// 查询用户
//	tlog.CtxInfof(ctx, "2222 login, username:[%s], email:[%s], phone:[%s]", req.Username, req.Email, req.Phone)
	u, err := user.QueryUser(ctx, &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		tlog.CtxErrorf(ctx, "query user error: %v", err)
		return "", err
	}
	// 密码校验
	if !utils.ComparePasswords(u.Password, req.Password) {
		tlog.CtxErrorf(ctx, "password error")
		return "", err
	}
	token, err := createToken(u)
	if err != nil {
		tlog.CtxErrorf(ctx, "set token into redis error, msg:%s", err.Error())
		return "", err
	}
	return token, nil
}

func createToken(user *entity.User) (string, error) {
	var err error
	token := base64.StdEncoding.EncodeToString([]byte(utils.GetSnowflakeUUID()))
	userData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	// 过期时间30分钟
	err = config.GetRedis().Set(token, userData, 30*time.Minute).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}