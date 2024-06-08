all: clean build test

clean:
	rm -f bin/api-server

build:
	go build -o bin/api-server cmd/api/server.go

test:
	go test internal ./...

envup:
	cd dev && docker compose up -d

envdown:
	cd dev && docker compose down
