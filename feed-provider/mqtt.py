import paho.mqtt.client as mqtt
from marshmallow_dataclass import dataclass
import marshmallow
from typing import List
import json


@dataclass
class MqttQuery:
    '''
    Accept something like this
    {
        'tag': 'tagname',
        'args' : {
                    'arg1': 'value1',
                    'arg2': 'value2'
                 }
    }
    '''
    tag: str
    args: dict


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
            message = MqttMessage.Schema().load(msg.payload)
        except marshmallow.exceptions.ValidationError:
            print("Received bad message format (ValidationError)")


    def init(self):
        self._client = mqtt.Client()
        self._client.on_connect = self._on_connect
        self._client.on_message = self._on_message

    def start(self):
        self._client.connect(self._broker_host, self._broker_port, 60)
        self._client.loop_forever()
