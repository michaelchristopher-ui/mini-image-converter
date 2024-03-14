.PHONY: deploy build clean config run

all:
	make setup
	make build
	make run

setup:
	go mod vendor

build:
	go build -v -o main ./cmd/app/main.go

run:
	./main --configpath ./cmd/app/config.yaml

test: test.unittest

test.unittest:
ifeq ($(cover),true)
	mkdir -p coverage
	go test -v -short -coverprofile coverage/unittest_cover.out -race ./...
else
	go test -v -short ./...
endif
