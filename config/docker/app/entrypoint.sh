#!/bin/sh

# Start the reflex script in the background
echo "Starting reflex"
reflex -r "\\.go$" -s -- sh -c 'echo "change detected!" && go run /var/www/cmd/server/server.go'

wait