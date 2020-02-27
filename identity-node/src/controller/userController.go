package controller

import (
	"identity-node/src/model"
	"identity-node/src/repository"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	repository.SaveSession(model.Session{User: username, Token: token})
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
	dateofbirth, _ := time.Parse("1994-04-30", c.PostForm("date-of-birth"))
	gender := c.PostForm("gender")
	phone := c.PostForm("phone")
	email := c.PostForm("email")

	log.Println(username, password)

	u := model.User{}
	u.Username = username
	u.Password = password
	u.DateOfBirth = dateofbirth
	u.Gender = gender
	u.Phone = phone
	u.Email = email

	if _, err := repository.SaveUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": err.Error(),
		})
		return
	}

	// If the user is created, set the token in a cookie and log the user in
	/*token := generateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)*/

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

	c.JSON(http.StatusOK, user)

}

// ListUserInterests - ritorna la lista di interessi dell'utente che effettua la richiesta
func ListUserInterests(c *gin.Context) {
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
	c.JSON(http.StatusOK, repository.ListInterestsByUser(user.Username))
}