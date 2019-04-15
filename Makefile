GO_BIN ?= go
CURL_BIN ?= curl
SHELL_BIN ?= sh

update:
	$(GO_BIN) get -u
	$(GO_BIN) mod tidy

linter-install: check-gopath
	cd ~
	$(CURL_BIN) -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | $(SHELL_BIN) -s -- -b ${GOPATH}/bin v1.16.0
	$(GO_BIN) get -u github.com/Quasilyte/go-consistent
	$(GO_BIN) get -u github.com/mgechev/revive

test:
	$(GO_BIN) test -failfast ./...

lint:
	golangci-lint run
	go-consistent -pedantic -v ./...
	revive -config revive.toml ./...

check-gopath:
ifndef GOPATH
	$(error GOPATH is undefined)
endif
