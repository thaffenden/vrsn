BINARY_NAME=vrsn
DIR ?= ./...
VERSION ?= $(shell head -n 1 VERSION)

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X github.com/thaffenden/vrsn/cmd.Version=${VERSION}" -o ${BINARY_NAME}

.PHONY:
build-image:
	@docker build --tag ${BINARY_NAME}:local .

.PHONY: fmt
fmt:
	@go fmt ${DIR}

.PHONY: lint
lint:
	@golangci-lint run -v ${DIR}

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover
