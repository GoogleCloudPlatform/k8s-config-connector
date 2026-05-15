.PHONY: all
all: build

.PHONY: build
build: controllerbuilder benchmark

.PHONY: controllerbuilder
controllerbuilder:
	go build -o ${GOPATH}/bin/controllerbuilder main.go

.PHONY: benchmark
benchmark:
	go build -o ${GOPATH}/bin/benchmark cmd/benchmark/type/main.go
