.PHONY: run build test clean

BINARY_NAME=ToDo

# "make run" will make the program run directly:
run:
	go run cmd/todo/main.go

# "make build" will compile the code into a binary file named "temp_conv"
build:
	go build ${BINARY_NAME} main.go

# "make test" will run all your tests with verbal output
test:
	go test -v ./...

# "make clean" will remove the binary file
clean:
	go clean
	rm -f ${BINARY_NAME}
