package productdto

type ProductResponse struct {
	ID            int    `json:"id"`
	ProductName   string `json:"product_name" form:"product_name"`
	Price         int    `json:"price" form:"price"`
	Stock         int    `json:"stock" form:"stock"`
	Description   string `json:"description" form:"description"`
	TransactionID int    `json:"transaction_id" form:"transaction_id"`
}
