#!/usr/bin/env bash

# shellcheck disable=SC1090
source "$(dirname "$0")/../vendor/knative.dev/hack/e2e-tests.sh"
source "$(dirname "$0")/e2e-common.sh"

set -Eeuox pipefail

export TEST_IMAGE_TEMPLATE="${EVENTING_TEST_IMAGE_TEMPLATE}"

env

failed=0

(( !failed )) && install_serverless || failed=1

(( failed )) && exit 1

success

