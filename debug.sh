#!/bin/bash

set -e

go build -o setupjunkie

docker ps > /dev/null 2>&1
if [ $? -eq 0 ]; then
    DOCKER="docker"
else
    DOCKER="sudo docker"
fi

$DOCKER build -t test-setupjunkie .
$DOCKER run -it test-setupjunkie /bin/bash
