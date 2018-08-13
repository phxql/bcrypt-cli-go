#!/usr/bin/env bash

mkdir bin

GOOS=darwin GOARCH=amd64 go build -o bin/bcrypt-cli-go-darwin
GOOS=windows GOARCH=amd64 go build -o bin/bcrypt-cli-go-windows
GOOS=linux GOARCH=amd64 go build  -o bin/bcrypt-cli-go-linux
