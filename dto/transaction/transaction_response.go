package transactiondto

type TransactionResponse struct {
	UserID          int    `json:"user_id"`
	Name            string `json:"name" gorm:"type: varchar(255)"`
	Email           string `json:"email" gorm:"type: varchar(255)"`
	Phone           string `json:"phone" gorm:"type: varchar(255)"`
	Address         string `json:"address" gorm:"type: varchar(255)"`
	TotalCounterQty int    `json:"total_counterqty"`
	TotalPrice      int    `json:"total_price"`
	Status          string `json:"status" gorm:"type: varchar(255)"`
	ProductID       int    `json:"product_id"`
}
