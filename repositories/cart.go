package repositories

import (
	"waysbean/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCart() ([]models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	GetCart(id int) (models.Cart, error)
}

func NewCartRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FindCart() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("User").Preload("Product").Find(&carts).Error

	return carts, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("User").Preload("Product").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}
