#!/bin/bash
image_name=grpc-aes
container_name=grpc-aes

mode=$1

if [[ $mode == "build" ]]; then
    docker rmi $(docker images | grep ${image_name} | awk '{print $3}')
    docker build -t ${image_name} .
    docker rmi $(docker images | grep none | awk '{print $3}')
fi

if [[ $mode == "run" ]]; then
    docker run -it --rm \
        -p 3000:3000 \
        --env-file .env \
        --name ${container_name} \
        ${image_name}
fi

if [[ $mode == "proto" ]]; then
    docker run \
        -v $PWD/src:/defs \
        namely/protoc \
        --go_out=./cipher --go_opt=paths=source_relative \
        --go-grpc_out=./cipher --go-grpc_opt=paths=source_relative \
        --proto_path=./proto \
        cipher.proto
fi

if [[ $mode == "proto-python" ]]; then
    python -m grpc_tools.protoc \
        -I./src/proto \
        --python_out=. \
        --grpc_python_out=. \
        --pyi_out=. \
        cipher.proto
fi
