# syntax=docker/dockerfile:1

FROM golang:1.25.6

WORKDIR /websocket

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /img-ws

ENV WS_ADDRESS=":8081"

EXPOSE 8081

CMD ["/img-ws"]