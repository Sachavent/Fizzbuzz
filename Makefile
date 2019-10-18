build:
	go build  -o fizzbuzz-api cmd/main.go

run:
	make build
	./fizzbuzz-api

test:
	go test ./...
