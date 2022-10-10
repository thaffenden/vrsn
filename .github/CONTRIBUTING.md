# Contributing to `vrsn`

## Setting up for local development

1. Fork the repo.
2. Run `make build` to ensure you can build the binary before making any changes.
3. Checkout a new branch for your changes.
4. Once made don't forget to run `make fmt` and `make lint` to ensure your
changes are inline with the repo's code standards.
5. Make sure you have added any new unit tests for new functionality.
6. Run the tests with `make test`.
7. Create a pull request 🎉

## Adding support for a new version file type

Supported version file types are stored in the `versionFileMap()` function in
`internal/files/get_version.go`.

This function maps the version file name to the function used to get the
version value from the file.

All version funcs should accept an instance of `*bufio.Scanner` and return the
version as a string or an error if the line with the version in cannot be
found.

### Adding unit tests

Unit tests are essential to keep everything working properly as the code
changes. `internal/files/get_version_test.go` contains tests for every
supported file type.
If you are adding a new one you should add a valid example of that file type
to the `internal/files/testdata/all` directory and an example that does not
include the version to `internal/files/testdata/no-version`.

You should  add a unit test for successfully reading the value from the file
and one for throwing an error when the version cannot be found. Both can be
added to the existing table tests in the same format as the current tests.
