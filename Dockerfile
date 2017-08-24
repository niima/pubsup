FROM alpine:3.5

RUN apk --no-cache add ca-certificates


COPY ./sync-service /
COPY ./config.json /

WORKDIR /

ENTRYPOINT /sync-service