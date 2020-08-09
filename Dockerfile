FROM golang

ADD . /go/src/github.com/ferkze/backend-test

RUN go get -u github.com/gorilla/mux

RUN go get github.com/robfig/cron/v3@v3.0.0

RUN go install github.com/ferkze/backend-test

ENTRYPOINT /go/bin/backend-test