#!/bin/bash -exu

ROOT=${PWD}
GOPATH=${PWD}/go
PATH=${GOPATH}/bin:$PATH

function main() {
  go get github.com/golang/dep/cmd/dep

  mkdir -p "${GOPATH}/src/github.com/christianang"
  ln -s "${ROOT}/staticfile-resource" "${GOPATH}/src/github.com/christianang/staticfile-resource"
  pushd "${GOPATH}/src/github.com/christianang/staticfile-resource" > /dev/null
    dep ensure
    ./scripts/build
  popd > /dev/null

  cp -R ${ROOT}/staticfile-resource/* "${ROOT}/workspace"
}

main
