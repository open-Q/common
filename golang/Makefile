all: test lint

dep:
	go get -u github.com/micro/protoc-gen-micro/v2

test:
	go test -p 1 -coverpkg=./...,./log/... ./... ./log/...

lint:
	golangci-lint cache clean
	golangci-lint run --config .golangci.yml --timeout=5m

todo:
	grep -rn --exclude-dir=.git --exclude=Makefile --exclude-dir=.idea --exclude=TODO.md '// TODO' . > 'TODO.md'

gen-protobuf:
	mkdir -p ./proto/user
	rm -rf ./proto/user/*
	protoc -I ${GOPATH}/src/github.com/open-Q/common/proto/ --micro_out=./proto/user --go_out=plugins=micro:./proto/user ${GOPATH}/src/github.com/open-Q/common/proto/user.proto
