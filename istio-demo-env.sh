#!/usr/bin/env sh

unset LOCAL_ISTIO_IMGS DOCKER_SOCKET_MOUNT BUILD_ZTUNNEL BUILD_ZTUNNEL_REPO ISTIO_ZTUNNEL_BASE_URL ISTIO_ENVOY_BASE_URL

ARCH=arm64
export TAG="ben-local-$ARCH-$BUILDSHA-debug"
export HUB=gcr.io/solo-oss/bleggett
export LOCALHUB=localhost:5000
export DOCKER_ARCHITECTURES="linux/$ARCH"
export TARGET_ARCH=$ARCH
export LOCAL_ISTIO_IMGS="pilot,ztunnel,istioctl,install-cni,proxyv2,app,app_sidecar_ubuntu_jammy,ext-authz,app"
export INTEGRATION_TEST_FLAGS='--istio.test.nocleanup  --istio.test.ambient --log_output_level=all:debug  --istio.test.kube.helm.values="global.logging.level=all:debug" -v'
