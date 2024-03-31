build:
	rm -rf build
	mkdir -p build
	cp -r client/assets/. build/
	env GOOS=js GOARCH=wasm go build -o build/game.wasm -tags dev client/cmd/client/client.go

buildprod:
	rm -rf build
	mkdir -p build
	cp -r client/assets/. build/
	env GOOS=js GOARCH=wasm go build -o build/game.wasm -tags prod client/cmd/client

try:
	@trap 'kill $(shell lsof -t -i:8080)' SIGINT; \
	go run github.com/hajimehoshi/wasmserve@latest ./client/cmd/client &
ifeq ($(OS),Windows_NT)
	explorer.exe http://localhost:8080
else
	xdg-open http://localhost:8080
endif