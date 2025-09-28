FROM nginx:latest

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

COPY ./landing /usr/share/nginx/html
COPY ./assets /usr/share/nginx/html/assets
COPY ./editor /usr/share/nginx/html/editor

COPY ./nginx/99-autoreload.sh /docker-entrypoint.d/99-autoreload.sh

RUN chmod +x /docker-entrypoint.d/99-autoreload.sh

EXPOSE 80 443
