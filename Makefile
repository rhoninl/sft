build:
	GOOS=linux GOARCH=amd64 go build -a -o outoput/sft_linux_amd64 cmd/sft/main.go
	GOOS=linux GOARCH=arm64 go build -a -o outoput/sft_linux_arm64 cmd/sft/main.go
	GOOS=windows GOARCH=amd64 go build -a -o outoput/sft_windows_amd64.exe cmd/sft/main.go
	GOOS=windows GOARCH=arm64 go build -a -o outoput/sft_windows_arm64.exe cmd/sft/main.go