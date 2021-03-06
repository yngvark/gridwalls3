IMAGE = ghcr.io/yngvark/frontend

.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init: generate-dotenv-file
	@echo "If you don't have npx, run: npm i -g npx"

generate-dotenv-file: ## Generate .env file template
	@echo "BACKEND_URL=ws://localhost:8080/zombie" >> .env
	@echo "PORT=3000" >> .env
	@echo "CERTIFICATE_FILE= #Optional, use 'mkcert localhost' to create certificate and key file" >> .env
	@echo "KEY_FILE= #Optional, use 'mkcert localhost' to create certificate and key file" >> .env

build-docker: ## -
	docker build -t $(IMAGE) .

run-docker: build-docker ## -
	docker kill zombie-frontend || true
	docker run \
		-d \
		-e PORT=3000 \
		-e BACKEND_URL=ws://localhost:8080/zombie \
		--name zombie-frontend \
		--rm \
		-p 30000:3000 \
		$(IMAGE)

push: build-docker ## -\
	docker push $(IMAGE)

run-watch: ## -
	npm run watch

run-no-watch: ## -
	npm run dev

build: ## -
	npm run build

run: ## -
	# npm run watch runs rollup watch, so our bundled javascript (dist/game.js) stays up to date
	# - npx launches the server, which hosts the app like it would in production, which uses node. It watches the server
	# directory and automatically refreshes the server.
	npm run watch & npx nodemon server/node_server.ts

.PHONY: test
test: ## -
	npm run test
