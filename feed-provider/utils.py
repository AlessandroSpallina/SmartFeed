import os


def get_config_from_env():
    return {
        "HTTP_SERVER_PORT": int(os.environ.get('PROVIDER_HTTP_SERVER_PORT', 3030)),
        "MQTT_BROKER_HOST": os.environ.get('PROVIDER_MQTT_BROKER_HOST', 'localhost'),
        "MQTT_BROKER_PORT": int(os.environ.get('PROVIDER_MQTT_BROKER_PORT', 1883)),
        "MQTT_PRODUCER_ID": os.environ.get('PROVIDER_MQTT_PRODUCER_ID', 'unnamed-node'),
        "MQTT_WELCOME_TOPIC": os.environ.get('PROVIDER_MQTT_WELCOME_TOPIC', 'provider/welcome')
    }
