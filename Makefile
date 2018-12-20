dep:
	dep ensure

dev:
	go run cmd/blucard/main.go

setup:
	go get -u github.com/golang/dep/cmd/dep