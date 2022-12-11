.DEFAULT_GOAL := build

build:
	./scripts/build_all.sh

docker:
	./scripts/docker_build_all.sh

test:
	go test ./...

hooks:
	cd .git/hooks; ln -s ../../githooks/* .

.PHONY: build docker test hooks
