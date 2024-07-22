all: clean build test

clean:
	rm -f build/app/bin/api-server

build:
	go build -o build/app/bin/api-server cmd/api/server.go

mocks:
	go generate ./...

test:
	go test ./...

envbuild:
	cd env && docker compose build

envup:
	cd env && docker compose up -d

envdown:
	cd env && docker compose down

genql:
	go run github.com/99designs/gqlgen generate

.PHONY: all build clean