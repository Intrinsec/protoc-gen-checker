empty :=
space := $(empty) $(empty)
NAME := checker
PACKAGE := github.com/intrinsec/protoc-gen-$(NAME)
PROTOC_GEN_VALIDATE := 1.3.3

# Pinned tool versions (Go itself is pinned by the `go` directive in go.mod).
GOLANGCI_LINT_VERSION := v2.15.0
GOVULNCHECK_VERSION := v1.9.0
GORELEASER_VERSION := v2.18.3

# protoc-gen-go parameters for properly generating the import path for PGV
GO_IMPORT_SPACES := M$(NAME)/$(NAME).proto=${PACKAGE}/$(NAME),\
	Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
	Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
	Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct,\
	Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
	Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor
GO_IMPORT:=$(subst $(space),,$(GO_IMPORT_SPACES))

.PHONY: build
build: bin/protoc-gen-$(NAME)

.PHONY: install
install: $(NAME)/$(NAME).pb.go
	@go install -v .

$(NAME)/$(NAME).pb.go: bin/protoc-gen-go bin/protoc-gen-validate $(NAME)/$(NAME).proto
	@cd $(NAME) && protoc -I . \
		--plugin=protoc-gen-go=$(shell pwd)/bin/protoc-gen-go \
		--go_opt=paths=source_relative \
		--go_out="${GO_IMPORT}:." $(NAME).proto

bin/protoc-gen-go:
	@GOBIN=$(shell pwd)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go

bin/protoc-gen-validate:
	@GOBIN=$(shell pwd)/bin go install github.com/envoyproxy/protoc-gen-validate@v$(PROTOC_GEN_VALIDATE)


bin/protoc-gen-$(NAME): $(NAME)/$(NAME).pb.go $(wildcard *.go) proto/validate
	@GOBIN=$(shell pwd)/bin go install .

.PHONY: proto/validate
proto/validate:
	@rm -rf $(shell pwd)/proto/*
	@mkdir -p $(shell pwd)/proto/validate/
	@cd $(shell pwd)/proto/validate/ \
		&& curl -s -L https://github.com/envoyproxy/protoc-gen-validate/archive/v$(PROTOC_GEN_VALIDATE).tar.gz | tar xzf - protoc-gen-validate-$(PROTOC_GEN_VALIDATE) \
		&& mv protoc-gen-validate-$(PROTOC_GEN_VALIDATE)/validate/*.proto . \
		&& rm -r protoc-gen-validate-$(PROTOC_GEN_VALIDATE)

.PHONY: test
test: build
	@echo "/////////////////////////////////////////////////////////////////////////////////////////////"
	@echo "This test is supposed to FAIL. It illustrate the various good/wrong ways of using this plugin."
	@echo "/////////////////////////////////////////////////////////////////////////////////////////////"
	@protoc -I . \
		-I $(shell pwd)/proto/validate/ \
		--plugin=protoc-gen-go=$(shell pwd)/bin/protoc-gen-go \
		--go_out="tests/" tests/*.proto
	@protoc -I . \
		-I $(shell pwd)/proto/validate/ \
		--plugin=protoc-gen-$(NAME)=$(shell pwd)/bin/protoc-gen-$(NAME) \
		--checker_opt=strict \
		--$(NAME)_out=tests tests/*.proto


.PHONY: tools
tools:
	@GOBIN=$(shell pwd)/bin go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	@GOBIN=$(shell pwd)/bin go install golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION)

.PHONY: tools-release
tools-release:
	@GOBIN=$(shell pwd)/bin go install github.com/goreleaser/goreleaser/v2@$(GORELEASER_VERSION)


.PHONY: lint
lint:
	@golangci-lint run .


.PHONY: vuln
vuln:
	@govulncheck ./...


.PHONY: unit
unit:
	@go test ./... -count=1 -race


.PHONY: release-snapshot
release-snapshot:
	@goreleaser release --snapshot --clean

.PHONY: release-check
release-check:
	@goreleaser check


.PHONY: clean
clean:
	@rm -fv tests/*.pb.go


.PHONY: distclean
distclean: clean
	@rm -frv bin/protoc-gen-go bin/protoc-gen-validate bin/protoc-gen-$(NAME) bin/golangci-lint bin/govulncheck bin/goreleaser $(NAME)/$(NAME).pb.go proto/* dist/
