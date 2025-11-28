#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./day.sh <day number>"
    exit 1
fi

day=$(printf "%02d" "$1")

cargo init "day$day"
cp _template/main.rs "day$day/src/main.rs"
cp _template/test.txt "day$day/src/test.txt"
cp _template/data.txt "day$day/src/data.txt"
mv "day$day" "$day"

echo "Created $day"
