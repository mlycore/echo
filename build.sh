#! /bin/bash

GOOS=linux GOARCH=amd64 go build echo.go
# docker build -t echo:`date +%F` .
#if [ -e echo ] 
#then
#    rm echo
#fi
[ -e echo ] && rm echo