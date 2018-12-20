dep:
	dep ensure

dev:
	go run cmd/blucard/main.go

dynamo:
	dynamodb create-table --cli-input-json file://create-records-table.json --region us-west-2

localstack:
	localstack start --docker

setup:
	go get -u github.com/golang/dep/cmd/dep
