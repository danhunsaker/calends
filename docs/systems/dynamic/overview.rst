.. _calendar-system-dynamic:

.. index:: Dynamic Calendars, calendars; Dynamic Calendars

Dynamic Calendars
=================

Unlike the other calendar systems supported by Calends, and documented here, the
dynamic calendar system isn't actually a calendar at all, at least on its own.
Instead, it's a way to provide a custom calendar system from information about
how it's constructed, rather than writing code to run calculations directly.
This approach isn't as fast or easily optimized as writing code, but it's more
easily accessible to those who don't have coding experience.

For example, an author writing a work of fiction with a calendar system unique
to their setting might wish to be able to track or convert dates and times in
their setting to equivalent dates and times in other systems in the same
setting, or even the author's own native calendar system. Using a dynamic
calendar, the author can add info about their setting's calendar system into
Calends, rather than needing to learn to code, write code as well as their work
of fiction, and/or hire a software developer to write their calendar system for
them. As of this writing, this feature is unique to Calends, among libraries
with similar functionalities and/or goals.

Because the dynamic calendars feature isn't actually a calendar system itself,
it doesn't have a calendar name of its own - instead, each dynamic calendar
system will provide its *own* name, and Calends will refer to that. Each
calendar system also defines its own formatting strings, and offset strings are
derived from the calendar system's units. More information on the structure and
creation of dynamic calendar systems is found below.

*Calendar Name:*
  *varies*

*Supported Input Types:*
  - string
  - integer
  - arbitrary-precision floating point number of seconds

*Supported Format Strings:*
  - *varies*

*Offsets:*
  - string with ``<value> <unit>`` combinations, where ``<unit>``\ s are defined
    by the calendar system itself

.. toctree::
   :maxdepth: 2

   calendar.rst
   unit.rst
   era.rst
   format.rst
   intercalation.rst
