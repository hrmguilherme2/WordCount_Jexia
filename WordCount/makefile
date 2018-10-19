SHELL := /bin/bash
PLATFORM := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)
GOPATH := $HOME/GO
GOBIN := $(GOPATH)/bin
PATH := PATH=$PATH:$HOME/go/bin:$GOPATH/bin

build:
	
	
	export PATH
	dep init
	dep ensure -add github.com/chrislusf/glow/flow
	dep status
	

run:
	go run main.go -input metamorphosis.txt -output Jexia.csv

test:
	go test