FROM golang:latest

ENV GOPATH /go
RUN go env -w GO111MODULE=on

COPY . /go/src/app/GoSample

WORKDIR /go/src/app/GoSample

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["GoSample"]