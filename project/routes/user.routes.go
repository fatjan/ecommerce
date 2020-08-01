package routes

import (
	"myproject/project/handler"

	"github.com/labstack/echo"
)

// UserConfigRoutes : struct for User config routes
type UserConfigRoutes struct {
	handler *handler.UserConfigHandler
}

// NewUserConfigRoutes : function to init User config routes
func NewUserConfigRoutes(h *handler.UserConfigHandler) *UserConfigRoutes {
	return &UserConfigRoutes{handler: h}
}

// Register : function for register handler in routes
func (r *UserConfigRoutes) Register(v1 *echo.Group) {
	hello := v1.Group("/user")
	hello.GET("/:id", r.handler.GetUserConfig)
	hello.POST("", r.handler.CreateUserConfig)
	hello.PUT("/:id", r.handler.UpdateUserConfig)
	hello.DELETE("/:id", r.handler.DeleteConfig)
	hello.GET("", r.handler.GetHelloWorld)
}
