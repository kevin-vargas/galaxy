#!/bin/bash
IMAGE_NAME="docker.fast.ar/galaxy"
TAG="latest"
build(){
    docker build -t ${IMAGE_NAME}:${TAG} .
}

push(){
    docker image push ${IMAGE_NAME}:${TAG}
}

move_directory(){
    local DIRECTORY="deployment"
    cd ${DIRECTORY}
}

deploy(){
    local FILENAME="deployment.yml"
    kubectl apply -f ${FILENAME}
}

{
    build &&
    push &&
    move_directory &&
    deploy
}