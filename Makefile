PROTO := $(shell find api -name "*.proto")
GO := $(shell find . -name "*.go")

proto: $(PROTO)
	protoc -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate -I . $(PROTO) --go_out=plugins=grpc:. --validate_out="lang=go:."

.PHONY: clean

run-notifications: bin/notifications
	./bin/notifications

bin/notifications: proto $(GO)
	 go build -o bin/notifications cmd/main.go

docker.build:
	docker build -f docker/notifications/Dockerfile . -t notifications:latest
	docker-compose up

docker.push:
	docker build -f docker/notifications/Dockerfile . -t notifications:latest
	docker tag notifications:latest cr.yandex/crpbivccj2pdgffg9qug/syntok-notifications:latest
	docker push cr.yandex/crpbivccj2pdgffg9qug/syntok-notifications:latest
