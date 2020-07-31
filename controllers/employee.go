package controllers

import (
	"ecommerce/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetEmployees(c echo.Context) error {
	result := models.GetEmployee()
	println("Employee is sucessfully obtained.")
	return c.JSON(http.StatusOK, result)
}
