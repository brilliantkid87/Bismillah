package models

type User struct {
	ID int `json:"id" gorm:"primary_key:auto_increment"`
	// IsAdmin     bool                    `json:"is_admin" gorm:"type: bool"`
	Role        string                  `json:"role" gorm:"type: varchar(255)"`
	Name        string                  `json:"name" gorm:"type: varchar(255)"`
	Email       string                  `json:"email" gorm:"type: varchar(255)"`
	Password    string                  `json:"-" gorm:"type: varchar(255)"`
	CartID      int                     `json:"cart_id"`
	Profile     ProfileResponse         `json:"profile"`
	Transaction TransactionUserResponse `json:"transaction"`
}

type UserTransactionResponse struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

type UserCartResponse struct {
	ID int `json:"id"`
	// IsAdmin bool   `json:"is_admin"`
	Role string `json:"role"`
	Name string `json:"name"`
}
type UserProfileResponse struct {
	ID int `json:"id"`
	// IsAdmin bool   `json:"is_admin"`
	Role string `json:"role"`
	Name string `json:"name"`
}

func (UserTransactionResponse) TableName() string {
	return "users"
}

func (UserProfileResponse) TableName() string {
	return "users"
}

func (UserCartResponse) TableName() string {
	return "users"
}
