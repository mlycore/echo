#! /bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o echoserver echo.go
docker build -t echo:`date +%F` .
#if [ -e echo ] 
#then
#    rm echo
#fi
[ -e echoserver ] && rm echoserver
