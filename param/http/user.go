package param

type GetUserByIDRequest struct {
	UserID string `json:"user_id"`
}

type GetUserByIDResponse struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
