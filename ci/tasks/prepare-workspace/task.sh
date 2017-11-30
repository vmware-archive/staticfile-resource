#!/bin/bash -exu

ROOT=${PWD}
GOPATH=${PWD}/go
PATH=${GOPATH}/bin:$PATH

function main() {
  mkdir -p "${GOPATH}/src/github.com/christianang"
  ln -s "${ROOT}/staticfile-resource" "${GOPATH}/src/github.com/christianang/staticfile-resource"
  pushd "${GOPATH}/src/github.com/christianang/staticfile-resource" > /dev/null
    ./scripts/build
  popd > /dev/null

  cp -R ${ROOT}/staticfile-resource/* "${ROOT}/workspace"
}

main
