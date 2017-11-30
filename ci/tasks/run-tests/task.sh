#!/bin/bash -exu

ROOT=${PWD}
GOPATH=${PWD}/go
PATH=${GOPATH}/bin:$PATH

function main() {
  go get github.com/onsi/ginkgo/ginkgo

  mkdir -p "${GOPATH}/src/github.com/pivotal-cf"
  ln -s "${ROOT}/staticfile-resource" "${GOPATH}/src/github.com/pivotal-cf/staticfile-resource"
  ginkgo -r "${GOPATH}/src/github.com/pivotal-cf/staticfile-resource"
}

main
