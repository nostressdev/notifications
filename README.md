# Service template

### First steps

0) Rename module in go.mod
1) Copy private deps to .libs in Makefile
```makefile
privatedeps: ;
	mkdir -p .libs/proto
	cp -r ../../proto/go/* .libs/proto
```
2) Set name of service in docker.build in Makefile
```makefile
docker.build: privatedeps
	docker build . -t servicename:latest
```
3) Set subspace in cmd/main.go
```go
subspace.Sub("grpc_fdb_service")
```