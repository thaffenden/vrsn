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

### Download from GitHub

Find the latest version for your system on the
[GitHub releases page](https://github.com/thaffenden/vrsn/releases).

### Run the Docker container

Get the Docker container from the
[GitHub container registry](https://github.com/thaffenden/vrsn/pkgs/container/vrsn).

```bash
docker pull ghcr.io/thaffenden/vrsn:latest
```

### Use the CircleCI Orb

TBD

See [Running in Docker](## Running in Docker) for more details.

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

## Running in Docker

To run `vrsn` in a docker container you just need to mount the repo as a
volume, and `vrsn` can do it's thing, **however** git's
[safe.directory](https://git-scm.com/docs/git-config/2.35.2#Documentation/git-config.txt-safedirectory)
settings would prevent `vrsn` from being able to use it's git based smarts.

To deal with this a directory called `/repo` is set as a safe directory as part
of the Docker Build process, and is configured as the container's working
directory so it's recommended you use that as the destination of the volume
mount. e.g.:

```bash
docker run --rm -it -v $PWD:/repo vrsn:latest check
```
