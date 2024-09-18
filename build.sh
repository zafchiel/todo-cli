#!/bin/bash

# Name of your application
APP_NAME="todoc"

# Platforms to build for
PLATFORMS="windows/amd64 windows/386 darwin/amd64 darwin/arm64 linux/amd64 linux/386 linux/arm linux/arm64"

# Build for each platform
for PLATFORM in $PLATFORMS; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    OUTPUT="${APP_NAME}-${GOOS}-${GOARCH}"
    if [ $GOOS = "windows" ]; then
        OUTPUT="${OUTPUT}.exe"
    fi

    echo "Building for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o build/$OUTPUT
    if [ $? -ne 0 ]; then
        echo "An error has occurred! Aborting the script execution..."
        exit 1
    fi
done