.PHONY: build
build:
	go mod tidy
	go build server.go

.PHONY: clean
clean:
	rm server

.PHONY: config
config:
	cd config \
	&& protoc -I . --go_out=. config.proto \
	&& protoc-go-inject-tag -input=config.pb.go