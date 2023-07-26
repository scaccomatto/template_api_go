lFROM golang:1.18-buster AS dependencies

WORKDIR /app

COPY api ./api
COPY cmd ./cmd
COPY config ./config
COPY internal ./internal
COPY vendor ./vendor
COPY go.mod ./
COPY go.sum ./

FROM dependencies AS build
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/

FROM alpine:latest
RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add --no-cache \
                curl \
                wget \
                bash \
                openssh \
                npm \
                git \
                jq

WORKDIR /app
COPY --from=build /app/app .
COPY --from=build /app/config/app_config.yaml ./config/app_config.yaml

EXPOSE 80
CMD ["./app"]