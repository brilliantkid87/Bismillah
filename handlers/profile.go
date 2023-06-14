package handlers

import (
	"net/http"
	dto "waysbean/dto/result"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfiletRepository
}

func HandlerProfile(ProfileRepository repositories.ProfiletRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) FindProfile(c echo.Context) error {
	profiles, err := h.ProfileRepository.FindProfile()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: profiles})
}
