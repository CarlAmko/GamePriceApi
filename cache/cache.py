import json

import redis

r = redis.Redis()


def put(provider: str, query: str, result: dict):
    one_hour = 3600
    key = f'{provider}:{query}'
    r.setex(name=key, time=one_hour, value=json.dumps(result))


def get(provider: str, query: str) -> dict:
    key = f'{provider}:{query}'
    return json.loads(r.get(name=key))
