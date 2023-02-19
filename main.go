package main

import (
	"bird-app/controllers/authcontroller"
	"bird-app/controllers/postcontroller"
	"bird-app/middlewares"
	"bird-app/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found. Will proceed and attempt to use existing environment variables.")
	}

	services.ConnectDB()

	router := gin.Default()
	api := router.Group("/", middlewares.CORS)

	posts := api.Group("/posts")
	posts.GET("/", postcontroller.Index)
	posts.POST("/", postcontroller.Create)

	auth := api.Group("/auth")
	auth.POST("/register", authcontroller.Register)
	auth.POST("/login", authcontroller.Login)

	router.Run()
}
