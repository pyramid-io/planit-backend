#!/bin/sh
envsubst '\$APP_PORT \$SERVER_NAME' < /etc/nginx/conf.d/app.conf.template > /etc/nginx/conf.d/app.conf
nginx -g 'daemon off;'