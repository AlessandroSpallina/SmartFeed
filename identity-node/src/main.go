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
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"

	"identity-node/src/handler"
)

type config struct {
	debug            bool
	httpServerPort   string
	mqttBrokerHost   string
	mqttBrokerPort   string
	mqttProducerID   string
	mqttWelcomeTopic string
}

func getConfigFromEnv() *config {
	conf := &config{}
	conf.debug, _ = strconv.ParseBool(os.Getenv("IDENTITY_DEBUG"))
	conf.httpServerPort = os.Getenv("IDENTITY_HTTP_SERVER_PORT")
	conf.mqttBrokerHost = os.Getenv("IDENTITY_MQTT_BROKER_HOST")
	conf.mqttBrokerPort = os.Getenv("IDENTITY_MQTT_BROKER_PORT")
	conf.mqttProducerID = os.Getenv("IDENTITY_MQTT_PRODUCER_ID")
	conf.mqttWelcomeTopic = os.Getenv("IDENTITY_MQTT_WELCOME_TOPIC")

	return conf
}

// Il producer mqtt (identity node) riceve richieste su un welcomeTopic fissato,
//   questo è un topic sul quale gli InfoColumn possono fare Publish, ma al quale non
//   possono fare Subscribe, possono quindi scrivere sul topic, ma non leggere i
//   messaggi inviati sul topic; l'unico a poterlo fare è il producer mqtt.
//   Quanto appena detto è possibile grazie ad una gestione dei permessi finemente configurabile
//   see: https://docs.vernemq.com/configuration/file-auth and https://docs.vernemq.com/configuration/db-auth
//
// Se gli InfoNode possono fare publish su welcomeTopic vuol dire che il broker mqtt
//   riconosce id e password forniti a connession time e consente/nega la connessione
//
// I messaggi inviati dagli InfoColumn sul welcomeTopic hanno il seguente formato
//   {"username":"nomeUtenteRiconosciuto", "responseTopic":"nomeTopic"}
//   il producer mqtt può quindi fare una publish sul topic richiesto dall'InfoColumns
//   il formato della risposta sul responseTopic è del tipo [{"tag":"weather", "args":["city":["catania", ...]]}]
//   => PATTERN REQUEST-REPLY su mqtt
//
// see https://github.com/eclipse/paho.mqtt.golang/blob/master/cmd/simple/main.go
func startMQTTProducer(brokerHost, brokerPort, clientID, welcomeTopic string, debug bool) {
	if debug {
		mqtt.DEBUG = log.New(os.Stdout, "", 0)
	}
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker("tcp://" + brokerHost + ":" + brokerPort).SetClientID(clientID)
	opts.SetKeepAlive(2 * time.Second)

	f := handler.RequestReplyRoutine

	opts.SetDefaultPublishHandler(f)

	opts.SetPingTimeout(1 * time.Second)
	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe(welcomeTopic, 0, nil); token.Wait() && token.Error() != nil {
		log.Println("[IDENTITY] MQTT Subscribe err:", token.Error())
		os.Exit(1)
	}

	/*for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is stocazzo #%d!", i)
		token := c.Publish(welcomeTopic, 0, false, text)
		token.Wait()
	}

	time.Sleep(60 * time.Second)

	if token := c.Unsubscribe(welcomeTopic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)

	time.Sleep(1 * time.Second)*/

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

	go startMQTTProducer(conf.mqttBrokerHost, conf.mqttBrokerPort, conf.mqttProducerID, conf.mqttWelcomeTopic, conf.debug)

	startHTTPServer(conf.httpServerPort, conf.debug)

	//model.User
}
