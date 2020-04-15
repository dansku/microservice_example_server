
## Microservice example project Server

[![Lint Status](https://github.com/dansku/microservice_example_server/workflows/golangci-lint/badge.svg)](https://github.com/dansku/microservice_example_server/actions) [![Test Status](https://github.com/dansku/microservice_example_server/workflows/code-test/badge.svg)](https://github.com/dansku/microservice_example_server/actions)


### Build docker container
* From the `srv` folder, build container with `docker build . -t srv`
* Run container `docker run -d -p 5300:5300 srv`
* 

### Notes
* Generate protobuf file: `protoc -I. --go_out=plugins=grpc:. calculate.proto`

# CI/CD Pipelines
The CI/CD pipelines are built using Github Actions Workflow. This is just a sample project, with a lot more to add in the future.