DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

all: deps
	go-bindata -ignore=.gitignore assets/...
	go install ./...

test: test-deps
	go list ./... | xargs -n1 go test

release: release-deps
	gox ./...

deps:
	go get github.com/websysd/websysd

test-deps:
	go get github.com/stretchr/testify

release-deps:
	go get github.com/mitchellh/gox

rocker: rocker-deps
	rocker build

rocker-deps:
	go get github.com/grammarly/rocker

dockerhub: rocker
	docker push iankent/websysd

.PNONY: all test release deps test-deps release-deps rocker rocker-deps dockerhub
