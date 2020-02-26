package main

// see https://github.com/demo-apps/go-gin-app/blob/a4cfa04a9146109ca88e0ecaba8b53b2af2159d9/handlers.user.go

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "pong",
	})
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if !isUserValid(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failure",
			"message": "invalid credentials provided",
		})
		return
	}

	token := generateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successful login",
	})

}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func register(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println(username, password)

	if _, err := registerNewUser(username, password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": err.Error(),
		})
		return
	}

	// If the user is created, set the token in a cookie and log the user in
	token := generateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successful registration",
	})
}
