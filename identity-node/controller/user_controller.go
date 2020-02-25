package controller

var version = "1.0.0"

# findme see https://stackoverflow.com/questions/52026284/accessing-local-packages-within-a-go-module-go-1-11

import (
  "github.com/gin-gonic/gin"
)


func Ping(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "pongo",
  })
