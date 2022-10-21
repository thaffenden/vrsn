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

See [Running in Docker](#running-in-docker) for more details.

### Use the CircleCI Orb

For ease of running checks in your CI this repo includes a CircleCI orb.
Just import the orb:

```yaml
orbs:
  vrsn: thaffenden/vrsn@volatile
```

Then use the `check-version` job in your workflow like:

```yaml
workflows:
  build:
    jobs:
      - vrsn/check-version
```

For an example you can look at this repo's CircleCI config, which uses the orb.

See the page in the CircleCI docs for more details on the options you can pass
to make sure the job works for your needs.

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

Run `vrsn bump` to increment the current version file.
It will prompt you to select the bump type and then write the new valid semver
version in your version file.

**Coming soon:** flags so you can avoid the picker.
Just run `vrsn bump --patch` to increment a patch version.

**Coming soon:** auto commit your version bump.
Just run `vrsn bump --commit` to automatically commit your version bump so you
don't need to do anything once the file is bumped.

Customise the commit message with the `--commit-msg` flag if you don't like the
default.

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
