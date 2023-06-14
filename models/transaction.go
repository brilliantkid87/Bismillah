package models

type Transaction struct {
	ID                 int                     `json:"id" gorm:"primary_key:auto_increment"`
	UserID             int                     `json:"user_id"`
	User               UserTransactionResponse `json:"user"`
	Name               string                  `json:"name" gorm:"type: varchar(255)"`
	Email              string                  `json:"email" gorm:"type: varchar(255)"`
	Phone              string                  `json:"phone" gorm:"type: varchar(255)"`
	Address            string                  `json:"address" gorm:"type: varchar(255)"`
	ProductID          int                     `json:"product_id"`
	ProductTransaction []ProductTransaction    `json:"products" gorm:"foreignKey:TransactionID"`
	TotalCounterQty    int                     `json:"total_counterqty" gorm:"type:int"`
	TotalPrice         int                     `json:"total_price" gorm:"type: int"`
	Status             string                  `json:"status" gorm:"type: varchar(255)"`
}

type TransactionUserResponse struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	TotalCounterQty int    `json:"total_counterqty"`
	TotalPrice      int    `json:"total_price"`
	Status          string `json:"status" gorm:"type: varchar(255)"`
}

func (TransactionUserResponse) TableName() string {
	return "transactions"
}
