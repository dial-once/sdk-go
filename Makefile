
all: build test coverage

build:
	go build -o dialonce

coverage:
	go tool cover -html=coverage.out -o coverage.html

lint:
	golint

test:
	go test -coverprofile=coverage.out
