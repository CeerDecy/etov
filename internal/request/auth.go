package request

type HasRegisteredRequest struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
