FROM berliozcloud/golang-1.10

ADD . /go/src/app
RUN dep ensure

RUN go build -o /app/main examples/hello/*
