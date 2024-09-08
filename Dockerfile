# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /app
ADD . /app/
RUN go build -o ./cmd/go-chat-app/main.go .
EXPOSE 8000
ENTRYPOINT ["./cmd/go-chat-app/main.go"]