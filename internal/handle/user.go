package handle

import (
	"errors"

	"github.com/sirupsen/logrus"

	"etov/internal/dao"
	"etov/internal/repo"
	"etov/internal/response"
	"etov/internal/svc"
)

func UserInfo(ctx *svc.Context) {
	userId, exists := ctx.Get("userID")
	if !exists {
		err := errors.New("can't get userID from context")
		logrus.Error(err)
		ctx.Error(err)
		return
	}

	var (
		userRepo repo.UserRepo
		resp     response.UserInfoResponse
	)
	userRepo = dao.NewUserDao(ctx.DB)
	user, err := userRepo.GetByID(userId.(int64))
	if err != nil {
		logrus.Error(err)
		ctx.Error(errors.New("无法获取用户信息"))
		return
	}
	if user == nil {
		ctx.Error(errors.New("无法获取用户信息"))
		return
	}

	resp = response.UserInfoResponse{
		Id:       user.Id,
		NickName: user.NickName,
		Email:    user.Email,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
		ApiKey:   user.ApiKey,
		Validate: user.Validate == "1",
	}
	ctx.Success(resp)
}
