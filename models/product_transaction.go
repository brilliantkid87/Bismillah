package models

type ProductTransaction struct {
	ID              int    `json:"id" gorm:"primary_key:auto_increment"`
	ProductID       int    `json:"product_id"`
	ProductName     string `json:"product_name" gorm:"type: varchar(255)"`
	ProductImage    string `json:"product_image" gorm:"type: varchar(255)"`
	ProductPrice    int    `json:"product_price" gorm:"type: int"`
	ProductQuantity int    `json:"product_quantity" gorm:"type: int"`
	TransactionID   int    `json:"transaction_id"`
}
