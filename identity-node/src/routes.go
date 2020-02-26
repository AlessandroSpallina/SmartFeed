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
		// qui index di profilo router.GET("/", )
		userRoutes.GET("/", ensureLoggedIn(), controller.ShowUser)
		userRoutes.POST("/login", ensureNotLoggedIn(), controller.Login)
		userRoutes.GET("/logout", ensureLoggedIn(), controller.Logout)
		userRoutes.POST("/register", ensureNotLoggedIn(), controller.Register)
	}
}
