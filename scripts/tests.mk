COVER_OUTPUT_RAW := coverage.out
COVER_OUTPUT_HTML := coverage.html

#####################
##@ Tests            
#####################

test: test-unit ## Run all tests

test-unit: ## Run unit tests
	go test -coverprofile $(COVER_OUTPUT_RAW) --short -cover  -failfast ./...

test-watch: ## Run tests in watch mode
	gow test -coverprofile $(COVER_OUTPUT_RAW) --short -cover  -failfast ./...


test-gen-coverage:
	grep -v -E -f ${PWD}/.covignore $(COVER_OUTPUT_RAW) > coverage.filtered.out
	mv coverage.filtered.out $(COVER_OUTPUT_RAW)

test-cover: test-gen-coverage ## generate html coverage report + open
	go tool cover -html=$(COVER_OUTPUT_RAW) -o $(COVER_OUTPUT_HTML)
	open coverage.html

