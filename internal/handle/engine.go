package handle

import (
	"github.com/sirupsen/logrus"

	"etov/internal/dao"
	"etov/internal/repo"
	"etov/internal/response"
	"etov/internal/svc"
)

func GetSupportEngine(ctx *svc.Context) {
	var (
		apiKeyRepo repo.APIKeyRepo = dao.NewAPIKeyDao(ctx.DB)
		resp       response.GetSupportEngineResponse
	)

	platform, err := apiKeyRepo.GetEngineByUid(1)
	if err != nil {
		logrus.Error(err)
		ctx.Error(err)
		return
	}
	resp.Platform = make([]response.Engine, 0, len(platform))
	for _, v := range platform {
		resp.Platform = append(resp.Platform, response.Engine{
			Name: v.KeyName,
			ID:   v.ID,
		})
	}

	uid, ok := ctx.GetUserId()
	resp.Custom = make([]response.Engine, 0)
	if ok {
		custom, _ := apiKeyRepo.GetEngineByUid(uid)
		resp.Custom = make([]response.Engine, 0, len(custom))
		for _, v := range custom {
			resp.Custom = append(resp.Custom, response.Engine{
				Name: v.KeyName,
				ID:   v.ID,
			})
		}
	}
	ctx.Success(resp)
}
