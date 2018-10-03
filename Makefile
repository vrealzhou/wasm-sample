GO=$(shell which go)

.PHONY: wasm run

wasm:
	mkdir -p build/web
	cp web/* build/web/
	GOARCH=wasm GOOS=js $(GO) build -o build/web/lib.wasm cmd/wasm/main.go

run:
	mkdir -p build/server
	$(GO) run cmd/server/main.go -dir=build/web