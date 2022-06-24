VERSION := v0.0.1

.PHONY: default
default: build

build:
	go build -o employee_management
	chmod +x employee_management

test:
	go test -v ./models