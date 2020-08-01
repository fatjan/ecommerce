package main

import (
	ecommerceController "ecommerce/controllers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome to our ecommerce.")
	})

	// Routes
	e.GET("/employess", ecommerceController.GetEmployees())
	e.POST("/users", ecommerceController.PostUser)
	e.GET("/users/:id", ecommerceController.GetUsers)
	e.PUT("/users/:id", ecommerceController.EditUser)
	e.DELETE("/users/:id", ecommerceController.DeleteUser)

	e.Logger.Fatal(e.Start(":8081"))
}
