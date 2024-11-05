-include .env

GOPATH:=$(shell go env GOPATH)

dev: build up

up:
	./tmp/app/main

build:
	go build -o ./tmp/app/main ./cmd/app/main.go
