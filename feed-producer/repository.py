import redis
import json

class Repository:

    def __init__(self):
        self.redis = redis.Redis(host="redis")
        
    def push(self, key, value):
        self.redis.set(key, json.dumps(value))
