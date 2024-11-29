# SFT

SFT is a command-line tool designed to facilitate Kubernetes operations such as port forwarding, logging, and managing edge devices. This tool is built using Go and leverages the Kubernetes client-go library for interacting with Kubernetes clusters.

## Features

- Port forwarding to Kubernetes pods
- Logging and monitoring of Kubernetes resources
- Management of edge devices
- Installation and uninstallation of telemetry services

## Installation

To install SFT, ensure you have Go installed and set up on your machine. Then, run the following command:
```bash
go install github.com/rhoninl/sft/cmd/sft@latest
```

## Development

### Prerequisites

- Go 1.23.1 or later
- Access to a Kubernetes cluster

### Building from Source

Clone the repository and build the project using:

```bash
git clone https://github.com/rhoninl/sft.git
cd sft
go build -o sft cmd/sft/main.go
```

### Running Tests

To run tests, use:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

