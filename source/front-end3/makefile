IMAGE = ghcr.io/yngvark/frontend

.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init:
	npm i -g npx

build-docker: ## -
	docker build -t $(IMAGE) .

run-docker: build-docker ## -
	docker run -d --name zombie-frontend --rm -p 3001:3001 $(IMAGE)

push: build-docker ## -
	docker push $(IMAGE)

run: ## -
	npm run dev

build: ## -
	npm run build

run-prod: ## Run the webapp through Node, as it would in production. Run make build before running this.
	npx nodemon node_server.ts

.PHONY: test
test: ## -
	npm run test

watch: ## -
	npm run watch

