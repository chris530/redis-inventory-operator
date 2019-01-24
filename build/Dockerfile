FROM alpine:3.8

RUN apk upgrade --update --no-cache

USER nobody

ADD build/_output/bin/redis-inventory-operator /usr/local/bin/redis-inventory-operator
