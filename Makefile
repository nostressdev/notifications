PROTO := $(shell find api -name "*.proto")

all: build

build:
	go build -o bin/service cmd/main.go

docker.build: privatedeps
	docker build . -t notifications:latest

.PHONY: clean privatedeps proto

privatedeps: proto
	mkdir -p .libs/proto

proto: $(PROTO)
	protoc -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate -I . $(PROTO) --go_out=plugins=grpc:. --validate_out="lang=go:."

clean:
	rm -rf .libs