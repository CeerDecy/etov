package response

type UserInfoResponse struct {
	Id       int64  `json:"id"`
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	ApiKey   string `json:"apiKey"`
	Validate bool   `json:"validate"`
}
