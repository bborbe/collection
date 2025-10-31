# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.17.0
- Add Intersect function to find common elements between two slices
- Intersect preserves order from first slice and handles duplicates automatically
- Add comprehensive test suite with 21 tests covering strings, integers, and custom types
- Update Go version from 1.25.2 to 1.25.3
- Update dependencies (github.com/bborbe/run, github.com/onsi/ginkgo/v2, and others)

## v1.16.0

- Add MarshalJSON and UnmarshalJSON methods to Set interface for JSON serialization
- Add MarshalJSON and UnmarshalJSON methods to SetEqual interface for JSON serialization
- Add MarshalJSON and UnmarshalJSON methods to SetHashCode interface for JSON serialization
- Support JSON arrays for all Set types (primitives, structs, maps, nested objects)
- Add comprehensive test coverage for JSON marshaling (39 new tests)
- Add tests for Sets embedded in structs with single and multiple Set fields
- Add tests for nested structs containing Set fields
- Add tests for round-trip JSON marshal/unmarshal operations

## v1.15.0

- Add UnmarshalText and MarshalText methods to Set interface for text encoding/decoding support
- Add comprehensive GoDoc comments for UnmarshalText and MarshalText methods

## v1.14.2

- Fix Set.UnmarshalText to properly support custom string-based types using unsafe.Pointer conversion
- Add comprehensive tests for Set with custom string types and type aliases

## v1.14.1

- Add MarshalText implementation for Set to enable automatic parsing with github.com/bborbe/argument

## v1.14.0

- Add ParseSetFromStrings function for converting string slices to Set with ~string constraint
- Add ParseSetFromString function for parsing comma-separated strings to Set with ~string constraint
- Add UnmarshalText implementation for Set to enable automatic parsing with github.com/bborbe/argument
- Add comprehensive test coverage for UnmarshalText and parsing functions

## v1.13.1

- refactor Strings() method to eliminate code duplication using elementToString helper
- refactor String() method to use formatSetString helper for consistent output
- optimize memory usage: change map[T]bool to map[T]struct{} in ContainsAny and ContainsAll
- standardize map lookup pattern to idiomatic two-value form
- add package-level documentation (doc.go) describing library features and architecture
- enhance HasHashCode documentation with security guidance for hash collision prevention
- add performance notes to ContainsAny and ContainsAll functions
- update copyright headers to 2025
- improve test coverage to 99.3%

## v1.13.0

- add ContainsAny function for checking if any element from one slice exists in another
- add ContainsAll and ContainsAny methods to Set, SetEqual, and SetHashCode interfaces
- add Strings() method to all Set implementations for sorted string slice output
- enhance Remove methods to accept variadic parameters for batch removal with single mutex lock
- improve String() methods to use sorted output for deterministic debugging
- optimize Strings() implementation with type switch for better performance (Stringer, string, default)
- add comprehensive test coverage for new ContainsAll and ContainsAny methods

## v1.12.0

- add String() method to Set, SetEqual, and SetHashCode interfaces for human-readable output
- implement efficient string representation using strings.Builder
- add comprehensive test coverage for String() methods with 100% code coverage
- document non-deterministic ordering behavior for map-based sets
- document insertion order preservation for SetEqual string representation

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
