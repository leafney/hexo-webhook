FROM alpine:3.7
MAINTAINER leafney "babycoolzx@126.com"

RUN apk add --no-cache git pcre nginx supervisor && \
    mkdir -p /etc/supervisor.d && \
    rm -rf /var/cache/apk/*

COPY ./supervisor_bloghook.ini /etc/supervisor.d/supervisor_bloghook.ini
COPY ./nginx_bloghook.conf /etc/nginx/conf.d/nginx_bloghook.conf
COPY ./build.sh ./main /app/

VOLUME ["/app"]

EXPOSE 80 8080

CMD ["supervisord","-c","/etc/supervisord.conf"]