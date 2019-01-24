FROM golang:latest

ARG SOURCE_LOCATION=/go/src/app
WORKDIR ${SOURCE_LOCATION}
COPY . ${SOURCE_LOCATION}
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go
CMD ["./app"]