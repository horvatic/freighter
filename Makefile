.PHONY: format
format:
	go fmt ./...

.PHONY: mod
mod:
	go mod tidy

.PHONY: build
build:  clean mod
	go build -o bin/freighter cmd/freighter/main.go

.PHONY: buildTest
buildTest:  clean mod
		go build -o bin/freighterTest test/service/cmd/freighter/main.go

.PHONY: clean
clean:
	rm -rf nohup.out
	go clean
	rm -rf bin

.PHONY: unitTest
unitTest:    clean mod
	go clean -testcache
	go test ./...

.PHONY: runTestEnv
runTestEnv: clean mod build buildTest
	nohup bin/freighter & >/dev/null &
	nohup bin/freighterTest & >/dev/null &
	@echo
	@echo
	@echo "****Process Running****"
	@ps -a | grep "freighter"

.PHONY: run
run:    clean mod build
	nohup bin/freighter & >/dev/null &
	@echo
	@echo
	@echo "****Process Running****"
	@ps -a | grep "freighter"
