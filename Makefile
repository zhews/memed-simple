test:
	go test ./...

hooks:
	cd .git/hooks; ln -s ../../githooks/* .

.PHONY: test hooks
