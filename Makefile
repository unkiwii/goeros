SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=eros

VERSION=$(shell git rev-parse --short HEAD)
YEAR_COMPILED=$(shell date +%Y)
MONTH_COMPILED=$(shell date +%m)
DAY_COMPILED=$(shell date +%d)

LDFLAGS=-ldflags "\
-X github.com/unkiwii/goeros/info.Version=${VERSION}\
-X github.com/unkiwii/goeros/info.YearCompiled=${YEAR_COMPILED}\
-X github.com/unkiwii/goeros/info.MonthCompiled=${MONTH_COMPILED}\
-X github.com/unkiwii/goeros/info.DayCompiled=${DAY_COMPILED}\
"

.DEFAULT_GOAL: ${BINARY}

$(BINARY): $(SOURCES)
	@echo Building...
	@go build ${LDFLAGS} -o ${BINARY} eros.go

.PHONY: clean
clean:
	@echo Cleaning...
	@if [ -f ${BINARY} ]; then rm ${BINARY}; fi

.PHONY: run
run: $(BINARY)
	@echo Running...
	@./$(BINARY)
