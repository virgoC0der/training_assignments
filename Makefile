VERSION := v0.0.1
TEST_TARGET := ./models
BINARY_NAME := employee_management

.PHONY: default
default: build

build:
	go build -o $(BINARY_NAME)
	chmod +x $(BINARY_NAME)

test:
	go test -v $(TEST_TARGET)

convey:
	goconvey