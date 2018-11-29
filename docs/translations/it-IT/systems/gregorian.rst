.. _calendar-system-gregorian:

.. index:: Gregorian Calendar, calendars; Gregorian Calendar

The Gregorian Calendar
======================

Supports dates and times in the Gregorian calendar system, the current
international standard for communicating dates and times.

*Calendar Name:*
  ``gregorian``

*Supported Input Types:*
  - string

*Supported Format Strings:*
  - Golang ``time`` package format strings (**RFC1123 layout**)
  - C-style ``strptime()``/``strftime()`` format strings

*Offsets:*
  - string with Gregorian time units

    - must be relative times
    - use full words instead of abbreviations for time units (such as
      ``seconds`` instead of just ``s``)
