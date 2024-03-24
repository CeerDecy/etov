package tools

import (
	"io"
	"os"
	"strings"

	"github.com/google/uuid"

	"etov/internal/gpt/etovs"
	"etov/internal/gpt/message"
	"etov/internal/request"
	"etov/internal/response"
	"etov/internal/svc"
)

var SummarySupportFiles = map[string]struct{}{
	"docx": {},
	"DOCX": {},
	"doc":  {},
	"DOC":  {},
	"md":   {},
	"MD":   {},
}

func Summary(ctx *svc.Context) {
	var req request.SummaryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}
	summary := etovs.NewSummary()
	msg := message.NewMessages()
	msg.AddChatMessageRoleUserMsg(req.Content)

	if req.FilePath != "" {
		basePath, _ := os.Getwd()
		file, err := os.Open(basePath + req.FilePath)
		if err != nil {
			ctx.Error(err)
			return
		}
		bytes, err := io.ReadAll(file)
		if err != nil {
			ctx.Error(err)
			return
		}
		msg.AddChatMessageRoleUserMsg(string(bytes))
	}
	summary.AppendContext(*msg)
	client, err := ctx.ClientCache.GetClient(req.EngineId, ctx.DB)
	if err != nil {
		ctx.Error(err)
		return
	}
	sess, err := summary.Execute(client)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Stream(sess.HandleStream(nil))
}

// FileUpload /**
func FileUpload(ctx *svc.Context) {
	formFile, err := ctx.FormFile("files")
	if err != nil {
		ctx.Error(err)
		return
	}
	fileName := uuid.New().String()
	split := strings.Split(formFile.Filename, ".")
	fileType := split[1]
	_, ok := SummarySupportFiles[fileType]
	if !ok {
		ctx.ErrorMsg("can't support this file")
		return
	}
	//保存path路径
	filePath := "/static/files/" + fileName + "." + fileType
	basePath, _ := os.Getwd()
	path := basePath + filePath
	err = ctx.SaveUploadedFile(formFile, path)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Success(response.UploadFileResponse{Path: filePath, SourceName: formFile.Filename})
}
