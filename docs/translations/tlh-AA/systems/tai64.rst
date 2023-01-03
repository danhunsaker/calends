.. _calendar-system-tai64:

.. index:: TAI64 Time, calendars; TAI64 Time

TAI64 Time
==========

Supports times that are seconds since ``CE 1970-01-01 00:00:00 TAI Gregorian``
(plus |2^62|, when in hexadecimal), as defined at
https://cr.yp.to/libtai/tai64.html (though this library includes extensions to
the formats described there). These values are also used internally, so this
calendar system can be used to directly expose the underlying internal values in
a manner that allows them to be used elsewhere.

*Calendar Name:*
  ``tai64``

*Supported Input Types:*
  - string
  - ``TAITime``

*Supported Format Strings:*
  - ``decimal``    - decimal; full (45 decimal places) resolution; number of
    seconds since ``CE 1970-01-01 00:00:00 TAI Gregorian``
  - ``tai64``      - hexadecimal; just seconds; TAI64 External Representation
  - ``tai64n``     - hexadecimal; with nanoseconds; TAI64N External
    Representation
  - ``tai64na``    - hexadecimal; with attoseconds; TAI64NA External
    Representation
  - ``tai64nar``   - hexadecimal; with rontoseconds; TAI64NAR External
    Representation
  - ``tai64naru``  - hexadecimal; with udectoseconds; TAI64NARU External
    Representation
  - ``tai64narux`` - **hexadecimal; with xindectoseconds**; TAI64NARUX External
    Representation

*Offsets:*
  - ``TAITime`` object
  - arbitrary-precision floating point number of seconds
  - string with ``decimal`` format layout (above)

.. |2^62| replace:: 2\ :sup:`62`
