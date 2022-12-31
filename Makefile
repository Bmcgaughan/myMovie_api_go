build:
	GOBIN=${PWD}/functions go install ./...

phony: build