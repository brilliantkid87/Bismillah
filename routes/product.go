package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.NewProductRepository(mysql.DB)
	h := handlers.NewProductHandler(productRepository)

	e.GET("/products", h.FindProduct)
	e.POST("/product", h.CreateProduct)
}
