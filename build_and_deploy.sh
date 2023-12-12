#!/bin/bash
if [!-d ./build]; then
    echo "Create 'build' folder"
    mkdir -p ./build;
fi;

cp main.exe ./build
cp -r config ./build