package handle

import (
	"etov/internal/dao"
	"etov/internal/repo"
	"etov/internal/request"
	"etov/internal/response"
	"etov/internal/svc"
)

func HasRegistered(ctx *svc.Context) {
	var (
		userRepo repo.UserRepo
		req      request.HasRegisteredRequest
	)
	userRepo = dao.NewUserDao(ctx.DB)
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}
	user, err := userRepo.GetByEmail(req.Email)
	if err != nil || user == nil {
		ctx.Success(&response.HasRegisteredResponse{Flag: false})
	} else {
		ctx.Success(&response.HasRegisteredResponse{Flag: true})
	}
}
