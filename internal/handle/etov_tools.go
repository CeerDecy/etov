package handle

import (
	"etov/internal/dao"
	"etov/internal/repo"
	"etov/internal/response"
	"etov/internal/svc"
)

func GetPublicTools(ctx *svc.Context) {
	var (
		toolRepo repo.ToolRepo
		resp     []response.GetPublicToolResponse
	)
	toolRepo = dao.NewToolDao(ctx.DB)

	tools, err := toolRepo.GetAllPublic()
	if err != nil {
		ctx.Error(err)
	}
	for _, tool := range tools {
		resp = append(resp, response.GetPublicToolResponse{
			Name:        tool.Name,
			Logo:        tool.Logo,
			URL:         tool.URL,
			Description: tool.Description,
			Disable:     tool.Disable == "Y",
		})
	}
	ctx.Success(resp)
}
