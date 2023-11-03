gen:
	mkdir -p lib
	rm -r lib/*
	protoc -I=./../proto --go_out=lib --go-grpc_out=lib ./../proto/*.proto

fmt:
	gofmt -s -l -w .