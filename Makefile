test:
	go test -count=1 -v -race -cover ./...
.PHONY: test

audit:
	go list -json -m all | nancy sleuth
.PHONY: audit

build:
	go build ./...
.PHONY: build

lint:
	exit
.PHONY: lint
