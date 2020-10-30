all: test lint

dep:
	go get -u github.com/micro/protoc-gen-micro/v2

test:
	go test -p 1 -coverpkg=./log/...,./service/...,./storage/... ./log/... ./service/... ./storage/...

lint:
	golangci-lint cache clean
	golangci-lint run --config .golangci.yml --timeout=5m

todo:
	grep -rn --exclude-dir=.git --exclude=Makefile --exclude-dir=.idea --exclude=TODO.md '// TODO' . > 'TODO.md'

gen-protobuf:
	mkdir -p ./proto/user
	rm -rf ./proto/user/*
	protoc -I ${GOPATH}/src/github.com/open-Q/common/proto/ --micro_out=./proto/user --go_out=plugins=micro:./proto/user ${GOPATH}/src/github.com/open-Q/common/proto/user.proto

start_images:
	docker run --rm -d -p 27017:27017 --name mongodb mongo:4.2 --replSet rs0
	chmod +x ./scripts/checkMongo.sh
	./scripts/checkMongo.sh
	docker exec -it mongodb mongo --eval 'rs.initiate({"_id":"rs0","members":[{"_id":0,"host":"localhost:27017"}]})'

stop_images:
	echo "Stopping docker"
	docker rm -f mongodb || true
