.DEFAULT_GOAL: help

help:
	@# 20s is the width of the first column
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## generate antlr code
	@echo Generate antlr code
	@cd ./sql-grammar && ./generate.sh Go sqlgrammar ../sqlgrammar
