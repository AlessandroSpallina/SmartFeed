package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	router.Use(setUserStatus())

	router.GET("/ping", pong)

	userRoutes := router.Group("/u")
	{
		// qui index di profilo router.GET("/", )
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}
}
