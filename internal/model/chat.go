package model

type ChatRequest struct {
	Uid     int64  `json:"uid"`
	Content string `json:"content"`
}
