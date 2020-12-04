VERSION?="0.1.0"

PROJECTNAME=$(shell basename "$(PWD)")

.PHONY: all
all:
	lint
	build
	test

.PHONY: build
build:
	echo "Building $(PROJECTNAME)"
	GOOS=windows GOARCH=386 go build -o bin/main.exe cmd/jogtrot/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-arm64 cmd/jogtrot/main.go

.PHONY: run
run:
	echo "Starting $(PROJECTNAME)"
	go run cmd/jogtrot/main.go

.PHONY: docs
docs:
	cd cmd; cd jogtrot; swag init .../cmd/jogtrot

.PHONY: test
test:
	echo "Testing $(PROJECTNAME)"	
	go test ./cmd/jogtrot/... -v
	# ginkgo ./cmd/jogtrot/

.PHONY: lint
lint:
	golangci-lint run

