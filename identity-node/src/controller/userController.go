package controller

import (
	"fmt"
	"identity-node/src/model"
	"identity-node/src/repository"
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
	var u model.User
	c.BindJSON(&u)

	if !repository.IsValidUser(u.Username, u.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failure",
			"message": "invalid credentials provided",
		})
		return
	}

	token := generateSessionToken()
	repository.SaveSession(model.Session{User: u.Username, Token: token})
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
	var u model.User
	c.BindJSON(&u)

	fmt.Println(u)

	if _, err := repository.SaveUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successful registration",
	})
}

// ShowUser -
func ShowUser(c *gin.Context) {
	// se qui è già loggato, perchè sono dietro il middleware ensureLoggedIn
	token, _ := c.Cookie("token")
	user, err := repository.FindUserBySession(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": err.Error(),
		})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, user)

}
