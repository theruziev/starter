# Starter 

[![Go Report Card](https://goreportcard.com/badge/github.com/theruziev/starter)](https://goreportcard.com/report/github.com/theruziev/starter)
[![Go Version](https://img.shields.io/github/go-mod/go-version/theruziev/starter)](https://github.com/theruziev/starter)
[![License](https://img.shields.io/github/license/theruziev/starter)](https://github.com/theruziev/starter/blob/main/LICENSE)

## Description
Starter is a robust project template for Go applications. It provides a solid foundation with essential components like command-line interface, HTTP server, logging, health checks, and version information.

This template follows Go best practices and has a clean, modular structure that makes it easy to extend and maintain.

## Features
- Clean architecture with clear separation of concerns
- Command-line interface with environment variable support
- HTTP server with health checks and graceful shutdown
- Structured logging with context propagation
- Task-based build system
- Multi-stage Docker build with minimal final image
- Ready-to-use fly.io deployment configuration

## Getting Started

### Prerequisites
- Go 1.18 or higher
- [go-task](https://taskfile.dev/) for running tasks

### Installation
1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/your-project.git
   cd your-project
   ```

2. Update the module name
   ```bash
   # Edit go.mod to change the module name
   # Update PACKAGE_PREFIX in Taskfile.yml
   ```

3. Install dependencies
   ```bash
   go mod tidy
   ```

## Dependencies
- [go-task](https://taskfile.dev/) - A task runner / simpler Make alternative written in Go
- [alecthomas/kong](https://github.com/alecthomas/kong) - is a command-line parser for Go
- [joho/godotenv](https://github.com/joho/godotenv) - loads environment variables from .env file
- [go-chi/chi](github.com/go-chi/chi/v5) - lightweight, idiomatic and composable router for building Go HTTP services
- [gRPC-go](https://grpc.io/docs/guides/health-checking) - gRPC Health Checking Protocol
- [uber-go/zap](https://go.uber.org/zap) - Blazing fast, structured, leveled logging in Go

## Project structure
```
├── cmd - application entry points
├── internal - internal packages
│   ├── app - in here we create application runners
│   │   └── server - server runner
│   └── pkg - internal packages
│       ├── closer - closer package provide a way to close multiple resources
│       ├── logx - logx package provide a way to log messages
│       ├── info - info package provide a information about version and git commit
│       └── healthcheck - healthcheck package provide a healthcheck functionality
├── pkg - public packages
├── .theruziev - project related files (e.g. Docker, taskfile.yml, etc.)
└── .bin - binaries
```

## Usage 
```bash
# Run tests
task test
# Run linter
task lint
# Run build 
task build
```
## Example server
I use fly.io for deployment.
- [Main](https://theruziev-starter.fly.dev/)
- [Version](https://theruziev-starter.fly.dev/version)
- [Readiness probe](https://theruziev-starter.fly.dev/ready)
- [Liveness probe](https://theruziev-starter.fly.dev/live)

## AI Assistant Conventions
This project includes conventions and guidelines for AI assistants. See [CONVENTIONS.md](CONVENTIONS.md) for details.
