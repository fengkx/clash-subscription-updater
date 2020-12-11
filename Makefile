build: version
	mkdir -p ./build
	go build -o build
build_all: version
	mkdir -p ./build
	env GOOS=linux GOARCH=amd64    go build -o ./build/clash-subscription-updater-linux-amd64
	env GOOS=linux GOARCH=arm64    go build -o ./build/clash-subscription-updater-linux-arm64
	env GOOS=linux GOARCH=arm      go build -o ./build/clash-subscription-updater-linux-arm
	env GOOS=linux GOARCH=mips     go build -o ./build/clash-subscription-updater-linux-mips
	env GOOS=linux GOARCH=mipsle   go build -o ./build/clash-subscription-updater-linux-mipsle
	env GOOS=linux GOARCH=mips64   go build -o ./build/clash-subscription-updater-linux-mips64
	env GOOS=linux GOARCH=mips64le go build -o ./build/clash-subscription-updater-linux-mips64le
	env GOOS=darwin GOARCH=amd64   go build -o ./build/clash-subscription-updater-macos-amd64

version:
	@sh  version.sh

.PHONY: build