OUTPUT = main # Referenced as Handler in template.yaml
PACKAGED_TEMPLATE = cloudformation_deploy.yaml
S3_BUCKET := $(S3_BUCKET)
STACK_NAME := $(STACK_NAME)
TEMPLATE = cloudformation.yaml

.PHONY: install
install:
	go get ./...

main: ./src/main.go
	go build -o $(OUTPUT) ./src/main.go

.PHONY: clean
clean:
	rm -f $(OUTPUT) $(PACKAGED_TEMPLATE)

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) main

.PHONY: build
build: clean lambda

.PHONY: api
api: build
	sam local start-api -t $(TEMPLATE)

.PHONY: package
package: build
	sam package --template-file $(TEMPLATE) --s3-bucket $(S3_BUCKET) --output-template-file $(PACKAGED_TEMPLATE)

.PHONY: deploy
deploy: package
	sam deploy --stack-name $(STACK_NAME) --template-file $(PACKAGED_TEMPLATE) --capabilities CAPABILITY_IAM