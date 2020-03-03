package handler

import (
	"fmt"

	"github.com/eclipse/paho.mqtt.golang"
)

// ListUserInterests - ritorna la lista di interessi dell'utente che effettua la richiesta
/*func ListUserInterests(c *gin.Context) {
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
}*/

func requestReplayRoutine(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
