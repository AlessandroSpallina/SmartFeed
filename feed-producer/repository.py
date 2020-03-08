import redis
import json

class Repository:

    def __init__(self, cities):
        self.cities = cities
        self.redis = redis.Redis(host="redis")

    def _key(self, feed):
        return self.cities + feed

    def push_repo(self, feed, send_list):
        self.redis.set(self._key(feed), json.dumps(send_list))
        print(send_list)