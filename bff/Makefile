gen_proto:
	mkdir -p lib
	protoc -I=./../proto --go_out=lib --go-grpc_out=lib ./../proto/*.proto
	go mod tidy

gen_grapql:
	go generate ./...

fmt:
	gofmt -s -l -w .