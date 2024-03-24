package handle

import (
	"time"

	"github.com/sirupsen/logrus"

	"etov/internal/dao"
	"etov/internal/model"
	"etov/internal/repo"
	"etov/internal/request"
	"etov/internal/response"
	"etov/internal/svc"
)

func SaveAPIKey(ctx *svc.Context) {
	var req request.AddAPIKeyRequest
	var apikeyDao repo.APIKeyRepo = dao.NewAPIKeyDao(ctx.DB)
	if err := ctx.ShouldBind(&req); err != nil {
		logrus.Error(err)
		ctx.Error(err)
		return
	}
	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.ErrorMsg("can't get userid")
		return
	}
	apikey := model.APIKey{
		UID:       userId,
		KeyName:   req.TokenName,
		APIKey:    req.Token,
		ModelTag:  req.ModelTag,
		Host:      req.Host,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := apikeyDao.SaveAPIKey(apikey); err != nil {
		ctx.ErrorMsg("保存失败")
		return
	}
	ctx.Success(nil)
}

func GetAPIKeys(ctx *svc.Context) {
	var apikeyDao repo.APIKeyRepo = dao.NewAPIKeyDao(ctx.DB)
	var resp response.GetAPIKeysResponse

	userId, ok := ctx.GetUserId()
	if !ok {
		ctx.ErrorMsg("can't get userid")
		return
	}

	apiKeys, err := apikeyDao.GetEngineByUid(userId)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.APIKeys = make([]response.APIKey, len(apiKeys))

	for i, key := range apiKeys {
		length := len(key.APIKey)
		token := key.APIKey[0:5] + "*******************" + key.APIKey[length-5:length]
		resp.APIKeys[i].ID = key.ID
		resp.APIKeys[i].Name = key.KeyName
		resp.APIKeys[i].Token = token
		resp.APIKeys[i].Host = key.Host
		resp.APIKeys[i].Model = key.ModelTag
	}
	ctx.Success(resp)
}

func UpdateToken(ctx *svc.Context) {
	var req request.UpdateTokenRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.ErrorMsg("参数错误")
		return
	}
	var apikeyDao repo.APIKeyRepo = dao.NewAPIKeyDao(ctx.DB)
	apiKey, err := apikeyDao.GetEngineByiId(req.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	key := model.APIKey{
		ID:        req.ID,
		UID:       apiKey.UID,
		KeyName:   req.Name,
		APIKey:    req.Token,
		ModelTag:  req.Model,
		Host:      req.Host,
		CreatedAt: apiKey.CreatedAt,
		UpdatedAt: time.Now(),
	}
	if err = apikeyDao.SaveAPIKey(key); err != nil {
		ctx.Error(err)
		return
	}
	ctx.Success(nil)
}

func DeleteToken(ctx *svc.Context) {
	var req request.DeleteTokenRequest
	var apikeyDao repo.APIKeyRepo = dao.NewAPIKeyDao(ctx.DB)
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}
	if err := apikeyDao.DeleteAPIKey(req.ID); err != nil {
		ctx.Error(err)
		return
	}
	ctx.Success(nil)
}
