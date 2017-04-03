help:
	@echo "Please use \`make <ROOT>' where <ROOT> is one of"
	@echo "  update-dependencies  to update glide.lock (refs to dependencies)"
	@echo "  dependencies         to install the dependencies"
	@echo "  likelo	              to build the main binary for current platform"
	@echo "  test                 to run unittests"
	@echo "  docker               to build the docker image"
	@echo "  clean                to remove generated files"

clean:
	rm -rf likelo

update-dependencies:
	glide --version 2> /dev/null || curl https://glide.sh/get | sh
	glide up

dependencies:
	glide --version 2> /dev/null || curl https://glide.sh/get | sh
	glide install

likelo: *.go */*/*.go
	$(GO_VARS) $(GO) build -o="likelo" -ldflags="$(LD_FLAGS)" $(ROOT)/cmd/likelo

test: *.go */*.go */*/*.go
	$(GO_VARS) $(GO) test -v $(shell glide novendor) && echo -e "\nTesting is passed."

docker: likelo Dockerfile
	docker build -t $(DOCKER_IMAGE):$(VERSION) .
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest

test-docker:
	docker run $(DOCKER_IMAGE):$(VERSION) version

push:
	docker push $(DOCKER_IMAGE):$(VERSION)
	docker push $(DOCKER_IMAGE):latest


## Project Vars ##########################################################
ROOT := github.com/arastu/likelo
DOCKER_IMAGE := arastu/likelo
.PHONY: help clean update-dependencies dependencies test docker push test-docker

## Commons Vars ##########################################################
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S), Linux)
	OS ?= linux
endif
ifeq ($(UNAME_S), Darwin)
	OS ?= osx
endif
ARCH := $(shell uname -m)
ifeq ($(ARCH), unknown)
	ARCH := x86_64
endif
ifeq ($(ARCH), i386)
	ARCH = x86_32
endif

GO_VARS =
GO ?= CGO_ENABLED=1 go
GIT ?= git
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).Commit=$(COMMIT) -X $(ROOT).BuildTime=$(BUILD_TIME)