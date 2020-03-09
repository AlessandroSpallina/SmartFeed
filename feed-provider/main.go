package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"./handler"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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

func startMQTTProducer(brokerHost, brokerPort, clientID, welcomeTopic string, debug bool) {
	if debug {
		mqtt.DEBUG = log.New(os.Stdout, "", 0)
	}
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker("tcp://" + brokerHost + ":" + brokerPort).SetClientID(clientID)
	opts.SetKeepAlive(2 * time.Second)

	//f := handler.RequestReplyRoutine

	f := handler.RequestRoutine

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
}

/*
func startHTTPServer(port string, debug bool) {
	if !debug { // di default gin parte in debug
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	initializeRoutes(router)
	router.Run(":" + port)
}
*/

func main() {
	conf := getConfigFromEnv()
	log.Println("[IDENTITY] Conf:", *conf)

	go startMQTTProducer(conf.mqttBrokerHost, conf.mqttBrokerPort, conf.mqttProducerID, conf.mqttWelcomeTopic, conf.debug)

	//startHTTPServer(conf.httpServerPort, conf.debug)

	//model.User
}
