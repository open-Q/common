all: test lint

test:
	go test -p 1 -coverpkg=./... ./...

lint:
	golangci-lint cache clean
	golangci-lint run --config .golangci.yml --timeout=5m

todo:
	grep -rn --exclude-dir=.git --exclude=Makefile --exclude-dir=.idea --exclude=TODO.md '// TODO' . > 'TODO.md'
