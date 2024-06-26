.DEFAULT_GOAL := help

PROJECT_ROOT:=$(shell git rev-parse --show-toplevel)
COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell date +%Y-%m-%d-%H-%M-%S)
APP_NAME := $(shell basename `git rev-parse --show-toplevel`)

# Load env properties , db name, port, etc.
# nb: You can change the default config with `make ENV_CONTEXT=".env.uat" `
ENV_CONTEXT ?= .env.local
ENV_CONTEXT_PATH:=$(PROJECT_ROOT)/$(ENV_CONTEXT)

## Override any default values in the parent .env, with your own
-include $(ENV_CONTEXT_PATH)

MAKE_LIB:=$(PROJECT_ROOT)/scripts
-include $(MAKE_LIB)/tests.mk
-include $(MAKE_LIB)/lints.mk
-include $(MAKE_LIB)/logs.mk
-include $(MAKE_LIB)/tools.mk

#####################
##@ CI
#####################

ci: log test lint scan ## Run CI tasks

install-ci: tools-scan ## install tools for CI only

#####################
##@ Dev
#####################

build: ## build go files
	go build $(GO_BUILD_FLAGS) -o $(APP_NAME)

install: tools-all ## install golang / node dependencies

