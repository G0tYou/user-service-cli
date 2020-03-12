# user-service-cli/Dockerfile

# Build the user-service-cli
FROM golang:alpine AS build

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /go/src/app
WORKDIR /go/src/app

ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service-cli

# Copy the newly built user-service-cli to Alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates

RUN mkdir /app
WORKDIR /app
COPY data .
COPY --from=build /go/src/app/user-service-cli .