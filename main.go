package main

import (
	"fmt"
	"golang-microservice/config"
	"golang-microservice/handlers"
	"golang-microservice/models"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  fmt.Println("Hello, Microservices with Golang & Echo!")
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
      "message": "Welcome to Microservices with Golang & Echo!",
		})
	})

	config.ConnectDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config.DB.AutoMigrate(&models.User{}) // Auto to create users table if not exists

	e.POST("/login", handlers.Login)
	r := e.Group("")
	r.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("users/:id", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

