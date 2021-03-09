import paho.mqtt.client as mqtt
import paho.mqtt.publish as publish

import time


class MqttHandler:
    """
    Accetta:
        * una lista di topic_to_subscribe
        * una funzione che gestirà i messaggi ricevuti

    Quando riceve un messaggio da un topic passa un dizionario {"topic": <topic>: "message": <message>}
    alla funzione che implementa la logica di gestione del messaggio
    """
    def __init__(self, config, topics_to_subscribe,  received_message_handler):
        self._broker_host = config['MQTT_BROKER_HOST']
        self._broker_port = config['MQTT_BROKER_PORT']
        self._id = config['MQTT_PRODUCER_ID']

        self._topics_to_subscribe = topics_to_subscribe
        self._received_message_handler = received_message_handler

        self._client = None

    def _on_connect(self, client, userdata, flags, rc):
        print("Connected with result code " + str(rc))

        # Subscribing in on_connect() means that if we lose the connection and
        # reconnect then subscriptions will be renewed.
        for topic in self._topics_to_subscribe:
            print(f"Subscribed to {topic}")
            self._client.subscribe(topic)

    # The callback for when a PUBLISH message is received from the server.
    def _on_message(self, client, userdata, msg):
        print(msg.topic + " " + str(msg.payload))

        ret = self._received_message_handler({"topic": msg.topic, "message": msg.payload})
        if ret is not None:
            publish.single(ret["topic"], ret["message"], hostname=self._broker_host)

    def init(self):
        self._client = mqtt.Client()
        self._client.on_connect = self._on_connect
        self._client.on_message = self._on_message

    def start(self):
        try:
            self._client.connect(self._broker_host, self._broker_port, 60)
            self._client.loop_forever()
        except ConnectionRefusedError:
            print("Connection refused, retrying in 10 sec")
            time.sleep(10)
            self.start()
