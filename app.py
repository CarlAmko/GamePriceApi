#!/usr/bin/env python3

from flask import Flask, request, abort, Response, jsonify

from api.spiders.launcher import run

app = Flask(__name__)


@app.route("/prices")
def prices():
    provider = request.args.get('provider', None)

    query = request.args.get('query', None)
    if query is None:
        abort(Response(status=400, response='Failed to provide query param.'))

    if provider:
        run(spider=provider, query=query)

    return jsonify()


if __name__ == '__main__':
    app.run()
