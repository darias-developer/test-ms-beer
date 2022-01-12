#!/bin/bash

docker run \
 -e DB_SOURCE='mongodb+srv://user_free_cluster:9UjGcHJJgIu5vTts@cluster0.o1qf0.mongodb.net/beer-test?retryWrites=true&w=majority' \
 -e PORT='8080' \
 -e LOG_PATH='/app/logs/test-ms-beer.log' \
 -e ACCESS_KEY='e59979e596dc86b3aaea9f1727e41416' \
 --publish 8080:8080 \
 -v logs:/app/logs/test-ms-beer.log \
 test-ms-beer:v1.0 &