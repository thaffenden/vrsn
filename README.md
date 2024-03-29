<h1 align="center">vrsn</h1>

<p align="center">
  <em>A single tool for <strong>all</strong> of your semantic versioning needs.</em>
</p>

![vrsn-demo](https://user-images.githubusercontent.com/14163530/204849040-39c8d4e2-80ef-4722-a04a-14a575ea4638.gif)

## Contents

- [Why](#why)
- [Install](#install)
  - [Download from GitHub](#download-from-github)
  - [Build it locally](#build-it-locally)
  - [Run the Docker container](#run-the-docker-container)
  - [Use the CircleCI orb](#use-the-circleci-orb)
- [Commands](#commands)
- [Running in Docker](#running-in-docker)

## Why?

### Language agnostic

You can run `vrsn` in a project in any (supported) language and it will work.

Currently supported version files:

| File | Languages |
| --- | --- |
| `build.gradle`, `build.gradle.kts` | ![Java](https://img.shields.io/badge/java-%23ED8B00.svg?style=for-the-badge&logo=java&logoColor=white) ![Kotlin](https://img.shields.io/badge/kotlin-%237F52FF.svg?style=for-the-badge&logo=kotlin&logoColor=white) |
| `Cargo.toml` | ![Rust](https://img.shields.io/badge/rust-%23000000.svg?style=for-the-badge&logo=rust&logoColor=white) |
| `CMakeLists.txt` | ![C++](https://img.shields.io/badge/c++-%2300599C.svg?style=for-the-badge&logo=c%2B%2B&logoColor=white) |
| `package.json` | ![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white) ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E) |
| `pyproject.toml` | ![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) |
| `setup.py` | ![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) |
| `VERSION` | ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) + more |

Don't see your favourite version file type in that list?
See the [CONTRIBUTING guide](./.github/CONTRIBUTING.md) for how to (easily) add
support!

If you're the type of person that jumps between projects in different languages
you don't need to remember the `yarn` or `poetry` commands for each different
project, just use `vrsn` and get on with the important stuff.

### Simple CI checks

Ensuring you properly version releases is important.

I've had to write semantic version checks in CI pipelines in different ways for
different languages in different jobs. Now I can just use `vrsn` and not have
to worry about solving the same problems again.

## Install

### Download from GitHub

Find the latest version for your system on the
[GitHub releases page](https://github.com/thaffenden/vrsn/releases).

### Build it locally

If you have go installed, you can clone this repo and run:

```bash
make install
```

This will build the binary and then copy it to `/usr/bin/vrsn` so it will be
available on your path. Nothing more to it.

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
      - vrsn/check-version:
            filters:
              branches:
                ignore:
                  - main
```

For an example you can look at this repo's [CircleCI config](./.circleci/config.yml)
which uses the orb.

See the [CircleCI orb docs](https://circleci.com/developer/orbs/orb/thaffenden/vrsn)
for more specifics on how to customise the orb jobs to best suite your needs.

The orb is semantically versioned using the same number as the `vrsn` binary
and Docker container, so you can pin a specific version in your CI config or
use the `volatile` tag to always get the latest version of `vrsn`.

## Commands

### `--help`

Run `vrsn --help` for a full up to date usage guide to get started or
`vrsn [command] --help` if you want help with a specific command.

### `check`

Run `vrsn check` to automatically check versions on an existing git branch.

By default the `check` command can tell if you are on a branch that is not
the base branch (i.e. `main`) and will compare the version file on your current
branch with the version file on the base branch.

This command is super useful for running in CI, just run `vrsn check`, in your
pull request CI and `vrsn` will tell you if the version has been properly
bumped or not.

Name your base branch something other than `main`?
You can use the `--base-branch` flag to specify the name you use.

Want to run it from somewhere other than the root of your git repo? You can
use the `--was` and `--now` flags to pass in values from wherever you need to
grab them:

```bash
vrsn check --was $(<function to get previous value>) --now $(<function to get current value>)
```

You can use the `--file` flag to point at a file that is not in the root of the
git repo (like in a monorepo with independantly versioned services), e.g.:

```bash
vrsn check --file './services/service-name/VERSION'
```

### `bump`

Run `vrsn bump` to increment the current version file.
It will prompt you to select the bump type and then write the new valid semver
version in your version file.

If you want to avoid the interactive picker you can pass the increment level as
an argument to the `bump` command, e.g.:

```bash
vrsn bump patch
```

Want to automatically commit the version bump? Just use the `--commit` flag. 🙌

Don't like the default commit message? Provide your own custom one with
`--commit-msg`.

You can use the `--file` flag to point at a file that is not in the root of the
git repo (like in a monorepo with independantly versioned services), e.g.:

```bash
vrsn bump --file './services/service-name/VERSION'
```

This approach allows you to easily increment multiple versions in bulk, just
write a script to iterate over each service that needs bumping and use the
command `vrsn bump patch --file ./services/$SERIVCE_NAME/VERSION`.

## Running in Docker

To run `vrsn` in a docker container you just need to mount the repo as a
volume, and `vrsn` can do it's thing, **however** git's
[safe.directory](https://git-scm.com/docs/git-config/2.35.2#Documentation/git-config.txt-safedirectory)
settings would prevent `vrsn` from being able to use it's git based smarts 🧠.

To deal with this a directory called `/repo` is set as a safe directory as part
of the Docker build process, and is configured as the container's working
directory so it's recommended you use that as the destination of the volume
mount. e.g.:

```bash
docker run --rm -it -v $PWD:/repo vrsn:latest check
```
