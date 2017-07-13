#!/usr/bin/env bash

set -e -x

cp -r go-url-preview $GOPATH/src/github.com/EMC-Dojo/

pushd $GOPATH/src/github.com/EMC-Dojo/
  glide up
  ginkgo -r .
popd