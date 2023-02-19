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
		fmt.Println(".env file not found. Will proceed and attempt to use existing environment variables.")
	}

	services.ConnectDB()

	router := gin.Default()
	api := router.Group("/", middlewares.CORS)

	posts := api.Group("/posts")
	posts.GET("/", middlewares.JWT, postcontroller.Index)
	posts.POST("/", middlewares.JWT, postcontroller.Create)
	posts.PUT("/:postUuid", middlewares.JWT, postcontroller.Edit)      // todo
	posts.DELETE("/:postUuid", middlewares.JWT, postcontroller.Delete) // todo

	comments := posts.Group("/:postUuid")                                 /* route: /posts/:postUUid */
	comments.GET("/comments", middlewares.JWT, commentcontroller.Index)   // todo
	comments.POST("/comments", middlewares.JWT, commentcontroller.Create) // todo

	auth := api.Group("/auth")
	auth.POST("/register", authcontroller.Register)
	auth.POST("/login", authcontroller.Login)

	oauth := auth.Group("/oauth2")                                /* route: /auth/oauth2 */
	oauth.POST("/google", authcontroller.GoogleInit)              // todo
	oauth.POST("/google/callback", authcontroller.GoogleCallback) // todo

	profile := api.Group("/user")
	profile.GET("/:username", profilecontroller.View)
	profile.GET("/:username/posts", profilecontroller.ViewPosts)            // todo
	profile.GET("/", middlewares.JWT, profilecontroller.ViewSelf)           // todo
	profile.GET("/posts", middlewares.JWT, profilecontroller.ViewPostsSelf) // todo
	profile.POST("/", middlewares.JWT, profilecontroller.EditAvatar)

	router.Run()
}
