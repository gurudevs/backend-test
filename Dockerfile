FROM golang

ADD . /go/src/github.com/ferkze/backend-test

RUN go get -u github.com/gorilla/mux

RUN go install github.com/ferkze/backend-test

ENTRYPOINT /go/bin/backend-test