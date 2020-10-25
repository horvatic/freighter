# Freighter

![Build](https://github.com/horvatic/freighter/workflows/Build/badge.svg)

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

### make buildTest
Used to build the test service. Compiled result are located at bin/freighterTest

### make clean
Used to clean all build and running artifacts

### make unitTest
Used to run unit test

### make runTestEnv
Used to build and run a test environment where freighter and freighterTest are both Running

### make run
Used to build and run freighter
