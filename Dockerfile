FROM golang:1.10.4 AS builder

ARG PROJECT

COPY . /go/src/github.com/niima/${PROJECT}

WORKDIR /go/src/github.com/niima/${PROJECT}

RUN go get -v

RUN CGO_ENABLED=0 go build -v -o /go/bin/app

FROM alpine:3.7 AS app

ARG PROJECT

COPY --from=builder /go/bin/app /

RUN chmod +x /app

ENTRYPOINT /app