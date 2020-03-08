package handler

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func RequestRoutine(client mqtt.Client, msg mqtt.Message) {
	requestInterface := &RequestBody{}
	if err := json.Unmarshal(msg.Payload(), requestInterface); err != nil {
		log.Printf("[FEED-PROVIDER] Received bad request from %s: %s [error: %s]", msg.Topic(), msg.Payload(), err.Error())
		return
	}

	// Unmarshal
	request := json.Unmarshal(msg.Payload(), requestInterface)

	request
}

func getFeed(topic string) {

}
