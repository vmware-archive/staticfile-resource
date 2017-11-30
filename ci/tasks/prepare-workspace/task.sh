#!/bin/bash -exu

ROOT=${PWD}
GOPATH=${PWD}/go
PATH=${GOPATH}/bin:$PATH

function main() {
  mkdir -p "${GOPATH}/src/github.com/pivotal-cf"
  ln -s "${ROOT}/staticfile-resource" "${GOPATH}/src/github.com/pivotal-cf/staticfile-resource"
  pushd "${GOPATH}/src/github.com/pivotal-cf/staticfile-resource" > /dev/null
    ./scripts/build
  popd > /dev/null

  cp -R ${ROOT}/staticfile-resource/* "${ROOT}/workspace"
}

main
