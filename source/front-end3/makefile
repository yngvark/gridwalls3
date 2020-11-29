IMAGE = ghcr.io/yngvark/frontend

.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init: generate-dotenv-file
	@echo "If you don't have npx, run: npm i -g npx"

generate-dotenv-file: ## Generate .env file template
	@echo "BACKEND_URL=ws://localhost:8080/zombie" >> .env
	@echo "PORT=3001" >> .env
	@echo "CERTIFICATE_FILE= #Optional, use 'mkcert localhost' to create certificate and key file" >> .env
	@echo "KEY_FILE= #Optional, use 'mkcert localhost' to create certificate and key file" >> .env

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

