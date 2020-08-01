package handler

import (
	"net/http"

	"myproject/project/entity"
	"myproject/project/service"

	"github.com/labstack/echo"
)

// UserConfigHandler : struct for User config handler
type UserConfigHandler struct {
	userConfigService service.IUserConfigService
}

// H : struct for empty
type H map[string]interface{}

// NewUserConfigHandler : function to init User config handler
func NewUserConfigHandler(s *service.UserConfigService) *UserConfigHandler {
	return &UserConfigHandler{userConfigService: s}
}

// CreateTaxPayerConfig is handler to create new taxpayer config
func (h *UserConfigHandler) CreateUserConfig(c echo.Context) error {
	u := &entity.User{}

	if err := c.Bind(u); err != nil {
		return err
	}

	err := h.userConfigService.CreateOneService(u)

	if err == nil {
		return c.JSON(http.StatusCreated, H{
			"code":    200,
			"message": "User is successfully added",
		})
	}
	return err
}

// GetUserConfig is handler to get User datas by specific id
func (h *UserConfigHandler) GetUserConfig(c echo.Context) error {
	id := c.Param("id")

	getOne, err := h.userConfigService.GetOneService(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, H{"code": 4000001, "message": "User Not Found"})
	}
	return c.JSON(http.StatusOK, getOne)
}

// DeleteConfig is handler to delete User datas by specific id
func (h *UserConfigHandler) DeleteConfig(c echo.Context) error {
	id := c.Param("id")

	err := h.userConfigService.DeleteService(id)

	if err == nil {
		return c.JSON(http.StatusCreated, H{
			"code":    200,
			"message": "User is successfully deleted",
		})
	}
	return err
}

// UpdateUserConfig is handler to update user
func (h *UserConfigHandler) UpdateUserConfig(c echo.Context) error {
	u := &entity.User{}

	if err := c.Bind(u); err != nil {
		return err
	}
	id := c.Param("id")

	err := h.userConfigService.UpdateService(id, u)

	if err == nil {
		return c.JSON(http.StatusCreated, H{
			"code":    200,
			"message": "user is successfully updated",
		})
	}
	return err
}

// GetHelloWorld is handler to default path
func (h *UserConfigHandler) GetHelloWorld(c echo.Context) error {
	return c.JSON(http.StatusCreated, H{
		"code":    200,
		"message": "Hello World",
	})
}
