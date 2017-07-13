#!/bin/sh

set -e -x

cp -r go-url-preview $GOPATH/src/github.com/EMC-Dojo/

cd $GOPATH/src/github.com/EMC-Dojo/

glide up
ginkgo -r .
