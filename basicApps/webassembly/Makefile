.DEFAULT_GOAL := helper

helper: # Adapted from: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
	@echo "Available targets..."
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Builds the application
	GOOS=js GOARCH=wasm go build -o main.wasm
	@if [ ! -z "wasm_exec.js" ] ; then cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .; fi

run: ## Runs the application
	@if [ $(shell docker ps -qf name=go_webassembly -f status=running | wc -l) != 1 ] ; then docker rm go_webassembly; docker run -ditv $(shell pwd)/:/usr/share/nginx/html/:ro --name go_webassembly -p 80:80 nginx && docker exec -it go_webassembly sed -i "/\}/i \    application/wasm wasm;" /etc/nginx/mime.types; fi

clean: ## Cleans up any old/unneeded
	-rm main.wasm wasm_exec.js
	docker stop go_webassembly && docker rm go_webassembly
