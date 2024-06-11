all: clean build test

clean:
	rm -f bin/api-server

build:
	go build -o bin/api-server cmd/api/server.go

mocks:
	go generate ./...

test:
	go test ./...

envbuild:
	cd dev && docker compose build

envup:
	cd dev && docker compose up -d

envdown:
	cd dev && docker compose down

genql:
	go run github.com/99designs/gqlgen generate