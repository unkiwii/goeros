SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=eros

VERSION=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date +%Y-%m-%d\ %H:%M)

LDFLAGS=-ldflags "-X github.com/unkiwii/goeros/info.VERSION '${VERSION}' -X github.com/unkiwii/goeros/info.DATE '${BUILD_TIME}'"

.DEFAULT_GOAL: ${BINARY}

$(BINARY): $(SOURCES)
	@echo Building...
	@go build ${LDFLAGS} -o ${BINARY} eros.go

.PHONY: clean
clean:
	@echo Cleaning...
	@if [ -f ${BINARY} ]; then rm ${BINARY}; fi
