FROM golang:1.24
WORKDIR /src

COPY main_native.go main_native.go 
COPY main_wasm.go main_wasm.go
COPY main.go main.go

COPY internal/ internal/

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod tidy
RUN go build -o ./bin/qurtc .
RUN GOOS=js GOARCH=wasm go build -o ./bin/qurtc.wasm .


FROM nginx:latest

ARG MAIN_PATH="./website"
ENV APP_DOMAIN="qurt.tech"

COPY ${MAIN_PATH}/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ${MAIN_PATH}/landing /usr/share/nginx/html
COPY ${MAIN_PATH}/assets /usr/share/nginx/html/assets
COPY ${MAIN_PATH}/tour /usr/share/nginx/html/tour
COPY ${MAIN_PATH}/nginx/99-autoreload.sh /docker-entrypoint.d/99-autoreload.sh

COPY --from=0 /src/bin/* /usr/share/nginx/html/assets/wasm/
# COPY --from=0 /src/bin/qurtc.wasm /usr/share/nginx/html/assets/wasm/

RUN chmod +x /docker-entrypoint.d/99-autoreload.sh;

RUN find /usr/share/nginx/html -name "*.html" -o -name "*.css" -o -name "*.js" | xargs gzip -k -f

EXPOSE 80 443
