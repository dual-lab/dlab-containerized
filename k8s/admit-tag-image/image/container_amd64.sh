#!/bin/bash

# Buildah script used to buildmthe image that will run the
# custom admit controller webhook.
# -------
#
# The script is a multistage build
#		1. compile go code using golang:1.18 image
#		2. copy the executable on a distroless image gcr.io/distroless/base-debian11
#

set -Eeuo pipefail
export arch=amd64

buildah unshare ./image/container.sh