package param

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type RegisterResponse struct {
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Token     string `json:"token"`
}
