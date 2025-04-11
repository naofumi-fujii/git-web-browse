# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build, Lint and Test Commands
- Install: `go get -u github.com/naofumi-fujii/git-web-browse`
- Build: `go build`
- Run all tests: `go test ./...`
- Run a specific test: `go test -run TestName`
- Run test with verbose output: `go test -v ./...`
- Format code: `go fmt ./...`

## Code Style Guidelines
- Formatting: Follow standard Go formatting using `go fmt`
- Imports: Group imports by standard library, then third-party packages
- Error Handling: Check errors and provide meaningful error messages
- Function Names: Use camelCase for function names
- Variable Names: Use descriptive variable names, short names for limited scope
- Maintain test patterns already established in the codebase
- URL handling should properly manage different Git hosting services (GitHub, GitLab, Bitbucket)
- Follow Go idioms like early returns for error handling