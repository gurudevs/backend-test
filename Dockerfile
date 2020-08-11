FROM golang

ADD . /go/src/github.com/ferkze/backend-test

WORKDIR /go/src/github.com/ferkze/backend-test

RUN go get -d -v

RUN go install github.com/ferkze/backend-test

ENTRYPOINT /go/bin/backend-test