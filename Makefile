ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BUILD_DIR := ${PWD}/bin

.PHONY: lint
lint: #### Lint code.
	@go version
	cd ${ROOT_DIR} && \
		golangci-lint --build-tags testing run --max-same-issues 0 --max-issues-per-linter 0

PHONY: test
test:
	go test -race -tags testing -mod=vendor ${ROOT_DIR}/...
