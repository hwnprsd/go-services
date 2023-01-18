# !/bin/bash

docker context use default
docker compose build
docker compose push
docker context use opl 
