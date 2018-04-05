PKG = github.com/redradrat/tape/cmd
LDFLAGS += -X "$(PKG).GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(PKG).BuildStamp=$(shell date)"
TAGS = ""
BUILD_FLAGS = "-v"

all: build test

build:
	go install $(BUILD_FLAGS) -ldflags '$(LDFLAGS)' -tags '$(TAGS)' 

test:
	go test -cover -race $(TEST_OPTIONS) -v -covermode=atomic -coverprofile=coverage.out -run -timeout=20s ./...

cover: test
	go tool cover -html=coverage.out

vet:
	go tool vet tape.go
