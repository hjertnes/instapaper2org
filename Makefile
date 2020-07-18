run:
	go run main.go
clean:
	rm instapaper2org
build:
	go build
test:
	go test ./...
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
lint:
	golangci-lint run --fix
