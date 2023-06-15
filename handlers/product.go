package handlers

import (
	"net/http"
	productdto "waysbean/dto/product"
	dto "waysbean/dto/result"
	"waysbean/models"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductRepository repositories.ProductRepository
}

func NewProductHandler(productRepository repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{ProductRepository: productRepository}
}

func (h *ProductHandler) FindProduct(c echo.Context) error {
	products, err := h.ProductRepository.FindProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: products})
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	request := new(productdto.CreateProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	product := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Stock:       request.Stock,
		Image:       request.Image,
	}

	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(data)})

}

func convertResponseProduct(u models.Product) productdto.ProductResponse {
	return productdto.ProductResponse{
		ID:          u.ID,
		Name:        u.Name,
		Price:       u.Price,
		Description: u.Description,
		Stock:       u.Stock,
		Image:       u.Image,
	}
}
