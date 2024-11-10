
#!/bin/zsh
day=$1

if [ -z "$day" ]; then
    echo "No day specified"
    exit 1
else
    mkdir "$day"
    cd "$day"
    cargo init --name "day$day"
    echo "common = { path = \"../common\" }" >> Cargo.toml
    mkdir inputs
    touch inputs/input.txt inputs/test.txt

    # workspace modifications
    cd ..
    sed -i '' 's/\["common"\]/["common", "day'$day'"]/' Cargo.toml
    exit 0
fi
