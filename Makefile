ORG     := $(shell basename $(realpath ..))
NAME    := $(shell basename $(PWD))

build:
	go build .
.PHONY: build

check:
	go vet $(shell go list ./... | grep -v /vendor/)
.PHONY: check

test:
	go test -v $(shell go list ./... | grep -v /vendor/) -cover -race -p=1
.PHONY: test

tools:
	go get -u github.com/roboll/ghr github.com/mitchellh/gox
.PHONY: tools

cross:
	@mkdir -p dist
	gox -os '!freebsd' -arch '!arm' -output "dist/${NAME}_{{.OS}}_{{.Arch}}"
.PHONY: cross

release: cross
	@ghr -b ${BODY} -t ${GITHUB_TOKEN} -u ${ORG} ${TAG} dist
.PHONY: release

TAG  = $(shell git describe --tags --abbrev=0 HEAD)
LAST = $(shell git describe --tags --abbrev=0 HEAD^)
BODY = "`git log ${LAST}..HEAD --oneline --decorate` `printf '\n\#\#\# [Build Info](${BUILD_URL})'`"
