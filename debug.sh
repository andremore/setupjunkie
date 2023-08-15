#!/bin/bash

set -e

go build -o setupjunkie
sudo docker build -t test-setupjunkie .
sudo docker run -it test-setupjunkie /bin/bash
