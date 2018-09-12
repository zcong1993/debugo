generate:
	@go generate ./...
.PHONY: generate

install:
	@go get -u -v github.com/golang/dep/cmd/dep
	@dep ensure
.PHONY: install

test:
	@go test ./...
.PHONY: test

test.cov:
	@go test ./... -coverprofile=coverage.txt -covermode=atomic
.PHONY: test.cov
