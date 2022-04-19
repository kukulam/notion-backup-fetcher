ALL_SRC?=$$(find ./ -name '*.go')

.default: build

build:
	go build ./...
.PHONY: build

check: imports fmt lint vet
.PHONY: check

test:
	@go test -count=1 -v -coverprofile=coverage.txt ./...
.PHONY: test

cover: test
	go tool cover -html=coverage.txt
.PHONY: cover

install-lint-deps:
	@curl -sSfL -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.44.0
.PHONY: install-lint-deps

lint:
	@golangci-lint run ./...
.PHONY: lint

vet:
	@go vet ./...; if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs."; \
		exit 1; \
	fi
.PHONY: vet

fmt:
	@go fmt ./...
.PHONY: gmt

install-goimports:
	@go install golang.org/x/tools/cmd/goimports
.PHONY: install-goimports

imports:
	@goimports -w $(ALL_SRC)
.PHONY: imports