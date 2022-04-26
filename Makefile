
.PHONY: build clean test help default



BIN_NAME=Skyplatform

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')


default: test

help:
	@echo 'Management commands for Skyplatform:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X Skyplatform/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X Skyplatform/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

build-race:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -race -ldflags "-X Skyplatform/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X Skyplatform/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}-race


build-profile-mem:
	go build -ldflags "-X Skyplatform/version.Profile=on -X Skyplatform/version.ProfileType=MEM" -o bin/${BIN_NAME}-prof-mem

build-profile-cpu:
	go build -ldflags "-X Skyplatform/version.Profile=on -X Skyplatform/version.ProfileType=CPU" -o bin/${BIN_NAME}-prof-cpu

get-deps:
	dep ensure

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

