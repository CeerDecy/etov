package response

type HasRegisteredResponse struct {
	Flag bool `json:"flag"`
}

type RegisterResponse struct {
	Mode string `json:"mode"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
