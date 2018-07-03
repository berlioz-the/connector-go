FROM berliozcloud/golang-1.10

RUN go get -u github.com/apache/thrift/lib/go/thrift

ADD . /go/src/app
RUN dep ensure

RUN go build -o /app/main examples/hello/*
