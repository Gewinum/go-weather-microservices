FROM golang:1.23 as build-env

WORKDIR /app

RUN mkdir -p /app/common
COPY --from=common . /app/common

WORKDIR /app/main

COPY --from=main go.mod go.sum ./
RUN go mod download

COPY --from=main . .
RUN go build -o /out/app ./cmd/app/main.go

FROM debian:buster

WORKDIR /
COPY --from=build-env /out /
CMD ["./app"]