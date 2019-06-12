#FROM scratch
FROM alpine:latest 
MAINTAINER mlycore <maxwell92@126.com>
ADD echoserver /echoserver
ADD Dockerfile /Dockerfiler
ENTRYPOINT ["/echoserver"]

