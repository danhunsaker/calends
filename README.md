# Calends

[![Software License](https://img.shields.io/github/license/danhunsaker/calends.svg?style=flat-square)](LICENSE)
[![Gitter](https://img.shields.io/gitter/room/danhunsaker/calends.svg?style=flat-square)](https://gitter.im/danhunsaker/calends)

[![Latest Stable Version](https://img.shields.io/github/release/danhunsaker/calends.svg?label=stable&style=flat-square)](https://github.com/danhunsaker/calends/releases)
[![Latest Unstable Version](https://img.shields.io/github/release/danhunsaker/calends/all.svg?label=unstable&style=flat-square)](https://github.com/danhunsaker/calends/releases)
[![Build Status](https://img.shields.io/travis/danhunsaker/calends.svg?style=flat-square)](https://travis-ci.org/danhunsaker/calends)
[![Codecov](https://img.shields.io/codecov/c/github/danhunsaker/calends.svg?style=flat-square)](https://codecov.io/gh/danhunsaker/calends)
[![Total Downloads](https://img.shields.io/github/downloads/danhunsaker/calends/total.svg?style=flat-square)](https://github.com/danhunsaker/calends/releases)

A library for handling dates and times across arbitrary calendar systems

## Features

- [x] Large range and high precision.

  Calends understands dates 2<sup>62</sup> seconds into the future or past, in
  units as small as 10<sup>-45</sup> seconds – that's over 146 billion years
  into the past or future (146 138 512 313 years, 169 days, 10 hours, 5 minutes,
  and 28 seconds from CE 1970 Jan 01 00:00:00 TAI, Gregorian), at resolutions
  smaller than Planck Time (54x10<sup>−45</sup> seconds, and the smallest
  meaningful duration even on the quantum scale). That encompasses well beyond
  the expected lifespan of the Universe, at resolutions enough to represent
  quantum events.

- [x] Supports parsing, formatting, and calculating offsets between date (and
  time) values in multiple calendar systems.

  Supported out of the box are the following (all systems are proleptic
  \[extrapolated beyond the officially-defined limits] unless specified
  otherwise):

  - [x] **Unix time**: A count of the number of seconds since CE 1970 Jan 01
    00:00:00 UTC

  - [x] **[TAI64][]**: Essentially Unix time plus 2<sup>62</sup>, but using TAI
    seconds instead of UTC seconds, so times can be converted unambiguously (UTC
    uses leap seconds to keep the solar zenith at noon, while TAI is a simple,
    unadjusted count). Calends supports an extended version of this spec, with
    three more components, to encode out to 45 places instead of just 18; this
    is also actually the internal time scale used by Calends itself, which is
    how it can support such a broad range of dates at such a high resolution.

    - [ ] Automatic updates for handling leap second insertions

    - [ ] Estimation of undefined past and future leap second insertions

  - [x] **Gregorian**: The current international standard calendar system

  - [ ] **Julian**: The previous version of the Gregorian calendar

  - [ ] **Julian Day Count**: A count of days since BCE 4713 Jan 01 12:00:00 UTC
    on the proleptic Julian Calendar

  - [ ] **Hebrew**: ...

  - [ ] **Persian**: ...

  - [ ] **Chinese**: Several variants

  - [ ] **Meso-American**: Commonly called Mayan, but used by several cultures
    in the region

  - [ ] **Discordian**: ...

  - [ ] **Stardate**: Yes, the ones from Star Trek&trade;; several variants exist

- [x] Encodes both time spans (`start != end`, `duration != 0`) and instants
  (`start == end`, `duration == 0`) in a single interface.

  The library treats the time values it encodes as `[start, end)` sets (that is,
  the `start` point is included in the range, as is every point between `start`
  and `end`, but the `end` point itself is _not_ included in the range). This
  allows `duration` to accurately be `end - start` in all cases. (And yes, that
  also means you can create spans with `duration < 0`.)

- [x] Supports calculations and comparisons on spans and instants.

  Addition, subtraction, intersection, combination, gap calculation, overlap
  detection, and similar operations are all supported directly on Calends
  values.

- [ ] Conversion to/from native date/time types.

  While this is possible by using a string representation as an intermediary, in
  either direction, some data and precision is lost in such a conversion.
  Instead, Calends supports conversion to and from such types directly,
  preserving as much data and accuracy as each native type provides.

- [ ] Geo-temporally aware.

  The library provides methods for passing a location instead of a calendar
  system, and selecting an appropriate calendar based on which was most common
  in that location at that point in time. _(Some guess work is involved in this
  process when parsing dates, so it is still preferred to supply the calendar
  system, if known, when parsing.)_

- [x] Well-defined interfaces for extending the library.

  Add more calendar systems, type conversions, or geo-temporal relationships
  without forking/modifying the library itself.

## Installation

The steps here will vary based on which programming language(s) you're using.

For Golang, simply run `go get github.com/danhunsaker/calends`, and then place
`"github.com/danhunsaker/calends"` in the `import` wherever you intend use it.

Other languages will use Calends through a language-specific wrapper around the
compiled Golang lib. So for PHP, as an example, you'd install the `calends`
extension, probably via PECL.

## Usage

Calends exposes a very small handful of things for use outside the library
itself. One is the `Calends` class, which should be the only interface users of
the library ever need to touch. Another is the `TAI64NAXURTime` class, used to
store and manipulate the instants of time which make up a `Calends` instance.
The rest are interfaces for extending the library's functionality.

_**.TODO.**_

[GitHub]:https://github.com/danhunsaker/calends
[TAI64]:http://cr.yp.to/libtai/tai64.html
