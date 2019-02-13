TESTDIRS := ./cmd/... ./internal/...

dep:
	dep ensure

dev:
	go run cmd/blucard/main.go

dynamo:
	awslocal dynamodb create-table --cli-input-json file://create-records-table.json --region us-west-2

localstack:
	DATA_DIR=/tmp/localstack/data localstack start --docker

setup:
	go get -u github.com/golang/dep/cmd/dep
	pip install awscli-local

test:
	go test -race $(TESTDIRS)

lint:
	gometalinter ./...
