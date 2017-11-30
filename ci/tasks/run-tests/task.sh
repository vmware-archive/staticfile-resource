#!/bin/bash -exu

ROOT=${PWD}
GOPATH=${PWD}/go
PATH=${GOPATH}/bin:$PATH

function main() {
  go get github.com/onsi/ginkgo/ginkgo
  go get github.com/golang/dep/cmd/dep

  mkdir -p "${GOPATH}/src/github.com/christianang"
  ln -s "${ROOT}/staticfile-resource" "${GOPATH}/src/github.com/christianang/staticfile-resource"
  dep ensure "${GOPATH}/src/github.com/christianang/staticfile-resource"
  ginkgo -r "${GOPATH}/src/github.com/christianang/staticfile-resource"
}

main
