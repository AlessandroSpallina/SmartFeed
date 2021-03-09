import utils
from marshmallow_dataclass import dataclass
from typing import List
import json
import marshmallow

import repository
from mqtt import MqttHandler


# json to python object the easy way: https://stackoverflow.com/a/54792440
@dataclass
class MqttQuery:
    """
    Accept something like this
    {
        'tag': 'tagname',
        'args' : ['value1', 'value2']
    }
    """
    tag: str
    args: List[str]


@dataclass
class MqttMessage:
    """
    Messaggio del tipo
    {query: [mqttquery, ...], response_topic: stringa}
    """
    query: List[MqttQuery]
    response_topic: str


def make_message_handler():
    db = repository.InterestRepository()

    # this function accept a dict like this: {"topic": <topic>: "message": <message>}
    # returns a dict like {"topic":<topic>: "message": <message>} that have to be sent by the MqttHandler object
    def received_message_handler(received_dict):
        to_return = None

        try:
            received_dict["message"] = json.loads(received_dict["message"])
            received_dict["message"] = MqttMessage.Schema().load(received_dict["message"])
            to_publish = []
            for q in received_dict["message"].query:
                to_publish.append({'query': q.__dict__, 'data': db.read(q.tag, q.args)})
            to_return = {"topic": received_dict["message"].response_topic, "message": json.dumps(to_publish)}
        except marshmallow.exceptions.ValidationError:
            print("Received bad message format (ValidationError)")
        except json.decoder.JSONDecodeError:
            print("Received bad message format (JSONDecodeError)")

        return to_return

    return received_message_handler


if __name__ == '__main__':
    config = utils.get_config_from_env()

    mqtt = MqttHandler(config, [config['MQTT_WELCOME_TOPIC']], make_message_handler())
    mqtt.init()
    mqtt.start()
