#!/usr/bin/env bash

set -xeuo pipefail

repo_root_dir=$(dirname "$(realpath "${BASH_SOURCE[0]}")")/..

tmp_dir=$(mktemp -d)
git clone --branch main https://github.com/openshift-knative/hack "$tmp_dir"

pushd "$tmp_dir"
go build -o "$tmp_dir"/generate cmd/generate/generate.go
popd

"$tmp_dir"/generate \
  --root-dir "${repo_root_dir}" \
  --generators dockerfile \
  --dockerfile-image-builder-fmt "registry.ci.openshift.org/openshift/release:rhel-8-release-golang-%s-openshift-4.16"

set +x
