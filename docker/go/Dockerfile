FROM golang:1.23.0-alpine3.20 AS build-env

WORKDIR /app

RUN mkdir -p /app/common
COPY --from=common . /app/common

WORKDIR /app/main

COPY --from=main go.mod go.sum ./
RUN go mod download

COPY --from=main . .
RUN go build -o /out/app ./cmd/app/main.go

FROM alpine:3.20

WORKDIR /
COPY --from=build-env /out /
CMD ["./app"]