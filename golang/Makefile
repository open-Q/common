all: test lint

test:
	go test -p 1 -coverpkg=./... ./...

lint:
	golangci-lint cache clean
	golangci-lint run --config .golangci.yml --timeout=5m

todo:
	grep -rn --exclude-dir=.git --exclude=Makefile --exclude-dir=.idea --exclude=TODO.md '// TODO' . > 'TODO.md'

gen-protobuf:
	protoc -I ${GOPATH}/src/github.com/open-Q/common/proto/ --micro_out=./proto/user --go_out=plugins=micro:./proto/user ${GOPATH}/src/github.com/open-Q/common/proto/user.proto
