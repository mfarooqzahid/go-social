package models

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
