# Calends

[![Software License](https://img.shields.io/github/license/danhunsaker/calends.svg?style=flat-square)](LICENSE)
[![Gitter](https://img.shields.io/gitter/room/danhunsaker/calends.svg?style=flat-square)](https://gitter.im/danhunsaker/calends)
[![Go Report Card](https://goreportcard.com/badge/github.com/danhunsaker/calends)](https://goreportcard.com/report/github.com/danhunsaker/calends)
[![GoDoc Reference](https://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danhunsaker/calends)

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

    - [x] Automatic calculation of leap second offsets

    - [ ] Estimation of undefined past and future leap second insertions

    - [ ] Automatic updates for handling leap second insertions

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

- [ ] Time zone support.

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

`Calends` objects are immutable - all methods return a new `Calends` object
where they might otherwise alter the current one. This is true even of the
`Set*()` methods. This has the side effect of using more memory to perform
manipulations than updating values on an existing object would. It makes many
operations safer, though, than mutable objects would allow.

Language-specific documentation is available, and may give a more concrete idea
of how to use Calends in a given language/environment, but the general usage
information given here should be valid for all of them.

### Create

- `calends.Create(value, calendar, format)`

  Creates a new `Calends` object, using `calendar` to select a calendar system,
  and `format` to parse the contents of `value` into the `Calends` object's
  internal instants. The contents of `value` can vary based on the calendar
  system itself, but generally speaking can always be a string. In any case, the
  value can always be a string→value map (associative array, hash, or whatever
  your language of choice prefers to call it), where the keys are any two of
  `start`, `end`, and `duration`. If all three are provided, `duration` is
  ignored in favor of calculating it directly. If only one is provided, `value`
  is passed to the calendar system itself unchanged. The calendar system then
  converts `value` to a `TAI64NAXURTime` instant, which the `Calends`
  object sets to the appropriate internal value.

### Read

- `c.Date(calendar, format)` / `c.EndDate(calendar, format)`

  Retrieves the start or end date of the `Calends` object `c` as a string. The
  value is generated by the calendar system given in  `calendar`, according to
  the format string in `format`.

- `c.Duration()`

  Retrieves the duration of the `Calends` object `c` as an arbitrary-precision
  floating point number. This value will be `0` if the `Calends` object is an
  instant.

### Update

- `c.SetDate(value, calendar, format)` / `c.SetEndDate(value, calendar, format)`

  Sets the start or end date of a `Calends` object, based on the `Calends`
  object `c`. The inputs are the same as for `Create()`, above, except the
  string→value map option isn't available, since you're already specifically
  setting the start or end value explicitly, depending on which method you call.

- `c.SetDuration(duration, calendar)` / `c.SetDurationFromEnd(duration, calendar)`

  Sets the duration of a `Calends` object, adjusting the end or start point
  accordingly, based on the `Calends` object `c`. The `duration` value is
  interpreted by the calendar system given in `calendar`, so is subject to any
  of its rules. `SetDurationFromEnd()` will adjust the start point, using the
  end as the anchor for the duration.

### Manipulate

- `c.Add(offset, calendar)` / `c.AddFromEnd(offset, calendar)`

  Increases the corresponding date in the `Calends` object `c` by `offset`, as
  interpreted by the calendar system given in `calendar`.

- `c.Subtract(offset, calendar)` / `c.SubtractFromEnd(offset, calendar)`

  Works the same as `Add()` / `AddFromEnd()`, except it decreases the
  corresponding date, rather than increasing it.

- `c.Next(offset, calendar)` / `c.Previous(offset, calendar)`

  Returns a `Calends` object of `offset` duration (as interpreted by the
  calendar system given in `calendar`), which abuts the `Calends` object `c`. If
  `offset` is empty, `calendar` is ignored, and the duration of `c` is used
  instead.

### Combine

- `c1.Merge(c2)`

  Returns a `Calends` object spanning from the earliest start date to the latest
  end date between `Calends` objects `c1` and `c2`.

- `c1.Intersect(c2)`

  Returns a `Calends` object spanning the overlap between `Calends` objects `c1`
  and `c2`. If `c1` and `c2` don't overlap, returns an error.

- `c1.Gap(c2)`

  Returns a `Calends` object spanning the gap between `Calends` objects `c1` and
  `c2`. If `c1` and `c2` overlap (and there is, therefore, no gap to return),
  returns an error.

### Compare

- `c1.Difference(c2, mode)`

  Returns the difference of `Calends` object `c1` minus `c2`, using `mode` to
  select which values to use in the calculation. Valid `mode`s include:

  - `start` - use the start date of both `c1` and `c2`
  - `duration` - use the duration of both `c1` and `c2`
  - `end` - use the end date of both `c1` and `c2`
  - `start-end` - use the start of `c1`, and the end of `c2`
  - `end-start` - use the end of `c1`, and the start of `c2`

- `c1.Compare(c2, mode)`

  Returns `-1` if `Calends` object `c1` is less than `Calends` object `c2`, `0`
  if they are equal, and `1` if `c1` is greater than `c2`, using `mode` to
  select which values to use in the comparison. Valid modes are the same as for
  `Difference()`, above.

- `c1.Contains(c2)`

  Returns a boolean value indicating whether `Calends` object `c1` contains all
  of `Calends` object `c2`.

- `c1.Overlaps(c2)`

  Returns a boolean value indicating whether `Calends` object `c1` overlaps with
  `Calends` object `c2`.

- `c1.Abuts(c2)`

  Returns a boolean value indicating whether `Calends` object `c1` abuts
  `Calends` object `c2` (that is, whether one begins at the same instant the
  other ends).

- `c1.IsSame(c2)`

  Returns a boolean value indicating whether `Calends` object `c1` covers the
  same span of time as `Calends` object `c2`.


- `c1.IsShorter(c2)` / `c1.IsSameDuration(c2)` / `c1.IsLonger(c2)`

  Returns a boolean comparing the duration of `Calends` objects `c1` and `c2`.

- `c1.IsBefore(c2)` / `c1.StartsBefore(c2)` / `c1.EndsBefore(c2)`

  Returns a boolean comparing `Calends` object `c1` with the start date of
  `Calends` object `c2`. `IsBefore` compares the entirety of `c1` with `c2`;
  `StartsBefore` compares only the start date of `c1`; `EndsBefore` compares
  only the end date of `c1`.

- `c1.IsDuring(c2)` / `c1.StartsDuring(c2)` / `c1.EndsDuring(c2)`

  Returns a boolean indicating whether `Calends` object `c1` lies between the
  start and end dates of `Calends` object `c2`. `IsDuring` compares the entirety
  of `c1` with `c2`; `StartsDuring` compares only the start date of `c1`;
  `EndsDuring` compares only the end date of `c1`.

- `c1.IsAfter(c2)` / `c1.StartsAfter(c2)` / `c1.EndsAfter(c2)`

  Returns a boolean comparing `Calends` object `c1` with the end date of
  `Calends` object `c2`. `IsAfter` compares the entirety of `c1` with `c2`;
  `StartsAfter` compares only the start date of `c1`; `EndsAfter` compares only
  the end date of `c1`.

### Calendar Systems

Currently supported calendar systems, and the options available for each, are
listed below. Formats in **bold** are the default format for that calendar.

- `tai64`
  - Formats
    - Calends `TAI64NAXURTime` object _(input only)_
    - string with TAI instant representation in one of the following layouts:
      - `decimal` - number of seconds since 1970.01.01 00:00:00 TAI
      - `tai64` - TAI64 External Representation; the hexadecimal version of
        `decimal` plus 2<sup>62</sup>, with no fractional seconds (16 hexits
        total)
      - `tai64n` - TAI64N External Representation; `tai64` with 9 decimal places
        encoded as 8 additional hexadecimal digits (24 hexits total)
      - `tai64na` - TAI64NA External Representation; `tai64n` with 9 more
        decimal places (18 total) encoded as 8 additional hexadecimal digits (32
        hexits total)
      - `tai64nax` - TAI64NAX External Representation; `tai64na` with 9 more
        decimal places (27 total) encoded as 8 additional hexadecimal digits (40
        hexits total)
      - `tai64naxu` - TAI64NAXU External Representation; `tai64nax` with 9 more
        decimal places (36 total) encoded as 8 additional hexadecimal digits (48
        hexits total)
      - **`tai64naxur` - TAI64NAXUR External Representation; `tai64naxu` with 9
        more decimal places (45 total) encoded as 8 additional hexadecimal
        digits (56 hexits total)**
  - Offsets
    - Calends `TAI64NAXURTime` object
    - arbitrary-precision floating point number of seconds
    - string with `decimal` format layout (above)

- `unix`
  - Formats
    - **number of seconds since 1970-01-01 00:00:00 UTC**
    - input can be integer or float, in either numeric or string representation
    - output uses Golang `fmt.Print()` conventions
  - Offsets
    - number of seconds
      - same input formatting as above

- `gregorian`
  - Formats
    - Golang `time.Time` object _(input only)_
    - Golang `time` package format strings (**RFC1123 layout**)
    - C-style `strptime()`/`strftime()` format strings
  - Offsets
    - Golang `time.Duration` object
    - string with Gregorian time units
      - must be relative times
      - use full words instead of abbreviations for time units (such as
        `seconds` instead of just `s`)

## Contributing

Pull requests are always welcome! That said, please be open to discussing the PR
content, and possibly revising it if requested. Not all requests can be merged,
and not all changes are desired.

## Security Reporting

Report all security-related issues to [dan (dot) hunsaker (plus) calends (at)
gmail](mailto:dan.hunsaker+calends@gmail.com), and use PGP or GPG protections on
your message. Security issues will be addressed internally before making any
vulnerability announcements.

[GitHub]:https://github.com/danhunsaker/calends
[TAI64]:http://cr.yp.to/libtai/tai64.html
