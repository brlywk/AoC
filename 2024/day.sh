#!/bin/zsh

day=$1

if [ "$day" -z ]; then
    echo "No day specified"
    exit 1
else
    mkdir "$day"
    cd "$day"
    cargo init --name "day_$day"
    exit 0
fi
