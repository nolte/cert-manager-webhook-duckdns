#!/usr/bin/env bash
mkdir -p __main__/hack
curl -sfL https://github.com/kubernetes-sigs/kubebuilder/releases/download/v2.3.2/kubebuilder_2.3.2_darwin_amd64.tar.gz | tar xvz --strip-components=1 -C __main__/hack