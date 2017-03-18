all: format test build

format:
	go fmt .
	go fmt ./iface

test:
	go test .

build:
	go build .
