/*Identity-Node ha 2 componenti:
  - Server HTTP per caso d'uso 1-1a
	- Client MQTT per caso d'uso 3

Ci saranno quindi due goroutine dispatcher, uno per protocollo, quindi i task verranno eseguiti in parallelo
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	//"identity-node/src/model"
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

// see https://github.com/eclipse/paho.mqtt.golang/blob/master/cmd/simple/main.go
func startMQTTProducer(brokerHost, brokerPort, clientID, defaultRequestChannel string, debug bool) {
	if debug {
		mqtt.DEBUG = log.New(os.Stdout, "", 0)
	}
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	f := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	}

	opts := mqtt.NewClientOptions().AddBroker("tcp://" + brokerHost + ":" + brokerPort).SetClientID(clientID)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// sul request channel l'identity node riceve le richieste e risponde su un canale dedicato per la colonnina (può essere dinamico o statico per colonnina @findme )
	//	if token := c.Subscribe(defaultRequestChannel)

}

func startHTTPServer(port string, debug bool) {
	if !debug { // di default gin parte in debug
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

	//model.User
}
