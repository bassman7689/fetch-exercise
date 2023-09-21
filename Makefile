.PHONY: build
build:
	@go build -o api cmd/api/main.go

test:
	@go test ./...

