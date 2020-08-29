.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build
	./gradlew npmBuild

publish: build ## Publish NPM package
	cd build/js/packages/zombie3-browser && npm publish
