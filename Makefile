.PHONY: all clean dist-clean ensure test

all: ensure

ensure: clean
	dep ensure -v -vendor-only
	go build -o server github.com/paksuy/golang-tutor/cmd/server
	go build -o client github.com/paksuy/golang-tutor/cmd/client

test: clean
	dep ensure -v -vendor-only
	go test -v -race -cover -vet=off ./... | tee /tmp/go-test.txt

clean: clean-vendor
	go clean -cache -testcache
	rm -f server client

clean-vendor:
	rm -rf vendor

dist-clean: clean
	rm -rf $(GOPATH)/pkg/*
