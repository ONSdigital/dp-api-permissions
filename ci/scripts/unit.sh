#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-authorisation
  make test
popd
