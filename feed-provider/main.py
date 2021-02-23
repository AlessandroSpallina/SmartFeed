import os
from mqtt import MqttHandler


def get_config_from_env():
    return {
        "HTTP_SERVER_PORT": int(os.environ['PROVIDER_HTTP_SERVER_PORT']),
        "MQTT_BROKER_HOST": os.environ['PROVIDER_MQTT_BROKER_HOST'],
        "MQTT_BROKER_PORT": int(os.environ['PROVIDER_MQTT_BROKER_PORT']),
        "MQTT_PRODUCER_ID": os.environ['PROVIDER_MQTT_PRODUCER_ID'],
        "MQTT_WELCOME_TOPIC": os.environ['PROVIDER_MQTT_WELCOME_TOPIC']
    }


if __name__ == '__main__':
    config = get_config_from_env()

    mqtt = MqttHandler(config)
    mqtt.init()
    mqtt.start()
