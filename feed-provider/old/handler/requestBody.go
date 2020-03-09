package handler

import (
	"encoding/json"
	"identity-node/src/repository"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// RequestReplyRoutine - Main routine for this microservice in mqtt protocol
//
// I messaggi inviati dagli InfoColumn sul welcomeTopic hanno il seguente formato
//   {"username":"nomeUtenteRiconosciuto", "response-topic":"nomeTopic"}
//   il producer mqtt può quindi fare una publish sul topic richiesto dall'InfoColumns
//   il formato della risposta sul responseTopic è del tipo [{"tag":"weather", "args":["city":["catania", ...]]}]
//   => PATTERN REQUEST-REPLY su mqtt
func RequestReplyRoutine(client mqtt.Client, msg mqtt.Message) {
	request := &RequestBody{}
	if err := json.Unmarshal(msg.Payload(), request); err != nil {
		// rispondi da qualche parte in mqtt che la richiesta è malformata
		log.Printf("[IDENTITY] Received bad request from %s: %s [error: %s]", msg.Topic(), msg.Payload(), err.Error())
		return
	}

	reply := repository.ListInterestsByUser(request.Username)

	toSend, _ := json.Marshal(reply)

	client.Publish(request.ResponseTopic, 0, false, toSend).Wait()
}
