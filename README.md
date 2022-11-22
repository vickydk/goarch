# goarch
Golang Clean Architetecture

### Update Config

1. go to folder `resources` and open file `config.json`
2. Update the value based on your environment

### Running HTTP

1. Clone this repo inside your $GOPATH/src/github.com/vickydk
2. Update config.json
3. Install all dependency `go mod tidy`
4. Try run Application Restful API with `make run` and access `http://localhost:8811`

### Running gRPC

1. Clone this repo inside your $GOPATH/src/github.com/vickydk
2. Update config.json
3. Install all dependency `go mod tidy`
4. Try run Application Restful API with `make run-grpc` and access `localhost:8822`

### Running gRPC & HTTP

1. Clone this repo inside your $GOPATH/src/github.com/vickydk
2. Update config.json
3. Install all dependency `go mod tidy`
4. Try run Application Restful API with `make run-app` 
5. Access `localhost:8822` for GRPC
6. Access `localhost:8811` for HTTP

### Check Http Up

Call this url: `http://localhost:8811/`, if you see response `OK` so the application already run successfully


### Proto File

#### Prerequisite :

1. Install `protoc` version 3.17.3

    - the binary can be found [here](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3
    - download it and put the binary in a folder registered on your `$PATH` or `/usr/local/bin`
    - run `protoc --version` should return protoc `libprotoc 3.17.3`

2. Install dependency library
    - go get -u google.golang.org/grpc
    - go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    - go install github.com/golang/protobuf/protoc-gen-go@latest // if cannot install line 17
    - go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    - go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
   
#### Compile Proto
1. go to folder `pkg/shared/grpc`
2. run this command to compile the proto file `./protoc.sh -p {protofile_name_without.dot.proto`, for example `./protoc.sh -p user`

## License

gorsk is licensed under the MIT license. Check the [LICENSE](LICENSE) file for details.

## Author

[VickyDk](https://github.com/vickydk)