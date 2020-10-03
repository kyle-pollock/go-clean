# go-clean

![Go](https://github.com/kyle-pollock/go-clean/workflows/Go/badge.svg?branch=master)

Example project of clean architecture in Go.

## Local development setup

### Config

1. Install the [config-wizard](https://github.com/spowelljr/config-wizard)
2. `cp .env.example .env` for a starting point and edit the contents of `.env` with configurations.
2. Run `source wizard -f` to inject the contents of `.env` as [environment variables](https://en.wikipedia.org/wiki/Environment_variable).

## Build

`make`

## Run

`make run`

## Test

`make test` to run all tests including long running tests such as integration tests.

`make test-short` to run only short tests such as unit tests.
