OUTPUT = main # Referenced as Handler in template.yaml
PACKAGED_TEMPLATE = template_deploy.yaml
S3_BUCKET := $(S3_BUCKET)
STACK_NAME := $(STACK_NAME)

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

# .PHONY: api
# api: build
# 	sam local start-api

.PHONY: package
package:
	aws package \
		--s3-bucket $(S3_BUCKET) \
		--output-template-file $(PACKAGED_TEMPLATE) \

.PHONY: deploy
deploy: package
	aws deploy \
		--template-file $(PACKAGED_TEMPLATE) \
		--stack-name $(STACK_NAME) \
		--no-fail-on-empty-changeset \
		--capabilities CAPABILITY_IAM

.PHONY: teardown
teardown: clean
	aws cloudformation delete-stack --stack-name $(STACK_NAME)


# Sync the templates folder with the bucket
#	Only works after deployment
.PHONY: sync-templates
sync-templates:
	aws s3 sync templates/ s3://go-mail-template/