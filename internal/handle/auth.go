package handle

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"etov/internal/apierrors"
	"etov/internal/dao"
	"etov/internal/model"
	"etov/internal/repo"
	"etov/internal/request"
	"etov/internal/response"
	"etov/internal/svc"
	"etov/internal/utils"
	"etov/tools"
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

func Login(ctx *svc.Context) {
	var (
		userRepo repo.UserRepo
		req      request.LoginRequest
		user     *model.User
		err      error
	)

	defer func() {
		if err != nil {
			ctx.Error(err)
		}
	}()

	userRepo = dao.NewUserDao(ctx.DB)
	if err = ctx.ShouldBind(&req); err != nil {
		return
	}
	email, phone, err := parseAccount(req.Account)
	if err != nil {
		return
	}

	if email == "" {
		user, _ = userRepo.GetByPhone(phone)
	} else {
		user, _ = userRepo.GetByEmail(email)
	}

	if user == nil {
		err = apierrors.UserNotExistError
		return
	}

	if tools.MD5Str(req.Password, user.Salt) != user.Password {
		err = apierrors.UserPasswordNotMatchError
	}

	token, err := utils.GenerateTokenUsingHs256(user.Id)
	if err != nil {
		err = errors.New("generate token error")
	}

	ctx.Success(&response.LoginResponse{Token: token})
}

func Register(ctx *svc.Context) {
	var (
		userRepo repo.UserRepo
		req      request.RegisterRequest
	)
	userRepo = dao.NewUserDao(ctx.DB)
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}

	email, phone, err := parseAccount(req.Account)
	if err != nil {
		logrus.Error(err)
		ctx.Error(errors.New("账号格式错误"))
		return
	}

	salt := utils.GenerateSalt(8)

	err = userRepo.Create(&model.User{
		Email:     email,
		Phone:     phone,
		Salt:      salt,
		Password:  tools.MD5([]byte(req.Password), []byte(salt)),
		NickName:  "User-" + strconv.Itoa(int(time.Now().Unix())),
		Avatar:    "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	if email == "" {
		ctx.Success(&response.RegisterResponse{Mode: "phone"})
	} else {
		ctx.Success(&response.RegisterResponse{Mode: "email"})
	}
}

func parseAccount(account string) (email string, phone string, err error) {
	if index := strings.Index(account, "@"); index != -1 {
		return account, "", nil
	}
	if len(account) == 11 {
		return "", account, nil
	}
	return "", "", errors.New("account format error")
}
