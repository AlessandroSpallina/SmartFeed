import requests
import json


# TODO: needs provider configuration, this static conf is ugly
class ProviderApi:
    def __init__(self):
        self._provider_host = "feed-provider"
        self._provider_port = 3030

    def produce(self, key, args, data):
        payload = {'tag': key, 'args': args, 'data': json.dumps(data)}
        requests.post(f"http://{self._provider_host}:{self._provider_port}/produce", data=json.dumps(payload))
