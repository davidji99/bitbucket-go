GO=$(shell which go)
GOVENDOr=$(shell which govendor)
TARGET=$(GOBIN)/go-bb
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
MAIN=$(shell pwd)/cmd/go-bb.go

install: sync
	@${GO} install -ldflags "-s -w -X main.BuildSHA=`git rev-parse --short HEAD`" ${MAIN}

clean:
	@rm -f ${TARGET}

fmt:
	gofmt -w $(GOFMT_FILES)

test:
	@${GO} test ./...

release: sync
	bin/build-release

sync:
	@${GOVENDOR} sync

.PHONY: install test release sync