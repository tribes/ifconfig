# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /ifconfig

##
## Deploy
##
FROM alpine

WORKDIR /

COPY --from=build /ifconfig /ifconfig

EXPOSE 80

USER nobody:nobody

ENTRYPOINT ["/ifconfig"]
