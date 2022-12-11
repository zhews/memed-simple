.DEFAULT_GOAL := build

build:
	./scripts/build_all.sh

docker:
	./scripts/docker_build_all.sh

test:
	go test ./...

test-env-up:
	cd test; docker-compose up -d

test-env-down:
	cd test; docker-compose down

hooks:
	cd .git/hooks; ln -s ../../githooks/* .

.PHONY: build docker test test-env-up test-env-down hooks
