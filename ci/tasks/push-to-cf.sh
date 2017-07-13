#!/usr/bin/env bash

set -e -x

cf api http://api.$CF_ENDPOINT --skip-ssl-validation
cf auth $CF_USERNAME "${CF_PASSWORD}"
cf target -o $CF_ORG -s $CF_SPACE

cd go-url-preview
cf push