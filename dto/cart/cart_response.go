package cartdto

type CartResponse struct {
	ID            int `json:"id"`
	ProductID     int `json:"product_id" gorm:"type: int"`
	OrderQuantity int `json:"order_quantity" gorm:"type: int"`
	UserID        int `json:"user_id"`
}