FROM golang:1.8.3-alpine3.6

RUN apk update && \
    apk add git && \
    mkdir -p $GOPATH/src/github.com/EMC-Dojo

RUN go get github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
RUN go get github.com/onsi/gomega         # fetches the matcher library