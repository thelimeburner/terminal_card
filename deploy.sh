#!/bin/bash
set -e
V=$1
REPO="thelimeburner"
PRE=veea
IMAGE_NAME=terminal_card

IMAGE=${REPO}/${PRE}/${IMAGE_NAME}:${V}


gobuilder () {
	if [ -z "$1" -o -z "$2" ]
	then
		echo "Usage: $0 5|6  PACKAGENAME"
		return
	fi
	if [ $1 -eq 5 ]
	then
		echo "Using arm"
		ARCH=arm
	elif [ $1 -eq 6 ]
	then
		echo "Using arm64"
		ARCH=arm64
	else
		echo "Wrong Architecture!"
		return
	fi
	NAME=$2
	OUTFILE=$(basename $NAME)
	env GOOS=linux GOARCH=${ARCH} go build -o ${OUTFILE}-${ARCH} ${NAME}
}

echo "######## BUILDING BINARY ########"
#gobuilder 6 github.com/Max2Inc/terminal-card
env GOOS=linux GOARCH=arm64 go build -o terminal-card-arm64 

echo "######## BUILDING CONTAINER ########"
docker build -t ${IMAGE} -f Dockerfile .


echo "######## CLEANING UP ########"

rm -rf terminal-card-arm64

echo "######## Pushing Container ########"
docker push ${IMAGE}
