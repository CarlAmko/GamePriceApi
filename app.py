#!/usr/bin/env python3
from urllib.parse import quote_plus

from flask import Flask, request, abort, Response, jsonify

from api.spiders.launcher import run
from cache import cache

app = Flask(__name__)


@app.route("/")
def prices():
	provider = request.args.get('provider', None)

	query = request.args.get('query', None)
	if query is None:
		abort(Response(status=400, response='Failed to provide query param.'))

	query = quote_plus(query)
	if provider:
		run(query=query, spider=provider)
		results = cache.get(query=query, provider=provider)
	else:
		run(query=query)
		results = cache.get_all(query=query)

	print(results)
	return jsonify(results)


if __name__ == '__main__':
	app.run(debug=True)
