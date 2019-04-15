FROM golang:latest

ARG SOURCE_LOCATION=/go/src/app

WORKDIR ${SOURCE_LOCATION}

ADD . ${SOURCE_LOCATION}

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 9090

CMD ["./main"]