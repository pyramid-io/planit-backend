# Start from the latest golang base image
FROM golang:latest

RUN go install github.com/cespare/reflex@latest

# Set the Current Working Directory inside the container
WORKDIR /var/www

# Copy go mod and sum files
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "src/main.go"]
