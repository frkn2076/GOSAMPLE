FROM golang:latest

ENV GOPATH /go

COPY ./config /go/src/app/GoSample/config
COPY ./controllers /go/src/app/GoSample/controllers
COPY ./db /go/src/app/GoSample/db
COPY ./infra /go/src/app/GoSample/infra
COPY ./logger /go/src/app/GoSample/logger
COPY ./middleware /go/src/app/GoSample/middleware
COPY ./router /go/src/app/GoSample/router
COPY ./main.go /go/src/app/GoSample/main.go

WORKDIR /go/src/app/GoSample

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["GoSample"]