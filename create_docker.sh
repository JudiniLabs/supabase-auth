#!/bin/bash

version=v2.149.0

docker buildx build --platform linux/arm64,linux/amd64 -t codegpt/gotrue:$version --push .