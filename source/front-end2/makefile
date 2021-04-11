IMAGE = ghcr.io/yngvark/gridwalls_zombie

.PHONY: help
help: ## Print this menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-docker: ## -
	docker build -t $(IMAGE) .

run-docker: build-docker ## -
	docker run --name zombie-frontend --rm -p 30000:3000 $(IMAGE)

push:
	docker push $(IMAGE)

run: ## -
	npm --prefix src run dev
