#!/usr/bin/env sh

unset LOCAL_ISTIO_IMGS LOCAL_BOOKINFO_IMAGES DOCKER_SOCKET_MOUNT BUILD_ZTUNNEL BUILD_ZTUNNEL_REPO ISTIO_ZTUNNEL_BASE_URL ISTIO_ENVOY_BASE_URL

ARCH=arm64
export LOCAL_ISTIO_IMGS="pilot,ztunnel,istioctl,install-cni,proxyv2,app,app_sidecar_ubuntu_jammy"
export LOCAL_BOOKINFO_IMGS="examples-bookinfo-details-v1,examples-bookinfo-productpage-v1"
export EXTERNAL_BOOKINFO_IMGS="examples-bookinfo-ratings-v1,examples-bookinfo-reviews-v1,examples-bookinfo-reviews-v2,examples-bookinfo-reviews-v3"
export TAG="ben-local-arm64-3a2a8e09ce-debug"
export HUB=gcr.io/solo-oss/bleggett
export LOCALHUB=localhost:5000
export EXTERNAL_BOOKINFO_TAG="1.18.0"
export EXTERNAL_BOOKINFO_HUB="docker.io/istio"
export INTERNAL_BOOKINFO_TAG="ben-local-arm64-5c13a38d88-debug"
export INTERNAL_BOOKINFO_HUB="$HUB"
export DOCKER_ARCHITECTURES="linux/$ARCH"
export TARGET_ARCH=$ARCH
export INTEGRATION_TEST_FLAGS='--istio.test.nocleanup  --istio.test.ambient --log_output_level=all:debug  --istio.test.kube.helm.values="global.logging.level=all:debug" -v'

copy_images_local() {
  for i in ${LOCAL_ISTIO_IMGS//,/ }; do
    INT_COPYREMOTE="${HUB}/${i}:${TAG}"
    INT_COPYLOCAL="localhost:5000/${i}:${TAG}"
    docker pull "${INT_COPYREMOTE}"
    docker tag "${INT_COPYREMOTE}" "${INT_COPYLOCAL}"
    docker push "${INT_COPYLOCAL}"
  done

  for i in ${LOCAL_BOOKINFO_IMGS//,/ }; do
    INT_COPYREMOTE="${INTERNAL_BOOKINFO_HUB}/${i}:${INTERNAL_BOOKINFO_TAG}"
    INT_COPYLOCAL="localhost:5000/${i}:${INTERNAL_BOOKINFO_TAG}"
    docker pull "${INT_COPYREMOTE}"
    docker tag "${INT_COPYREMOTE}" "${INT_COPYLOCAL}"
    docker push "${INT_COPYLOCAL}"
  done

  for i in ${EXTERNAL_BOOKINFO_IMGS//,/ }; do
    INT_COPYREMOTE="${EXTERNAL_BOOKINFO_HUB}/${i}:${EXTERNAL_BOOKINFO_TAG}"
    INT_COPYLOCAL="localhost:5000/${i}:${EXTERNAL_BOOKINFO_TAG}"
    docker pull "${INT_COPYREMOTE}"
    docker tag "${INT_COPYREMOTE}" "${INT_COPYLOCAL}"
    docker push "${INT_COPYLOCAL}"
  done
}
