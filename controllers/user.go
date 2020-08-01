package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetUsers is to get all users
func GetUsers(c echo.Context) err {
	result := models.getUser()
	println("Employee is sucessfully obtained.")
	return c.JSON(http.StatusOK, result)
}

// PostUser is to add a user
func PostUser(c echo.Context) err {
	result := models.createUser()
	println("User is sucessfully added.")
	return c.JSON(http.StatusOK, result)
}

// EditUser is to edit user
func EditUser(c echo.Context) err {
	result := models.updateUser()
	println("User is sucessfully updated.")
	return c.JSON(http.StatusOK, result)
}

// DeleteUser is to delete user
func DeleteUser(c echo.Context) err {
	result := models.deleteUser()
	println("User is sucessfully deleted.")
	return c.JSON(http.StatusOK, result)
}
