FROM nginx:latest

ENV APP_DOMAIN="qurt.tech"

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

COPY ./landing /usr/share/nginx/html
COPY ./assets /usr/share/nginx/html/assets
COPY ./tour /usr/share/nginx/html/tour

COPY ./nginx/99-autoreload.sh /docker-entrypoint.d/99-autoreload.sh

RUN chmod +x /docker-entrypoint.d/99-autoreload.sh;

RUN find /usr/share/nginx/html -name "*.html" -o -name "*.css" -o -name "*.js" | xargs gzip -k -f

EXPOSE 80 443
