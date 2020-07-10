# Compile stage
FROM golang:alpine AS build-env
RUN apk add --no-cache git bash

ARG SERVICE=""
ARG CONFIG="master"

#ENV GOPROXY=direct
#ENV GOPRIVATE=gitlab.com
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /service
ADD . .
COPY ./env/${CONFIG}.${SERVICE}_cfg.yaml /config.yaml
RUN ./build.sh ${SERVICE} /app

# Final stage
FROM alpine:3.7

# Port 8080 belongs to our application
EXPOSE 8080

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache ca-certificates bash

WORKDIR /

COPY --from=build-env /app /
COPY --from=build-env /config.yaml /

CMD ["/app"]
