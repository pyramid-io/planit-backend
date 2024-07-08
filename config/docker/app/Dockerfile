FROM golang:latest

WORKDIR /var/www

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
# RUN go install github.com/go-delve/delve/cmd/dlv@latest

ADD . /var/www


RUN go mod download
RUN go mod tidy
RUN go mod verify

RUN swag init -g cmd/server/server.go


EXPOSE 8080
# EXPOSE 2345


