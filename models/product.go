package models

type Product struct {
	ID          int                   `json:"id" gorm:"primary_key:auto_increment"`
	Name        string                `json:"name" gorm:"type: varchar(255)"`
	Price       int                   `json:"price" gorm:"type: int"`
	Description string                `json:"description" gorm:"type: text"`
	Image       string                `json:"image" gorm:"type: varchar(255)"`
	Stock       int                   `json:"stock" gorm:"type: int"`
	Cart        []CartProductResponse `json:"cart"`
}

type ProductCartResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
}

func (ProductCartResponse) TableName() string {
	return "products"
}
