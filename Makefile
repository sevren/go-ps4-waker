all: build-ps4-waker

.PHONY: build-ps4-waker
build-ps4-waker:
	CGO=0 go build -a -tags netgo cmd/ps4-waker/main.go

.PHONY: run
run:
	go run cmd/ps4-waker/main.go