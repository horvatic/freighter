.PHONY: format
format:
	go fmt ./...

.PHONY: mod
mod:
	go mod tidy

.PHONY: build
.NOTPARALLEL:
build: clean mod unit-test
	go build -o bin/freighter cmd/freighter/main.go

.PHONY: build-test
build-test: clean mod
		go build -o bin/freighterTest test/service/cmd/freighter/main.go

.PHONY: clean
clean:
	rm -rf nohup.out
	go clean
	rm -rf bin

.PHONY: unit-test
.NOTPARALLEL:
unit-test: clean mod
	go clean -testcache
	go test ./...

.PHONY: run-test-env
.NOTPARALLEL:
run-test-env: clean mod unit-test
	docker-compose up --build -d

.PHONY: run
.NOTPARALLEL:
run: clean mod unit-test
	docker-compose up --build -d web

.PHONY: stop
.NOTPARALLEL:
stop:
	docker-compose down

.PHONY: docker-build
.NOTPARALLEL:
docker-build: clean mod unit-test
	docker build . -t freighter

.PHONY: docker-run
docker-run: docker-build
	docker run -p 8080:8080 -d -name=freighter freighter
