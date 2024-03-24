package handle

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

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
			Params:      parseParams(tool.Params),
			Description: tool.Description,
			Disable:     tool.Disable == "Y",
		})
	}
	ctx.Success(resp)
}

func parseParams(param string) string {
	if param == "" {
		return param
	}
	dict := make(map[string]any)
	err := json.Unmarshal([]byte(param), &dict)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	var ps = make([]string, 0, len(dict))
	for k, v := range dict {
		ps = append(ps, fmt.Sprintf("%s=%v", k, v))
	}
	join := strings.Join(ps, "&")
	return join
}
