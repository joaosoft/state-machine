APP_NAME=logger
GOLANG_VERSION=latest

utest:

itest:

build:
	docker build . --build-arg APP_NAME=$(APP_NAME) --build-arg GOLANG_VERSION=$(GOLANG_VERSION)

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*