# Starter 

## Description
Starter - project template for my own projects.
It has a basic structure and configuration for the project I usually use.

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
# Run build
task build
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


