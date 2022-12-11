.DEFAULT_GOAL := build

build:
	./scripts/build_all.sh

test:
	go test ./...

hooks:
	cd .git/hooks; ln -s ../../githooks/* .

.PHONY: build test hooks
