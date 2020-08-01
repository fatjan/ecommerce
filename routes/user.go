package routes

import (
	"ecommerce/config"
	ecommerceController "ecommerce/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// User export
func User(e *echo.Echo) {
	// flag := controller.Handler{Flag: 0} // 0 means not a test case
	// User API
	user := e.Group(path + "/user")
	// Use JWT
	user.Use(middleware.JWT([]byte(config.JWTSecret)))
	// Register Endpoint
	user.POST("", ecommerceController.AddUser, acl("create_user"))
	user.GET("/all", ecommerceController.GetAllUser, acl("read_user"))
	user.POST("/dropdown", ecommerceController.GetAllUserDropDownByMultiProductID)
	user.GET("/dropdown/:user_id", ecommerceController.GetAllUserDropDown)
	user.GET("/:user_id", ecommerceController.GetUserByID, acl("read_user"))
	user.PUT("/:user_id", ecommerceController.UpdateUser, acl("update_user"))
	user.DELETE("/:user_id", ecommerceController.DeleteUserByID, acl("delete_customer"))
}
