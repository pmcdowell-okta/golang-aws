
GOPATH=$(shell pwd)


setup:
	go get "github.com/aws/aws-lambda-go/events"
	go get "github.com/aws/aws-lambda-go/lambda"

clean:
	rm -rf bin/*

build:
	mkdir -p bin
	@GOPATH=$(GOPATH) GOOS=linux GOARCH=amd64 go build -o ./bin/echo2 echo2.go
	@GOPATH=$(GOPATH) GOOS=linux GOARCH=amd64 go build -o ./bin/main main.go
	@GOPATH=$(GOPATH) GOOS=linux GOARCH=amd64 go build -o ./bin/webpage webpage.go


deploy:
	sls deploy


