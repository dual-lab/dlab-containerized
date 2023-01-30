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

appName=admit-tag-image
version=$(cat pkg/Makefile | grep "__version__ :=" | cut -d "=" -f 2 | xargs)
echo "Building architecture ${arch}"
# First stage build compile go code
buildcontainer=$(buildah from golang:1.18)
buildcontainerMnt=$(buildah mount ${buildcontainer})

trap "buildah unmount $buildcontainer" ERR
trap "buildah rm $buildcontainer" ERR

buildah run ${buildcontainer} -- mkdir work_dir
buildah config --workingdir work_dir ${buildcontainer}
buildah run ${buildcontainer} -- mkdir scripts
buildah run ${buildcontainer} -- mkdir pkg
cp -r scripts/* ${buildcontainerMnt}/work_dir/scripts/
cp -r pkg/* ${buildcontainerMnt}/work_dir/pkg/
cp go.mod ${buildcontainerMnt}/work_dir/
cp go.sum ${buildcontainerMnt}/work_dir/
cp KBUILD ${buildcontainerMnt}/work_dir/
cp Makefile ${buildcontainerMnt}/work_dir/
cp main.go ${buildcontainerMnt}/work_dir/

buildah run ${buildcontainer} -- go mod download
buildah run --env CGO_ENABLED=0 --env GOOS=linux --env GOARCH=${arch} ${buildcontainer} -- make
# Second stage copy executable on final image
appContainer=$(buildah from gcr.io/distroless/base-debian11)
appContainerMnt=$(buildah mount ${appContainer})

trap "buildah unmount $appContainer" ERR
trap "buildah rm $appContainer" ERR

cp ${buildcontainerMnt}/work_dir/build/${appName} ${appContainerMnt}/webhook

buildah config --port 443 $appContainer
buildah config --entrypoint '["./webhook"]' $appContainer
#buildah config --cmd '["-p 443", "--certFile admission_crt.pem", "--certKey admission_key.pem"]' $appContainer
buildah config --created-by "dmike16"  $appContainer
buildah config --author "dmike16 at dual-lab.yandex.com" $appContainer
buildah config --label name="dlabc/${appName}:${version}" $appContainer

buildah unmount ${appContainer}
buildah unmount ${buildcontainer}

buildah commit $appContainer dlabc/${appName}:${arch}-${version}

buildah rm ${appContainer}
buildah rm ${buildcontainer}