package userdto

type UserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	CartID int    `json:"cart_id" form:"cart_id"`
}
