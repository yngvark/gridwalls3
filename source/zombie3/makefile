SHELL=/bin/bash

.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build
	./gradlew npmBuild

publish: build ## Publish NPM package
	echo PATH: $(PATH)
	echo "Prerequisite, run:"
	echo "npm login --registry=https://npm.pkg.github.com"
	echo "with password=<github token from https://github.com/settings/tokens>"
	echo
	cd build/js/packages/zombie3-browser && npm publish
