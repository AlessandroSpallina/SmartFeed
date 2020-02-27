package controller

import (
	"identity-node/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListTags - ritorna tutti i tag disponibili sul sistema
func ListTags(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": repository.ListTags(),
	})
}
