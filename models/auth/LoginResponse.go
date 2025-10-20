package auth

type LoginResponse struct {
	Name    string `json:"name"`
	Token   string `json:"token"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Message string `json:"message"`
}
