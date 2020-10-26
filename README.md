# Freighter

![Push](https://github.com/horvatic/freighter/workflows/Push/badge.svg)

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/horvatic/freighter.svg)](https://github.com/horvatic/freighter)

[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

[Storyboard](https://trello.com/b/JUoGW49f/freighter)

## Overview
A lightweight standalone api manager with built in load balancing, and plugin middleware

## Commands
All commands use will be using make.

### make format
Used to run a go fmt

### make mod
Used to mod tidy

### make build
Used to build the main application. Compiled result are located at bin/freighter

### make build-test
Used to build the test service. Compiled result are located at bin/freighterTest

### make clean
Used to clean all build and running artifacts

### make unit-test
Used to run unit test

### make run-test-env
Used to build and run a test environment where freighter and freighterTest are both Running

### make run
Used to build and run freighter

### make docker-build
Used to build a docker image

### make docker-run
Used to run the docker image
