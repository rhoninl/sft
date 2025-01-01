.PHONY: build protoc_gen_go protoc_gen_ts dev

# Build binaries for different platforms
build:
	GOOS=linux GOARCH=amd64 go build -a -o output/sft_linux_amd64 cmd/sft/main.go
	GOOS=linux GOARCH=arm64 go build -a -o output/sft_linux_arm64 cmd/sft/main.go
	GOOS=windows GOARCH=amd64 go build -a -o output/sft_windows_amd64.exe cmd/sft/main.go
	GOOS=windows GOARCH=arm64 go build -a -o output/sft_windows_arm64.exe cmd/sft/main.go

# Generate Go code from protobuf
protoc_gen_go:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/shifu/shifu.proto

# Generate TypeScript code from protobuf
protoc_gen_ts:
	mkdir -p webview/src/proto/proto/shifu
	protoc -I=. proto/shifu/shifu.proto \
		--js_out=import_style=commonjs:webview/src/proto \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:webview/src/proto

# Generate all protobuf code
protoc: protoc_gen_go protoc_gen_ts

# Start development servers
dev:
	@echo "Starting backend server..."
	go run cmd/sft/main.go &
	@echo "Starting frontend server..."
	cd webview && npm start

# Install dependencies
install:
	@echo "Installing Go dependencies..."
	go mod tidy
	@echo "Installing frontend dependencies..."
	cd webview && npm install

# Clean generated files
clean:
	rm -rf output/
	rm -rf webview/src/proto/proto