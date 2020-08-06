test:
	go test -race ./...
build:
	go build .
run: build
	./go-app