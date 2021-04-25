#!/usr/bin/env bash

set -ex

version=0.0.1
host_name=registry.terraform.io
namespace=selectel
name=vscale
target0=darwin_amd64

provider_dir="${HOME}/.terraform.d/plugins/${host_name}/${namespace}/${name}/${version}/${target0}"
provider_bin="terraform-provider-${name}_v${version}"

go mod tidy
#go build main.go
go build -o ${provider_bin}

mkdir -p $provider_dir
cp $provider_bin $provider_dir/$provider_bin
rm $provider_bin
