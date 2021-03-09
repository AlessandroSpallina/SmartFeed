import sys
import os
import csv
import time

sys.path.insert(0, '../feed-provider')
from mqtt import MqttHandler
import utils

LOG_FILE = "mqtt_messages.csv"


def received_message_handler(received_dict):
    CSV_HEADER = ['timestamp', 'topic', 'byte_received']
    append_csv_header = False

    if not os.path.exists(LOG_FILE):
        append_csv_header = True

    with open(LOG_FILE, mode="a", newline='', encoding='utf-8') as f:
        csv_writer = csv.writer(f, delimiter=";")
        if append_csv_header:
            csv_writer.writerow(CSV_HEADER)
        csv_writer.writerow([time.time(), received_dict["topic"], len(received_dict["message"])])


def main():
    conf = utils.get_config_from_env()

    mqtt = MqttHandler(conf, ['identity/#', 'provider/#'], received_message_handler)
    mqtt.init()
    mqtt.start()


if __name__ == "__main__":
    main()
