package response

type APIKey struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Host  string `json:"host"`
	Model string `json:"model"`
}

type GetAPIKeysResponse struct {
	APIKeys []APIKey `json:"APIKeys"`
}
