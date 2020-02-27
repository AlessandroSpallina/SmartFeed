package controller

import (
	"identity-node/src/model"
	"identity-node/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// CreateUserInterest - salva l'interesse ricevuto per l'utente loggato
func CreateUserInterest(c *gin.Context) {
	token, _ := c.Cookie("token")
	user, err := repository.FindUserBySession(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": err.Error(),
		})
		return
	}

	var i model.Interest
	c.BindJSON(&i)

	i.User = user.Username

	repository.SaveInterest(i)

	c.JSON(http.StatusOK, i)
}
