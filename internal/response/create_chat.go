package response

type CreateChatResponse struct {
	ChatId string `json:"chatId"`
}

func NewCreateChatResponse(chatId string) *CreateChatResponse {
	return &CreateChatResponse{ChatId: chatId}
}
