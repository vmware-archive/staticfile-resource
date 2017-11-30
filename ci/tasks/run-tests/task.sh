#!/bin/bash -exu

ROOT=${PWD}
GOPATH=${PWD}/go
PATH=${GOPATH}/bin:$PATH

function main() {
  go get github.com/onsi/ginkgo/ginkgo
  go get github.com/golang/dep/cmd/dep

  mkdir -p "${GOPATH}/src/github.com/pivotal-cf"
  ln -s "${ROOT}/staticfile-resource" "${GOPATH}/src/github.com/pivotal-cf/staticfile-resource"
  pushd "${GOPATH}/src/github.com/pivotal-cf/staticfile-resource" > /dev/null
    dep ensure
    ginkgo -r .
  popd > /dev/null
}

main
