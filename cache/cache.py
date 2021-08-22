import json
from typing import List, Optional

import redis

r = redis.Redis(host='cache')

# Cache entries for 1 hour.
caching_time = 60 * 60


def put(provider: str, query: str, result: dict):
	key = f'{provider}:{query}'
	# Insert provider into results for bulk queries.
	result['provider'] = provider

	value = json.dumps(result)
	print(f'Cache value: {value}')
	r.setex(name=key, time=caching_time, value=value)


def get(provider: str, query: str) -> Optional[dict]:
	key = f'{provider}:{query}'
	result = r.get(name=key)
	if result:
		return json.loads(result)
	else:
		return None


def get_all(query: str) -> List[dict]:
	res = []
	for name in ['g2a', 'steam', 'hb']:
		result = get(name, query)
		if result:
			res.append(result)
	return res
