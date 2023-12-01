.PHONY: help

help:
	@echo 'Usage: make <target> [options]'
	@echo
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

create-day: ## Usage: make create-day l=golang y=2022 d=1
	@if [ -z "$(l)" -o -z "$(y)" -o -z "$(d)" ]; then \
		echo "Missing required variables"; \
		exit 1; \
	fi

	@cd scripts/createDays &&	go run main.go -l $(l) -y $(y) -d $(d)

run: ## Usage: make run l=golang y=2022 d=1 p=1
	@if [ -z "$(l)" -o -z "$(y)" -o -z "$(d)" -o -z "$(p)" ]; then \
		echo "Missing required variables"; \
		exit 1; \
	fi

	@if [ "$(l)" = "golang" ]; then \
		cd $(y)/$(l)/day$(d) && go run main.go -p $(p); \
		exit 0; \
	fi

	@if [ "$(l)" = "rust" ]; then \
		cd $(y)/$(l)/day$(d) && cargo run -- $(p); \
		exit 0; \
	fi

get-prompt: ## Usage: make get-prompt l=golang y=2022 d=1
	@if [ -z "$(l)" -o -z "$(y)" -o -z "$(d)" ]; then \
		echo "Missing required variables"; \
		exit 1; \
	fi

	@cd scripts/createDays && go run main.go -l $(l) -y $(y) -d $(d) -po 1

test: ## TODO
	@echo "TODO"
