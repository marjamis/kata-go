.DEFAULT_GOAL := helper
GIT_COMMIT ?= $(shell git rev-parse --short=12 HEAD || echo "NoGit")
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
TEXT_RED = \033[0;31m
TEXT_BLUE = \033[0;34;1m
TEXT_GREEN = \033[0;32;1m
TEXT_NOCOLOR = \033[0m

helper: # Adapted from: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
	@echo "Available targets..." # @ will not output shell command part to stdout that Makefiles normally do but will execute and display the output.
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

command: clean
	go run main.go -f=./test_data/massive.csv -s=Node1 -d=Node12
	cd go-diagrams && dot -Tpng dijkstra.dot > dijkstra.png
	code go-diagrams/dijkstra.png

build: ## Builds the application

test: ## Builds and then runs tests against the application
	go test -v main.go main_test.go

run: ## Runs the prod version of the application
	$(MAKE) command

clean: ## Cleans up any old/unneeded items
	-rm -r go-diagrams
