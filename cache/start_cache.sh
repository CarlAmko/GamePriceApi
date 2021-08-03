#!/usr/bin/env bash

docker run -p 6379:6379 -d --name=game_api_cache redis:latest
