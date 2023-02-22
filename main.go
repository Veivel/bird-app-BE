package main

import (
	"bird-app/controllers/authcontroller"
	"bird-app/controllers/commentcontroller"
	"bird-app/controllers/postcontroller"
	"bird-app/controllers/profilecontroller"
	"bird-app/middlewares"
	"bird-app/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[ERR] .env file not found. Will proceed anyway.")
	}

	services.ConnectDB()

	router := gin.Default()
	router.Use(middlewares.CORS)

	posts := router.Group("/posts")
	posts.GET("/", middlewares.JWT, postcontroller.Index)
	posts.POST("/", middlewares.JWT, postcontroller.Create)
	posts.GET("/:postUuid", middlewares.JWT, postcontroller.Show)
	posts.PUT("/:postUuid", middlewares.JWT, postcontroller.Edit)
	posts.DELETE("/:postUuid", middlewares.JWT, postcontroller.Delete)

	comments := posts.Group("/:postUuid") /* route: /posts/:postUuid */
	comments.GET("/comments", middlewares.JWT, commentcontroller.Index)
	comments.POST("/comments", middlewares.JWT, commentcontroller.Create)

	auth := router.Group("/auth")
	auth.POST("/register", authcontroller.Register)
	auth.POST("/login", authcontroller.Login)

	oauth := auth.Group("/oauth2")                                /* route: /auth/oauth2 */
	oauth.POST("/google", authcontroller.GoogleInit)              // todo
	oauth.POST("/google/callback", authcontroller.GoogleCallback) // todo

	profile := router.Group("/user")
	profile.GET("/:username", profilecontroller.View)
	profile.GET("/:username/posts", profilecontroller.ViewPosts)

	profile.GET("/", middlewares.JWT, profilecontroller.ViewSelf)
	profile.GET("/posts", middlewares.JWT, profilecontroller.ViewPostsSelf)
	profile.POST("/", middlewares.JWT, profilecontroller.EditAvatar)

	router.Run()
	// todo: refactor
	// todo: context
}
