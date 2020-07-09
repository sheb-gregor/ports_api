# Compile stage
FROM golang:alpine AS build-env
RUN apk add --no-cache git

ARG CONFIG="master"

#ENV GOPROXY=direct
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPRIVATE=gitlab.com

WORKDIR /service
ADD . .
COPY ./env/${CONFIG}.config.yaml /config.yaml
RUN go mod download && ./build.sh /app

# Final stage
FROM alpine:3.7

# Port 8080 belongs to our application
EXPOSE 8080

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache ca-certificates bash

WORKDIR /

COPY --from=build-env /app /
COPY --from=build-env /config.yaml /

# Run delve
CMD ["/app", "serve"]
