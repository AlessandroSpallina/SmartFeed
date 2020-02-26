package controller

import (
	"identity-node/src/model"
	"identity-node/src/repository"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

// Pong - check availability
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "pong",
	})
}

// Login -
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if !repository.IsValidUser(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failure",
			"message": "invalid credentials provided",
		})
		return
	}

	token := generateSessionToken()
	repository.SaveSession(username, token)
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successful login",
	})
}

// Logout -
func Logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// @findme : a questo punto la sessione su db dovrebbe essere distrutta
	// es. repository.DeleteSession()

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

// Register - register new user
func Register(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println(username, password)

	u := model.User{}
	u.Username = username
	u.Password = password

	if _, err := repository.SaveUser(u); err != nil {
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

func ShowUser(c *gin.Context) {

}
