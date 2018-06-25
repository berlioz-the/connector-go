FROM berliozcloud/golang-1.10

ADD . /go/src/app
RUN dep ensure

RUN ls -la /go/src/
RUN ls -la /go/src/app

RUN go build -o /app/main examples/hello/*
