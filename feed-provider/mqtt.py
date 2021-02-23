import paho.mqtt.client as mqtt
import paho.mqtt.publish as publish
from marshmallow_dataclass import dataclass
from typing import List
import json
import marshmallow
import time

import repository


@dataclass
class MqttQuery:
    '''
    Accept something like this
    {
        'tag': 'tagname',
        'args' : ['value1', 'value2']
    }
    '''
    tag: str
    args: List[str]


@dataclass
class MqttMessage:
    '''
    Messaggio del tipo
    {query: [mqttquery, ...], response_topic: stringa}
    '''
    query: List[MqttQuery]
    response_topic: str


class MqttHandler:
    def __init__(self, config):
        self._broker_host = config['MQTT_BROKER_HOST']
        self._broker_port = config['MQTT_BROKER_PORT']
        self._welcome_topic = config['MQTT_WELCOME_TOPIC']
        self._id = config['MQTT_PRODUCER_ID']

        self._client = None
        self._db = None

    def _on_connect(self, client, userdata, flags, rc):
        print("Connected with result code " + str(rc))

        # Subscribing in on_connect() means that if we lose the connection and
        # reconnect then subscriptions will be renewed.
        self._client.subscribe(self._welcome_topic)

    # The callback for when a PUBLISH message is received from the server.
    def _on_message(self, client, userdata, msg):
        print(msg.topic + " " + str(msg.payload))
        # Messaggio del tipo
        # {query: [mqttquery, ...], response-topic: stringa}
        try:
            msg.payload = json.loads(msg.payload)
            message = MqttMessage.Schema().load(msg.payload)
            # raccogliere le risposte delle query
            # comporre un messaggio di risposta adeguato
            # inviare sul giusto topic
            to_publish = []
            for q in message.query:
                to_publish.append({'query': q.__dict__, 'data': self._db.read(q.tag, q.args)})

            # print(f"{self._db.read(message.query[0].tag, message.query[0].args)}")
            publish.single(message.response_topic, json.dumps(to_publish), hostname=self._broker_host)

            # print(f"{message.response_topic}, {message.query[0].tag}")

        except marshmallow.exceptions.ValidationError:
            print("Received bad message format (ValidationError)")
        except json.decoder.JSONDecodeError:
            print("Received bad message format (JSONDecodeError)")

    def init(self):
        self._client = mqtt.Client()
        self._client.on_connect = self._on_connect
        self._client.on_message = self._on_message

    def start(self):
        try:
            self._client.connect(self._broker_host, self._broker_port, 60)
            self._db = repository.InterestRepository()
            self._client.loop_forever()
        except ConnectionRefusedError:
            print("Connection refused, retrying in 10 sec")
            time.sleep(10)
            self.start()
