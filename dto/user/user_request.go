package userdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	CartID   int    `json:"cart_id" form:"cart_id"`
}
