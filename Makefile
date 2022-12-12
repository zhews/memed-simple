.DEFAULT_GOAL := build

build:
	./scripts/build_all.sh

docker:
	./scripts/docker_build_all.sh

test:
	go test ./...

coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic -coverpkg=./... ./...

test-env-up:
	cd test; docker-compose up -d

test-env-down:
	cd test; docker-compose down

hooks:
	cd .git/hooks; ln -s ../../githooks/* .

.PHONY: build docker test coverage test-env-up test-env-down hooks
