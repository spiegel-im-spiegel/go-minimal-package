# go-minimal-package -- Minimal package for development of Go

[![check vulns](https://github.com/spiegel-im-spiegel/go-minimal-package/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/go-minimal-package/actions)
[![lint status](https://github.com/spiegel-im-spiegel/go-minimal-package/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/go-minimal-package/actions)
[![lint status](https://github.com/spiegel-im-spiegel/go-minimal-package/workflows/build/badge.svg)](https://github.com/spiegel-im-spiegel/go-minimal-package/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/go-minimal-package/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/go-minimal-package.svg)](https://github.com/spiegel-im-spiegel/go-minimal-package/releases/latest)

This package is required Go 1.16 or later.

## Build and Install

```
$ go install github.com/spiegel-im-spiegel/go-minimal-package@latest
```

## Binaries

See [latest release](https://github.com/spiegel-im-spiegel/go-minimal-package/releases/latest).

## Usage

```
$ go-minimal-package -h
Usage:
  go-minimal-package [flags]
  go-minimal-package [command]

Available Commands:
  help        Help about any command
  version     Print the version number

Flags:
  -h, --help   help for go-minimal-package

Use "go-minimal-package [command] --help" for more information about a command.
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

[go-minimal-package]: https://github.com/spiegel-im-spiegel/go-minimal-package "spiegel-im-spiegel/go-minimal-package: Minimal package for development of Go"
