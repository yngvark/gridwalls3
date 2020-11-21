SHELL         = bash
IMAGE = ghcr.io/yngvark/zombie

GO := $(shell command -v go 2> /dev/null)
ifndef GO
$(error go is required, please install)
endif

GOPATH			:= $(shell go env GOPATH)
GOBIN			?= $(GOPATH)/bin
GOFUMPT			:= $(GOBIN)/gofumpt

PKGS  = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
FILES = $(shell find . -name '.?*' -prune -o -name vendor -prune -o -name '*.go' -print)

gofumpt:
	$(GO) get -u mvdan.cc/gofumpt

fmt: gofumpt
	$(GO) fmt $(PKGS)
	$(GOFUMPT) -s -w $(FILES)

build-docker:
	docker build . -t $(IMAGE)

run-docker: build-docker
	docker run --name zombie --rm -p 8080:8080 $(IMAGE)

push:
	docker push $(IMAGE)

