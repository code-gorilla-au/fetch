.DEFAULT_GOAL := help

test: ## Test
	go test -cover -short -failfast ./...


scan: fmt vet ## run security scan
	gosec ./...

fmt: ## run go fmt
	go fmt ./...

vet: ## run go vet
	go vet ./...

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)	