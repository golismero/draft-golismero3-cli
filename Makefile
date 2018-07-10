.PHONY: g3cli
SHELL := /bin/bash
GO_VERSION := 1.10.3
GO_REPO := github.com/golismero/golismero3-cli


define godock
	@docker run --rm -it \
		-v $$(pwd)/.go:/go -v $$(pwd):/go/src/$(GO_REPO) \
		-w /go/src/$(GO_REPO) \
		golang:$(GO_VERSION) $(1) $(2)
endef

define gocmd
	$(call godock, go, $(1))
endef

all: g3cli

acceptance: g3cli
	$(call gocmd,get github.com/DATA-DOG/godog/cmd/godog)
	$(call gocmd,get github.com/DATA-DOG/godog)

	@docker run --rm -it \
		-v $$(pwd)/.go:/go -v $$(pwd):/go/src/$(GO_REPO) \
		-w /go/src/$(GO_REPO)/internal/features \
		golang:$(GO_VERSION) godog -t "~@wip" . 


g3cli: test
	$(call gocmd,install -ldflags "-s \
		-X main.Version=$$(cat VERSION) \
		-X main.BuildTime=$$(TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ') \
		-X main.GitHash=$$(git rev-parse HEAD)" $(GO_REPO)/cmd/g3cli/...)

test: deps
	$(call gocmd,test -race $(GO_REPO)/pkg/...)

deps:
	$(call gocmd,get github.com/spf13/cobra)