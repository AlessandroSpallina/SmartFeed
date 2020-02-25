/*Identity-Node ha 2 componenti:
  - Server HTTP per caso d'uso 1-1a
	- Client MQTT per caso d'uso 3

Ci saranno quindi due goroutine dispatcher, uno per protocollo, quindi i task verranno eseguiti in parallelo
*/
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type config struct {
	debug          bool
	httpServerPort string
	mqttBrokerHost string
	mqttBrokerPort string
}

func getConfigFromEnv() *config {
	conf := &config{}
	conf.debug, _ = strconv.ParseBool(os.Getenv("IDENTITY_DEBUG"))
	conf.httpServerPort = os.Getenv("IDENTITY_NODE_PORT")
	conf.mqttBrokerHost = os.Getenv("MQTT_BROKER_HOST")
	conf.mqttBrokerPort = os.Getenv("MQTT_BROKER_PORT")

	return conf
}

func startHTTPServer(port string, debug bool) {

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	initializeRoutes(router)
	router.Run(":" + port)
}

func main() {
	conf := getConfigFromEnv()
	log.Println("[IDENTITY] Conf:", *conf)

	startHTTPServer(conf.httpServerPort, conf.debug)
}
