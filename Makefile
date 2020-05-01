all: build-ps4-waker

.PHONY: build-ps4-waker
build-ps4-waker:
	CGO_ENABLED=0 go build -a -tags netgo -o bin/ps4-waker cmd/ps4-waker/ps4-waker.go

.PHONY: run
run:
	go run cmd/ps4-waker/ps4-waker.go