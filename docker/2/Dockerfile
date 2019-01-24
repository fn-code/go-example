FROM golang:latest AS builder

ARG SOURCE_LOCATION=/go/src/app
WORKDIR ${SOURCE_LOCATION}
COPY . ${SOURCE_LOCATION}
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apps ./cmd/main.go

FROM alpine:latest
ARG BUILDER_SOURCE=/go/src/app
WORKDIR /app
COPY --from=builder ${BUILDER_SOURCE}/ .
ENTRYPOINT [ "./apps" ]