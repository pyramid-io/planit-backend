# Start from the latest golang base image
FROM golang:latest

RUN go install github.com/cespare/reflex@latest
# RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set the Current Working Directory inside the container
WORKDIR /var/www

# Expose port 8080 to the outside world
EXPOSE 8080
# EXPOSE 2345

ADD . /var/www
RUN go mod tidy & go mod verify



