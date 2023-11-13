#!/bin/sh

pushd operator
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager ./cmd/manager
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o gke_addon_poststart ./cmd/gke_addon_poststart
popd

make -f Makefile.fifteen push-operator
