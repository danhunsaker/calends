# Calends

[![Software License](https://img.shields.io/github/license/danhunsaker/calends.svg?style=for-the-badge)](LICENSE)
[![Main Docs](https://img.shields.io/readthedocs/calends.svg?label=main+docs&style=for-the-badge)](https://calends.readthedocs.io/)
[![GoDoc Reference](https://img.shields.io/badge/GoDoc-reference-brightgreen.svg?style=flat-square)](https://godoc.org/github.com/danhunsaker/calends)
[![Gitter Chat](https://img.shields.io/gitter/room/danhunsaker/calends.svg?style=flat-square)](https://gitter.im/danhunsaker/calends)
[![Total Downloads](https://img.shields.io/github/downloads/danhunsaker/calends/total.svg?style=flat-square)](https://github.com/danhunsaker/calends/releases)

[![Latest Stable Version](https://img.shields.io/github/release/danhunsaker/calends.svg?style=for-the-badge)](https://github.com/danhunsaker/calends/releases)
[![GitHub Release Date](https://img.shields.io/github/release-date/danhunsaker/calends.svg?style=for-the-badge)](https://github.com/danhunsaker/calends)
[![Github commits (since latest release)](https://img.shields.io/github/commits-since/danhunsaker/calends/latest.svg?style=flat-square)](https://github.com/danhunsaker/calends)
[![GitHub last commit](https://img.shields.io/github/last-commit/danhunsaker/calends.svg?style=flat-square)](https://github.com/danhunsaker/calends)

[![Maintenance Status](https://img.shields.io/maintenance/yes/2023.svg?style=flat-square)](https://github.com/danhunsaker/calends)
[![GitHub branch checks state](https://img.shields.io/github/checks-status/danhunsaker/calends/main?style=flat-square)](https://github.com/danhunsaker/calends/actions)
[![Codecov coverage](https://img.shields.io/codecov/c/github/danhunsaker/calends.svg?style=flat-square)](https://codecov.io/gh/danhunsaker/calends)
[![Go Report Card](https://goreportcard.com/badge/github.com/danhunsaker/calends?style=flat-square)](https://goreportcard.com/report/github.com/danhunsaker/calends)
[![Libraries.io Dependency Check](https://img.shields.io/librariesio/github/danhunsaker/calends.svg?style=flat-square)](https://libraries.io/github/danhunsaker/calends)

[![Code Climate coverage](https://img.shields.io/codeclimate/coverage-letter/danhunsaker/calends.svg?style=flat-square)](https://codeclimate.com/github/danhunsaker/calends)
[![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability/danhunsaker/calends.svg?style=flat-square)](https://codeclimate.com/github/danhunsaker/calends)
[![Code Climate technical debt](https://img.shields.io/codeclimate/tech-debt/danhunsaker/calends.svg?style=flat-square)](https://codeclimate.com/github/danhunsaker/calends)
[![Code Climate issues](https://img.shields.io/codeclimate/issues/danhunsaker/calends.svg?style=flat-square)](https://codeclimate.com/github/danhunsaker/calends)

[![Crowdin](https://d322cqt584bo4o.cloudfront.net/calends/localized.svg)](https://crowdin.com/project/calends)
[![Liberapay receiving](https://img.shields.io/liberapay/receives/danhunsaker.svg?style=flat-square)](https://liberapay.com/danhunsaker/)

A library for handling dates and times across arbitrary calendar systems

## Features

More information about each of these features is available in [the full
documentation][full].

- [x] Large range and high precision
- [x] Supports date (and time) values in multiple calendar systems:
  - [x] **Unix time**
  - [x] **[TAI64][]**
    - [x] Automatic calculation of leap second offsets
    - [ ] Estimation of undefined past and future leap second insertions
    - [ ] Automatic updates for handling leap second insertions
  - [x] **Gregorian**
    - [ ] Disconnect from native `time.Time` implementation, and its limitations
  - [ ] **Julian**
  - [x] **Julian Day Count**
  - [ ] **Hebrew**
  - [ ] **Persian**
  - [ ] **Chinese**
  - [ ] **Meso-American**
  - [ ] **Discordian**
  - [x] **Stardate**
- [x] Encodes both time spans and instants in a single interface
- [x] Supports calculations and comparisons on spans and instants
- [ ] Conversion to/from native date/time types
- [ ] Geo-temporally aware
- [ ] Time zone support
- [x] Well-defined interfaces for extending the library
- [x] Shared library (`.so`/`.dll`/`.dylib`)
- [X] WebAssembly binary

## Installation

The steps here will vary based on which programming language(s) you're using.

For Golang, simply run `go get github.com/danhunsaker/calends`, and then place
`"github.com/danhunsaker/calends"` in the `import` wherever you intend use it.

For other languages, refer to [the full documentation][full].

## Usage

Usage data has been moved to [the full documentation][full].

## Calendar Systems

Currently supported calendar systems, and the options available for each, are
listed in [the full documentation][full]. Also provided there are the docs for
how to add your own.

## Contributing

Pull requests are always welcome! That said, please be open to discussing the PR
content, and possibly revising it if requested. Not all requests can be merged,
and not all changes are desired.

Or, you can contribute some money, instead! Check out [my
Patreon](https://www.patreon.com/DanHunsaker) for options, there. Other options
will likely be added for one-time donations in the future.

## Security Reporting

Report all security-related issues to [dan (dot) hunsaker (plus) calends (at)
gmail](mailto:dan.hunsaker+calends@gmail.com), and use PGP or GPG protections on
your message (the account's key is `44806AB9`, or you can look it up by the
email address). Security issues will be addressed internally before making any
vulnerability announcements.

[GitHub]:https://github.com/danhunsaker/calends
[TAI64]:http://cr.yp.to/libtai/tai64.html
[full]:https://calends.readthedocs.io/
