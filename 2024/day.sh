#/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./day.sh <day number>"
fi

day=$(printf "%02d" "$1")

if [ -d "$day" ]; then
    echo "Error: directory '$day' already exists"
    exit 1
fi

cp -r _template "$day"
echo "Created folder: '$day' from '_template'"
