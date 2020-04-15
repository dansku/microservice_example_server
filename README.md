
## Microservice example project Server

[![Lint Status](https://github.com/dansku/microservice_example_server/workflows/golangci-lint/badge.svg)](https://github.com/dansku/microservice_example_server/actions) [![Test Status](https://github.com/dansku/microservice_example_server/workflows/code-test/badge.svg)](https://github.com/dansku/microservice_example_server/actions)

This is a sample project to build a Golang server that communicates via gRPC and takes an array of numbers and a math operation and returns the result.
Operations possible are `sum`, `subtract`, `multiply`, and `divide`.

### Running this from a docker container
* Build: `docker build . -t srv`
* Run: `docker run -d -p 5300:5300 srv`

### Notes
* Generate gRPC protobuf file: `protoc -I. --go_out=plugins=grpc:. calculate.proto`