# syntax=docker/dockerfile:1
# check=error=true
FROM golang:1.27 AS build-stage

EXPOSE 8080

RUN mkdir /app
WORKDIR /app

# Download Go modules
COPY . ./
RUN go mod download

# le starting point
COPY ./docker/backend/entrypoint /
# just in case it's missing executable bit
RUN chmod +x /entrypoint

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./backend
