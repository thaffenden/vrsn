BINARY_NAME=vrsn
DIR ?= ./...
VERSION ?= $(shell head -n 1 VERSION)

.PHONY: build
build:
	@go build -ldflags "-X github.com/thaffenden/vrsn/cmd.Version=${VERSION}" -o ${BINARY_NAME}

.PHONY: fmt
fmt:
	@go fmt ${DIR}

.PHONY: lint
lint:
	@golangci-lint run -v ${DIR}

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover
