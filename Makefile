S3_BUCKET := $(S3_BUCKET)
STACK_NAME := $(STACK_NAME)
REGION := us-east-1

init:
	aws s3api create-bucket --bucket $(S3_BUCKET)
.PHONY: init

build:
	sam build
.PHONY: build

package:
	sam package --s3-bucket $(S3_BUCKET) --region $(REGION)
.PHONY: package

deploy:
	sam deploy --s3-bucket $(S3_BUCKET) --region $(REGION) \
	--stack-name $(STACK_NAME) --capabilities CAPABILITY_IAM
.PHONY: deploy

delete:
	aws cloudformation delete-stack --region $(REGION) --stack-name $(STACK_NAME)
.PHONY: delete
