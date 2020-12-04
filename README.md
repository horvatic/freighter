# Freighter

![Push](https://github.com/horvatic/freighter/workflows/Push/badge.svg)

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/horvatic/freighter.svg)](https://github.com/horvatic/freighter)

[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)



## Overview
A lightweight standalone api manager with built in load balancing, and plug-in middleware

[Storyboard](https://trello.com/b/JUoGW49f/freighter)

### Current Behavior:
- reverse proxy
- service registration
- API key

### Future State
- load balancing
- front end. [Link To Repo](https://github.com/horvatic/freighter-bridge)
- plug-in middleware
- mongo backing

### Dependencies
- docker
- docker-compose
- make
- golang 1.15

### Dev Notes
- Service registration happens with the calling of the endpoint of `{gatewayUri}/config/register`
- `/test/service/cmd/freighter/main.go` is a great example service. It will register with freighter on startup
- freighter must be running before any applications can register
- postman integration test are provided at `/test/integrationsfreighter.postman_collection.json`
- when running postman use the `localtest.postman_environment.json` environment


### Setting up a local test environment

#### Running
make run-test-env

#### Cleanup
make stop

## Commands


#### make format
Used to run a go fmt

#### make mod
Used to mod tidy

#### make build
Used to build the main application. Compiled result are located at bin/freighter

#### make build-test
Used to build the test service. Compiled result are located at bin/freighterTest

#### make clean
Used to clean all build and running artifacts

#### make unit-test
Used to run unit test

#### make run-test-env
Used to build and run a test environment where freighter and freighterTest are both Running

#### make stop
Used to stop and clean up a test environment where freighter and freighterTest are both Running

#### make run
Used to build and run freighter

#### make docker-build
Used to build a docker image

#### make docker-run
Used to run the docker image
