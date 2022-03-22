TEST?=$$(go list ./...)
PKG_NAME=registry

default: build

build: fmt generate
	go install

test: fmt generate
	go test $(TESTARGS) -timeout=30s $(TEST)

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s ./$(PKG_NAME)

# Currently required by tf-deploy compile


lint:
	@echo "==> Checking source code against linters..."
	@golangci-lint run ./$(PKG_NAME)

tools:
	@echo "==> installing required tooling..."
	go install github.com/client9/misspell/cmd/misspell@v0.3.4
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.0


generate:
	go generate  ./...


.PHONY: build test  fmt fmtcheck lint tools  generate
