package api

// RegisterRequest is only used for email and password registration
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID     int     `json:"id"`
	Name   *string `json:"name"`
	Email  string  `json:"email"`
	Avatar string  `json:"avatar"`
}
