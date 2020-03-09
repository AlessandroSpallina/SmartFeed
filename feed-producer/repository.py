import redis
import json

class Repository:

    def __init__(self, cities):
        self.cities = cities.lower()
        self.redis = redis.Redis(host="redis")

    def _key(self, feed):
        return feed + '|' + self.cities
        

    def push_repo(self, feed, send_list):
        self.redis.set(self._key(feed), json.dumps(send_list))
        #print(send_list)