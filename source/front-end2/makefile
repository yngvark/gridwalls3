.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-docker: ## -
	docker build -t gridwalls_zombie .

run-docker: build-docker ## -
	docker run --rm -p 30000:3000 gridwalls_zombie

run: ## -
	npm --prefix src run dev
