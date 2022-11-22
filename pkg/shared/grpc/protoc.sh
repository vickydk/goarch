#!/bin/bash

POSITIONAL=()
while [[ $# -gt 0 ]]
do
KEY="$1"

case $KEY in
    -p|--path)
    PROTO_PATH="$2"
    shift # past argument
    shift # past value
    ;;
esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

if [[ ! $PROTO_PATH ]]; then
    echo "Please specify proto folder with -p option"
    exit 1
fi

protoc -I proto/ \
--go_out=./pb --go_opt=paths=source_relative \
--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
--doc_out=./doc/$PROTO_PATH --doc_opt=html,$PROTO_PATH.html \
proto/$PROTO_PATH/*.proto
