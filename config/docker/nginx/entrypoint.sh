#!/bin/sh
envsubst '\$SERVER_PORT' < /etc/nginx/conf.d/app.conf.template > /etc/nginx/conf.d/app.conf
nginx -g 'daemon off;'