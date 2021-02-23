import redis
import json


class Repository:
    def __init__(self):
        self.redis = redis.Redis(host="redis")

    # Returns True/False
    def push(self, key, value):
        return self.redis.set(key, json.dumps(value))

    # Returns the read string
    def pull(self, key):
        return self.redis.get(key)


class InterestRepository(Repository):
    def _get_key(self, tag, args):
        redis_key = tag
        for arg in args:
            redis_key += f"|{args}"
        return redis_key

    def create(self, tag, args, value):
        return self.push(self._get_key(tag, args), value)

    def read(self, tag, args):
        return self.pull(self._get_key(tag, args))
