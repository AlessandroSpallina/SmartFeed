package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

/*const (
	httpServerPort = os.Getenv("IDENTITY_NODE_PORT")
)*/

type config struct {
	httpServerPort string
	mqttBrokerHost string
	mqttBrokerPort string
}

func getConfigFromEnv() *config {
	conf := &config{}
	conf.httpServerPort = os.Getenv("IDENTITY_NODE_PORT")
	conf.mqttBrokerHost = os.Getenv("MQTT_BROKER_HOST")
	conf.mqttBrokerPort = os.Getenv("MQTT_BROKER_PORT")

	return conf
}

/*
	Identity-Node ha 2 componenti:
	  * Server HTTP per caso d'uso 1-1a
		* Client MQTT per caso d'uso 3

	Ci saranno quindi due goroutine dispatcher, uno per protocollo, quindi i task verranno eseguiti in parallelo
*/
func main() {

	config := getConfigFromEnv()
	fmt.Println("Conf:", *config)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongo",
		})
	})

	r.Run(":3000")
}
