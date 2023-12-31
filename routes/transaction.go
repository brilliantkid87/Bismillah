package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	delta := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(delta)

	e.GET("/transactions", h.FindTransaction)
	// e.GET("/user/:id", h.c)
	e.POST("/transaction", h.CreateTransaction)
	e.GET("/transaction-user/:id", h.GetTransactionByUser)
}
