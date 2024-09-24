.PHONY: build run

build:
	go build -o bin/load-balancer main.go

run: build
	./bin/load-balancer