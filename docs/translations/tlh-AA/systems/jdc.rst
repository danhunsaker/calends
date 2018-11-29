.. _calendar-system-jdc:

.. index::
   single: Julian Day Count
   single: calendars; Julian Day Count
   see: JDC; Julian Day Count

Julian Day Count
================

A count of days since ``BCE 4713 Jan 01 12:00:00 UTC Julian (proleptic)``. Yes,
that's noon. This calendar system is used mostly for astronomy purposes, though
there is a modified variant with a narrower scope which counts from midnight
instead.

*Calendar Name:*
  ``jdc``

*Supported Input Types:*
  - string
  - integer
  - arbitrary-precision floating point number of seconds

*Supported Format Strings:*
  - ``full``     - the full, canonical Day Count
  - ``fullday``  - the full Day Count, without the fractional time part
  - ``fulltime`` - just the fractional time part of the full Day Count
  - ``modified`` - **an abbreviated Day Count, 2400000.5 less than the full
    (starts at midnight instead of noon)**
  - ``day``      - the modified Day Count, without the fractional time part
  - ``time``     - just the fractional time part of the modified Day Count

*Offsets:*
  - number of days, as integer or float, via numeric or string types

    - can include fractional days to indicate time
