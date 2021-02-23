from mqtt import MqttHandler
import utils


if __name__ == '__main__':
    config = utils.get_config_from_env()

    mqtt = MqttHandler(config)
    mqtt.init()
    mqtt.start()
