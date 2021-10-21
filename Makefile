empty :=
space := $(empty) $(empty)
NAME := checker
PACKAGE := github.com/intrinsec/protoc-gen-$(NAME)
PROTOC_GEN_VALIDATE := 0.6.1

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
	@GOBIN=$(shell pwd)/bin go install github.com/envoyproxy/protoc-gen-validate


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


.PHONY: clean
clean:
	@rm -fv tests/*.pb.go


.PHONY: distclean
distclean: clean
	@rm -frv bin/protoc-gen-go bin/protoc-gen-validate bin/protoc-gen-$(NAME) $(NAME)/$(NAME).pb.go proto/*
