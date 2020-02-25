/*Identity-Node ha 2 componenti:
  - Server HTTP per caso d'uso 1-1a
	- Client MQTT per caso d'uso 3

Ci saranno quindi due goroutine dispatcher, uno per protocollo, quindi i task verranno eseguiti in parallelo
*/
package main

import (
	"log"
	"os"

	//"identity-node/controller/user_controller"

	//"controller"

	"github.com/gin-gonic/gin"
)

var conf = getConfigFromEnv()

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

func startHTTPServer() {
	r := gin.Default()

	//r.GET("/ping", controller.Ping)

	r.Run(":" + conf.httpServerPort)
}

func main() {
	log.Println("[IDENTITY] Conf:", *conf)

	startHTTPServer()
}
