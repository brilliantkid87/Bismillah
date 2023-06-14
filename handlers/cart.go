package handlers

import (
	"net/http"
	dto "waysbean/dto/result"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

type handlerCard struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCard {
	return &handlerCard{CartRepository}
}

func (h *handlerCard) FindCart(c echo.Context) error {
	carts, err := h.CartRepository.FindCart()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: carts})
}
