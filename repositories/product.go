package repositories

import (
	"waysbean/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	GetProduct(id int) (models.Product, error)
}

func NewProductRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FindProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Cart").Find(&products).Error

	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("Cart").First(&product, ID).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}
