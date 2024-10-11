.PHONY=build

build-producer:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/producer producer/main.go

run-producer: build-producer
	@./bin/producer

build-consumer:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/consumer consumer/main.go

run-consumer: build-consumer
	@./bin/consumer

coverage:
	@go test -v -cover ./...

test:
	@go test -v ./...

