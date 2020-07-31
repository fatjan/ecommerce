import (
	"net/http"
	"ecommerce/models"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) err {
	result := models.getUser()
	println("Employee is sucessfully obtained.")
	return c.JSON(http.StatusOK, result)
}

func PostUsers(c echo.Context) err {
	result := models.postUser()
	println("User is sucessfully added.")
	return c.JSON(http.StatusOK, result)
}