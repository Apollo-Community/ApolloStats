#!/bin/bash

FILES="README.md src/main.go"
NEWVERSION=$1
OLDVERSION=$(grep -o "VERSION = \".*\"" src/main.go | grep -o "[0-9]\.[0-9]")
echo "old ver: '$OLDVERSION'"
echo "new ver: '$NEWVERSION'"

function commit {
    echo "Committing..."
    git commit $FILES -m "Increasing version to v$NEWVERSION."
    git tag -a "v$NEWVERSION"
    exit
}

function abort {
    echo "Aborting..."
    exit
}

sed -i "s/$OLDVERSION/$NEWVERSION/g" $FILES
git diff .

while true; do
    read -p "Do wish to commit and tag this new version (y/n)? " input
    case $input in
        [Yy]*) commit;;
        [Nn]*) abort;;
        *) echo "Please answear yes or no.";;
    esac
done
