# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.11.1

- update dependencies to resolve compatibility issues
- add exclude directives for cloud.google.com/go v0.26.0 and golang.org/x/tools v0.38.0

## v1.11.0

- add variadic constructors to Set, SetEqual, and SetHashCode for initialization with elements
- enhance Add methods to accept variadic parameters for batch operations with single mutex lock
- add comprehensive test coverage with 25 new test cases for variadic functionality
- improve GoDoc comments with performance characteristics and complexity analysis
- add performance comparison documentation between Set implementations
- add usage examples to all constructor documentation
- add go-modtool to Makefile for go.mod formatting
- update Go version to 1.25.2
- update dependencies (bborbe/errors, bborbe/run, and security tools)

## v1.10.2

- add GitHub Actions CI workflow for automated testing
- add GitHub Actions workflows for Claude Code integration
- add golangci-lint configuration
- enhance Makefile with additional quality checks (osv-scanner, gosec, trivy)
- improve test code formatting with golines
- update Go version to 1.24.6
- add comprehensive linting and security scanning tools

## v1.10.1

- add comprehensive documentation comments to all public functions, types, and interfaces
- improve documentation following Go doc best practices guidelines

## v1.10.0

**BREAKING CHANGES:**
- ContainsAll behavior changed: now only checks if all elements of second argument are present in first argument (superset check), instead of bidirectional equality check

## v1.9.1

- add tests
- go mod update

## v1.9.0

- remove vendor
- go mod update

## v1.8.0

- add StreamList
- go mod update

## v1.7.0

- add join
- go mod update

## v1.6.0

- add compare

## v1.5.0

- add Map helper
- go mod update

## v1.4.1

- go mod update

## v1.4.0

- add length method to sets
- go mod update

## v1.3.2

- go mod update

## v1.3.1

- add SetHashCode
- add SetEqual

## v1.3.0

- add Set

## v1.2.0

- add Equal
- add ContainsAll

## v1.1.0

- add ChannelFnCount

## v1.0.0

- Initial Version
