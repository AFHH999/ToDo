.PHONY: run build test clean check lint
BIN_DIR=bin
BINARY_NAME=${BIN_DIR}/todo

# "make check will test for compromised dependencies"
check:
		govulncheck ./...

# "make lint will perform a static check to find bugs or style issues"
lint:
	golangci-lint run

# "make run" will make the program run directly:
run:
	go run cmd/todo/main.go

# "make build" will compile the code into a binary file named "ToDo"
build:
	mkdir -p ${BIN_DIR}
	go build -o ${BINARY_NAME} cmd/todo/main.go

# "make test" will run all your tests with verbal output
test:
	go test -v ./...

# "make clean" will remove the binary file
clean:
	go clean
	rm -fr ${BIN_DIR}
