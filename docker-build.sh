#!/bin/bash

set -ex

PARENT_DIR=$(basename "${PWD%/*}")
CURRENT_DIR="${PWD##*/}"
IMAGE_NAME="$PARENT_DIR/$CURRENT_DIR"
TAG="${1}"

docker build -t ${IMAGE_NAME}:${TAG} .
docker tag ${IMAGE_NAME}:${TAG} ${IMAGE_NAME}:latest
docker push ${IMAGE_NAME}
docker tag ${IMAGE_NAME}:latest ${IMAGE_NAME}:${TAG}
docker push ${IMAGE_NAME}
