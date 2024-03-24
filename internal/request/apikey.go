package request

type AddAPIKeyRequest struct {
	TokenName string `json:"token_name"`
	Token     string `json:"token"`
	Host      string `json:"host"`
	ModelTag  string `json:"model_tag"`
}

type UpdateTokenRequest struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Host  string `json:"host"`
	Model string `json:"model"`
}

type DeleteTokenRequest struct {
	ID int64 `json:"ID"`
}
