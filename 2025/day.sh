#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./day.sh <day number>"
fi

day=$(printf "%02d" "$1")

cargo new day$day
echo "Created day$day"
