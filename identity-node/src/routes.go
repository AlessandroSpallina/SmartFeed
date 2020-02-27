package main

import (
	"identity-node/src/controller"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	router.Use(setUserStatus())

	router.GET("/ping", controller.Pong)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/", ensureLoggedIn(), controller.ShowUser)
		userRoutes.POST("/login", ensureNotLoggedIn(), controller.Login)
		userRoutes.GET("/logout", ensureLoggedIn(), controller.Logout)
		userRoutes.POST("/register", ensureNotLoggedIn(), controller.Register)

		// @findme : TODO
		// CRUD for user interests
		/*userRoutes.GET("/interests", ensureLoggedIn(), controller.ListUserInterests)

		userRoutes.GET("/interest/{:id}", ensureLoggedIn(), controller.ReadUserInterest)
		userRoutes.POST("/interest", ensureLoggedIn(), controller.CreateUserInterest)
		userRoutes.POST("/interest/{:id}", ensureLoggedIn(), controller.UpdateUserInterest)
		userRoutes.DELETE("/interest/{:id}", ensureLoggedIn(), controller.DeleteUserInterest)*/
	}

	router.GET("/tags", controller.ListTags)
	/*tagRoutes := router.Group("/tag")
	{
		@findme : qui casi d'uso CRUD per i tag inseriti da un "admin"
		tagRoutes.GET("/tags", controller.ListTags)
		...
	}*/
}
