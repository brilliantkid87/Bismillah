package models

type Cart struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int                 `json:"user_id" gorm:"type: int"`
	User          UserCartResponse    `json:"user"`
	ProductID     int                 `json:"product_id" gorm:"type: int"`
	Product       ProductCartResponse `json:"product"`
	OrderQuantity int                 `json:"order_quantity" gorm:"type: int"`
}

type CartUserResponse struct {
	ID            int `json:"id"`
	ProductID     int `json:"product_id"`
	OrderQuantity int `json:"order_quantity"`
	UserID        int `json:"user_id"`
}

type CartProductResponse struct {
	ProductID     int                 `json:"-"`
	Product       ProductCartResponse `json:"product"`
	OrderQuantity int                 `json:"order_quantity"`
	UserID        int                 `json:"user_id"`
}

func (CartUserResponse) TableName() string {
	return "carts"
}

func (CartProductResponse) TableName() string {
	return "carts"
}
