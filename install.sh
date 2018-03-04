#!/bin/bash

GOPATH=$(go env GOPATH)
REPO_ROOT="$GOPATH/src/github.com/tamalsaha/kubectl-plugin-demo"

pushd $REPO_ROOT

mkdir -p ~/.kube/plugins/tamal
cp plugin.yaml ~/.kube/plugins/tamal/plugin.yaml
go build -v
mv kubectl-plugin-demo ~/.kube/plugins/tamal/kubectl-plugin-demo

popd
