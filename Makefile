.PHONY: data
data:
	go run cmd/dataset/main.go

.PHONY: build
build:
	go build -o bin/main cmd/main/main.go

.PHONY: test
test:
	go test ./... -v

.PHONY: lint
lint:
	go tool golangci-lint run ./...

.PHONY: dataclean
dataclean:
	rm -rf data

.PHONY: clean
clean:
	rm -rf bin
