FROM golang:1.8.3-alpine3.6

RUN apk update && \
    apk add git && \
    apk add glide && \
    mkdir -p $GOPATH/src/github.com/EMC-Dojo

RUN go get github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI