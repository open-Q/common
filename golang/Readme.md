# [open-Q] Golang

## Golang shared files and libraries

### Run test
```bash
make test
```
or
```bash
go test -p 1 -coverpkg=./... ./...
```

### Run linter
```bash
make lint
```
or
```bash
golangci-lint run --config .golangci.yml --timeout=5m
```