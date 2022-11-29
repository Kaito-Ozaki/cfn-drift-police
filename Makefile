.DEFAULT_GOAL := help
.PHONY: all test clean

all:

init: ## build environment
	npm install serverless
	npm install serverless-go-build
	npm install --save-dev serverless-prune-plugin

invokeCheck: ## invoke lambda function
	serverless build
	sls invoke local -f check

invokeAlert: ## invoke lambda function
	serverless build
	sls invoke local -f alert

test: ## run all test
	go test ./test/...

help: ## Self-documented Makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'