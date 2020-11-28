IMAGE = ghcr.io/yngvark/frontend

.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-docker: ## -
	docker build -t $(IMAGE) .

run-docker: build-docker ## -
	docker run --name zombie-frontend --rm -p 30000:3000 $(IMAGE)

push: build-docker ## -
	docker push $(IMAGE)

run: ## -
	npm run dev

.PHONY: test
test: ## -
	npm run test

watch: ## -
	npm run watch

