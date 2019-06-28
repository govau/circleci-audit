#!/bin/bash

set -euo pipefail

docker build --tag circleci-audit .

docker run -it --rm \
  -e CIRCLE_TOKEN="${CIRCLE_TOKEN}" \
  -t circleci-audit

