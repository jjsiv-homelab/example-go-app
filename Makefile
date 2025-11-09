MODULE = github.com/jjsiv-homelab/example-go-app

deps:
	go mod download

clean:
	go fmt ./...
	go vet ./...
	go mod tidy

VERSION ?= $(shell git describe --tags --always)
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)
build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -trimpath -ldflags "-X $(MODULE)/internal/version.commit=$(COMMIT_SHA) -X $(MODULE)/internal/version.version=$(VERSION)" -o ./bin/example-app ./cmd/example-go-app/main.go
