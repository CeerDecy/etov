package response

type ChatItem struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type CreateChatResponse struct {
	Chat *ChatItem `json:"chat"`
}

type GetChatsResponse struct {
	Chats []*ChatItem `json:"chats"`
}
