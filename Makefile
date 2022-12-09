export CGO_ENABLED=0
export GOBIN=./bin

APP_VERSION = 0.0.1

.PHONY:
all: install

.PHONY:
install:
	go build -o bin/dsa main.go

.PHONY:
run:
	go run main.go

.PHONY:
solve:
	go test -v -timeout 30s -run ^ github.com/shubham-gaur/dsa-golang/problems/array

.PHONY:
clean:
	rm -rf bin