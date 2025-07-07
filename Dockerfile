FROM golang:1.24.2 AS  build-stage

WORKDIR /app

COPY cmd ./cmd
COPY go.mod go.sum ./
COPY api ./api
COPY internal ./internal
COPY conf ./conf

RUN CGO_ENABLED=0 GOOS=linux go build -o template-app ./cmd/


# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /app/template-app /app/template-app
COPY --from=build-stage /app/conf/config.yaml /app/conf/config.yaml

EXPOSE 8081

USER nonroot:nonroot


#RUN apk add --no-cache \
#                curl \
#                bash \
#                openssh \
#                git

ENTRYPOINT ["./template-app"]
