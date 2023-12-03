#!/bin/bash

wails3 generate bindings -d frontend

cd "frontend"

npm install

npm run build

if [ $? -ne 0 ]; then
    echo "Error: npm run build failed."
    exit $?
fi

cd ..

go mod tidy
go run .
