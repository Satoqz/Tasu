FROM alpine:latest
LABEL author="satoqz"

RUN apk update && apk add libstdc++

COPY bf /usr/local/bin/
COPY run.sh /var/run
