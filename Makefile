.DEFAULT_GOAL := all

.PHONY: all
all: test

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean ./...
