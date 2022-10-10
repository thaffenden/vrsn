# vrsn

A single tool for your semantic versioning needs.

## Why?

### Language agnotistic

You can run `vrsn` in a project in any (supported) language and it will work.

Currently supported version files:

- `Cargo.toml` - rust
- `package.json` - javascript, typescript
- `pyproject.toml` - python
- `VERSION` - go, python, various others etc

Don't see your favourite version file type in that list?

See the [CONTRIBUTING guide](./.github/CONTRIBUTING.md) for how to (easily) add
support!

If you're the type of person that jumps between projects in different languages
you don't need to remeber the `yarn` or `poetry` commands for each different
project, just use `vrsn` and get on with the important stuff.

## Install

TBD

## Commands

### `--help`

Run `vrsn --help` for a full up to date usage guide to get started or
`vrsn [command] --help` if you want help with a specific command.

### `check`

Run `vrsn check` to automatically check versions on an existsing git branch.

By default the `check` command can tell if you are on a branch that is not
the base branch (i.e. `main`) and will compare the version file on your current
branch with the version file on the base branch.

This command is super useful for running in CI, just run `vrsn check`, in your
pull request CI and `vrsn` will tell you if the version has been properly
bumped or not.

Name your base branch something other than `main`?
You can use the `--base-branch` flag to specify the name you use.

### `bump`

TBD
