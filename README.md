# Introduction

# Dependencies
- Go 1.20+
- Docker

## Go Modules
The easiest way to ensure that Go module dependencies are up to date is by running `go mod tidy`.

# Project Layout

## Top-level Directories
This project layout is loosely derived from the popular [Standard Go Project Layout](https://github.com/golang-standards/project-layout). The top level directories are as follows:

`/cmd` contains the entrypoint(s) of the application with nested directories representing individual binaries.

`/config` holds the application's configuration files.

`/docker` contains docker related items such a `.Dockerfile` or `.docker-compose.yml`. 

`/internal` is used for shared packages and tends to hold the majority of the application code. 

`/scripts` holds any project scripts.

`/test` contains all the project test cases with the exception of unit tests which live side-by-side with the application code.

`/tools` holds any project specific tooling

## Go Package Structure
This project embraces a layer-based strategy for package organization. The following blog post provides a good explanation of what this means: https://www.gobeyond.dev/packages-as-layers/

# Linting
This project uses golangci for linting. Golangci wraps a number of linting tools, which are configured in `.golangci.yml`.
Run the following command to lint the project:
`docker run --rm -v "$(pwd)":/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v`

# Testing

## Running The Tests
The test cases in this project are written in Go. To run them all, execute `go test ./...` at the root of the project. Specify a filepath to instead run a subset of the test cases. For example, `go test ./internal/...` will run all tests located within the `/internal` directory.

Note that certain types of test cases, such as integration tests, may require docker to be running.

## Dependencies
The [testify](https://github.com/stretchr/testify) library is used for assertions and mocks. See `/internal/example_unit_test.go` for examples.

The [Testcontainers](https://testcontainers.com/) framework should be used for external dependencies. See `/test/example_integration_test.go` for examples.

# Troubleshooting

## WSL2 for Windows
You may need to set DOCKER_HOST="npipe:////.//pipe//docker_engine" for test containers to work.