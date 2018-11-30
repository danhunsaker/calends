.. _calendar-system-unix:

.. index:: UNIX Time, calendars; UNIX Time

UNIX Time
=========

Supports times that are seconds since ``CE 1970-01-01 00:00:00 UTC Gregorian``,
commonly used by computer systems for storing date/time values, internally.

*Calendar Name:*
  ``unix``

*Supported Input Types:*
  - string
  - integer
  - arbitrary-precision floating point number of seconds

*Supported Format Strings:*
  - values are always number of seconds since ``CE 1970-01-01 00:00:00 UTC
    Gregorian``

    - ``%d`` - integer string
    - ``%f`` - **floating point string**

*Offsets:*
  - number of seconds
