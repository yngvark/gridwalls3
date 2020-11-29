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

run-watch: ## -
	npm run watch

run-no-watch: ## -
	npm run dev

build: ## -
	npm run build

run: ## -
	# We need a bit crazy setup to run both a server and a client:
	# - npx launches the server, which hosts the app like it would in production, which uses node. It watches the server
	# directory and automatically refreshes the server.
	# - rollup (npm run watch) makes sure that the bundled javascript (dist/game.js) file made by rollup (js bundler)
	# stays up to date. It watches the src folder and automatically creates a bundled dist/game.js.
	#
	# But wait, there's more:
	# - npx is hosted on http://localhost:3001, or http://localhost:3001 if you have this enabled (in .env file)
	# - rollup is hosted on http://localhost:3000. It cannot work alone, as it needs to call an endpoint /config, which
	# is hosted by npx, to get environmental variables. Maybe I should do all this some other way.
	npx nodemon server/node_server.ts & npm run watch

.PHONY: test
test: ## -
	npm run test

watch: ## -
	npm run watch

