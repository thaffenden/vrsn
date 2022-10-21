BINARY_NAME=vrsn
DIR ?= ./...
VERSION ?= $(shell head -n 1 VERSION)

define circleci-docker
	docker run --rm -v ${PWD}/.circleci:/repo circleci/circleci-cli:alpine 
endef

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X github.com/thaffenden/vrsn/cmd.Version=${VERSION}" -o ${BINARY_NAME}

.PHONY:
build-image:
	@docker build --tag ${BINARY_NAME}:local .

.PHONY: fmt
fmt:
	@go fmt ${DIR}

.PHONY: ghrc-login
ghcr-login:
	@op run --env-file='./.env' -- docker login ghcr.io -u "${GITHUB_USER}" -p "${GITHUB_TOKEN}"

.PHONY: install
install: build
	@sudo cp ./${BINARY_NAME} /usr/bin/${BINARY_NAME}

.PHONY: lint
lint:
	@golangci-lint run -v ${DIR}

.PHONY: push-tag
push-tag:
	@git tag -a ${VERSION}
	@git push origin ${VERSION}

.PHONY: release
release: ghrc-login push-tag
	@op run --env-file='./.env' -- goreleaser release --rm-dist

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover

.PHONY: validate-ci
validate-ci:
	@$(circleci-docker) config validate /repo/config.yml

.PHONY: validate-orb
validate-orb:
	@$(circleci-docker) orb validate /repo/orb.yml
