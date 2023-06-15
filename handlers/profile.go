package handlers

import (
	"net/http"
	profiledto "waysbean/dto/profile"
	dto "waysbean/dto/result"
	"waysbean/models"
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

func (h *handlerProfile) CreateProfile(c echo.Context) error {
	request := new(profiledto.CreateProfileRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	profile := models.Profile{
		Phone:        request.Phone,
		Address:      request.Address,
		UserID:       request.UserID,
		ImageProfile: request.ImageProfile,
	}

	data, err := h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProfile(data)})
}

func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		ID:           u.ID,
		Phone:        u.Phone,
		Name:         u.Name,
		PostCode:     u.PostCode,
		Address:      u.Address,
		ImageProfile: u.ImageProfile,
		UserID:       u.UserID,
		User:         u.User,
	}
}
