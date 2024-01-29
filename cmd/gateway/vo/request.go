package vo

type CreateUserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email  string	`json:"email"` 
}