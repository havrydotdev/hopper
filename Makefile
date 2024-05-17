HOPPER_BINARY=bin/hopper

hot:
	air

run:
	go run cmd/main.go

build:
	CGO_ENABLED=0 go build -v -o ${HOPPER_BINARY} cmd/main.go

test:
	go test -v ./internal