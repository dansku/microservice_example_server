
## Server
### Build docker container
* From the `srv` folder, build container with `docker build . -t srv`
* Run container `docker run -d -p 5300:5300 srv`


## Client

Example:
`go run main.go -host=127.1.0.1 -port=5300 -operation subtract -numbers 10.2,19,13.3`

# CI/CD Pipelines