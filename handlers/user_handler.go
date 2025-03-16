package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang-microservice/config"
	"golang-microservice/models"
)

func GetUsers(c echo.Context) error {
	var users []models.User
	config.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	newUser := new(models.User)
	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	config.DB.Create(newUser)
	return c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if config.DB.First(&user, id).Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	config.DB.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if config.DB.First(&user, id).Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	config.DB.Delete(&user)
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}
