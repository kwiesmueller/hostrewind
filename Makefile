all: test install
install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/hostrewind/*.go
vet:
	go tool vet .
	go tool vet --shadow .
lint:
	golint -min_confidence 1 ./...
errcheck:
	errcheck -ignore '(Close|Write)' ./...
check: lint vet errcheck
format:
	find . -name "*.go" -exec gofmt -w "{}" \;
	goimports -w=true .
prepare:
	go get -u github.com/golang/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/kisielk/errcheck
	go get -u github.com/golang/glog
