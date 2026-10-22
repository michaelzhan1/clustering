.PHONY: build
build:
	go build -o bin/main cmd/main/main.go

test:
	go test ./... -v

lint:
	go tool golangci-lint run ./...

clean:
	rm -rf bin
