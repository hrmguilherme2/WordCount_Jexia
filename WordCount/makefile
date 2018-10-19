
build:
	
	echo "$${PATH}"
	dep init
	dep ensure -add github.com/chrislusf/glow/flow
	dep status
	

run:
	go run main.go -input metamorphosis.txt -output Jexia.csv

test:
	go test

deps:
	ulimit -n 4096
	go get -u github.com/golang/dep/cmd/dep
