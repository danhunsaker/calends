.. _usage-cli:

.. index::
   pair: usage; CLI

Using Calends from the Command Line
===================================

Calends can be used from the command line directly, though some if its features
are limited or unavailable. Specifically, it doesn't support custom calendars,
so you'll need to ensure you build it with the calendar you want already loaded.

Command Line Options
--------------------

.. program:: calends

The available options for ``calends``, on the command line directly, are the
following:

.. option:: convert <from-calendar> <from-format> <to-calendar> <to-format> [<date>]

   :- from-calendar: The calendar system to parse the date/time with.
   :- from-calendar: The format the date/time is expected to use.
   :- to-calendar: The calendar system to format the date/time with.
   :- to-calendar: The format the date/time is expected to use.
   :- date: The value to convert.

   Converts a date from one calendar/format to another. If ``date`` isn't
   provided in the arguments, it's read from ``/dev/stdin`` instead.

.. option:: parse <from-calendar> <from-format> [<date>]

   :- from-calendar: The calendar system to parse the date/time with.
   :- from-calendar: The format the date/time is expected to use.
   :- date: The value to parse.

   Converts a date from the given calendar/format to a portable/unambiguous date
   stamp. The output from this command can then be used as input to others.

.. option:: format <to-calendar> <to-format> [<stamp>]

   :- to-calendar: The calendar system to format the date/time with.
   :- to-calendar: The format the date/time is expected to use.
   :- stamp: The value to format.

   Converts a date stamp from the :option:`parse` command to the given
   caledar/format.

.. option:: offset <offset-calendar> [<offset> [<stamp>]]

   :- offset-calendar: The calendar system to interpret the offset with.
   :- offset: The offset to add.
   :- stamp: The value to add the offset to.

   Adds an offset to the date stamp from the :option:`parse` command.

.. program:: calends compare

There is also a ``calends compare``, whose options are these:

.. option:: contains [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` contains
   ``stamp2``.

.. option:: overlaps [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` overlaps
   with ``stamp2``.

.. option:: abuts [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` abuts
   ``stamp2``.

.. option:: same [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` is the same
   as ``stamp2``.

.. option:: shorter [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` is shorter
   than ``stamp2``.

.. option:: same-duration [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` is the same
   duration as ``stamp2``.

.. option:: longer [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` is longer
   than ``stamp2``.

.. option:: before [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` is before
   ``stamp2``.

.. option:: start-before [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` starts
   before ``stamp2``.

.. option:: end-before [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` ends before
   ``stamp2``.

.. option:: during [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` is during
   ``stamp2``.

.. option:: start-during [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` starts
   during ``stamp2``.

.. option:: end-during [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` ends during
   ``stamp2``.

.. option:: after [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` is after
   ``stamp2``.

.. option:: start-after [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` starts
   after ``stamp2``.

.. option:: end-after [<stamp1> [<stamp2>]]

   :- stamp1: The value to compare.
   :- stamp2: The value to compare the other to.

   Compares ``stamp1`` to ``stamp2``, and returns whether ``stamp1`` ends after
   ``stamp2``.

Interactive/Batch Mode
----------------------

Create
++++++

.. program:: calends (batch-mode)

.. option:: parse <calendar> <format> <date> <target>

   :- calendar: The calendar system to parse the date/time with.
   :- format: The format the date/time is expected to use.
   :- date: The value to parse.
   :- target: The name to give the result.

   Creates a new ``Calends`` value, using ``calendar`` to select a calendar
   system, and ``format`` to describe the contents of ``date`` to parse. The
   result is stored as ``target``, so it can be used later on by other commands.

.. option:: parse-range <calendar> <format> <date> <end-date> <target>

   :- calendar: The calendar system to parse the dates/times with.
   :- format: The format the dates/times are expected to use.
   :- date: The start date to parse.
   :- end-date: The end date to parse.
   :- target: The name to give the result.

   Creates a new ``Calends`` value, using ``calendar`` to select a calendar
   system, and ``format`` to describe the contents of ``date`` and ``end-date``
   to parse. The result is stored as ``target``, so it can be used later on by
   other commands.

Read
++++

.. option:: date <calendar> <format> <source>

   :- calendar: The calendar system to format the date/time with.
   :- format: The format the date/time is expected to be in.
   :- source: The name of the ``Calends`` value to use.

   Retrieves the start date of ``source`` as a string. The value is generated by
   the calendar system given in ``calendar``, according to the format string in
   ``format``.

.. option:: end-date <calendar> <format> <source>

   :- calendar: The calendar system to format the date/time with.
   :- format: The format the date/time is expected to be in.
   :- source: The name of the ``Calends`` value to use.

   Retrieves the end date of ``source`` as a string. The value is generated by
   the calendar system given in ``calendar``, according to the format string in
   ``format``.

Update
++++++

.. option:: set-date <calendar> <format> <date> <source> <target>

   :- calendar: The calendar system to parse the date/time with.
   :- format: The format the date/time is expected to use.
   :- date: The value to parse the date/time from.
   :- source: The name of the ``Calends`` value to use.
   :- target: The name to give the result.

   Sets the start date of ``target`` based on ``source``'s current value. The
   inputs are the same as for :option:`parse`, above.

.. option:: set-end-date <calendar> <format> <date> <source> <target>

   :- calendar: The calendar system to parse the date/time with.
   :- format: The format the date/time is expected to use.
   :- date: The value to parse the date/time from.
   :- source: The name of the ``Calends`` value to use.
   :- target: The name to give the result.

   Sets the end date of ``target`` based on ``source``'s current value. The
   inputs are the same as for :option:`parse`, above.

Manipulate
++++++++++

.. option:: add <calendar> <offset> <source> <target>

   :- calendar: The calendar system to parse the offset with.
   :- offset: The value to parse the offset from.
   :- source: The name of the ``Calends`` value to use.
   :- target: The name to give the result.

   Increases the end date of ``source``'s current value by ``offset``, as
   interpreted by the calendar system given in ``calendar``.

.. option:: add-from-end <calendar> <offset> <source> <target>

   :- calendar: The calendar system to parse the offset with.
   :- offset: The value to parse the offset from.
   :- source: The name of the ``Calends`` value to use.
   :- target: The name to give the result.

   Increases the start date of ``source``'s current value by ``offset``, as
   interpreted by the calendar system given in ``calendar``.

.. option:: subtract <calendar> <offset> <source> <target>

   :- calendar: The calendar system to parse the offset with.
   :- offset: The value to parse the offset from.
   :- source: The name of the ``Calends`` value to use.
   :- target: The name to give the result.

   Works the same as :option:`add`, except it decreases the start date, rather
   than increasing it.

.. option:: subtract-from-end <calendar> <offset> <source> <target>

   :- calendar: The calendar system to parse the offset with.
   :- offset: The value to parse the offset from.
   :- source: The name of the ``Calends`` value to use.
   :- target: The name to give the result.

   Works the same as :option:`add-from-end`, except it decreases the end date,
   rather than increasing it.

.. option:: next <calendar> <offset> <source> <target>

   :- calendar: The calendar system to parse the offset with.
   :- offset: The value to parse the offset from.
   :- source: The name of the ``Calends`` value to use.
   :- target: The name to give the result.

   Sets ``target`` to a ``Calends`` value of ``offset`` duration (as interpreted
   by the calendar system given in ``calendar``), which abuts the end of
   ``source``. If ``offset`` is empty, ``calendar`` is ignored, and ``source``'s
   duration is used instead.

.. option:: previous <calendar> <offset> <source> <target>

   :- calendar: The calendar system to parse the offset with.
   :- offset: The value to parse the offset from.
   :- source: The name of the ``Calends`` value to use.
   :- target: The name to give the result.

   Sets ``target`` to a ``Calends`` value of ``offset`` duration (as interpreted
   by the calendar system given in ``calendar``), which abuts the start of
   ``source``. If ``offset`` is empty, ``calendar`` is ignored, and ``source``'s
   duration is used instead.

Combine
+++++++

.. option:: merge <source> <combine> <target>

   :- source: The name of the ``Calends`` value to use.
   :- combine: The ``Calends`` value to merge.
   :- target: The name to give the result.

   Sets ``target`` to a value spanning from the earliest start date to the
   latest end date between ``source`` and ``combine``.

.. option:: intersect <source> <combine> <target>

   :- source: The name of the ``Calends`` value to use.
   :- combine: The ``Calends`` value to intersect.
   :- target: The name to give the result.

   Sets ``target`` to a value spanning the overlap between ``source`` and
   ``combine``. If ``source`` and ``combine`` don't overlap, returns an error.

.. option:: gap <source> <combine> <target>

   :- source: The name of the ``Calends`` value to use.
   :- combine: The ``Calends`` value to gap.
   :- target: The name to give the result.

   Sets ``target`` to a value spanning the gap between ``source`` and
   ``combine``. If ``source`` and ``combine`` overlap (and there is, therefore,
   no gap to return), returns an error.

Compare
+++++++

.. option:: difference <source> <compare> <mode>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.
   :- mode: The comparison mode.

   Returns the difference of ``source`` minus ``compare``, using ``mode`` to
   select which values to use in the calculation. Valid ``mode``\ s include:

   - ``start`` - use the start date of both ``source`` and ``compare``
   - ``duration`` - use the duration of both ``source`` and ``compare``
   - ``end`` - use the end date of both ``source`` and ``compare``
   - ``start-end`` - use the start of ``source``, and the end of ``compare``
   - ``end-start`` - use the end of ``source``, and the start of ``compare``

.. option:: compare <source> <compare> <mode>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.
   :- mode: The comparison mode.

   Returns ``-1`` if ``source`` is less than ``compare``, ``0`` if they are
   equal, and ``1`` if ``source`` is greater than ``compare``, using ``mode`` to
   select which values to use in the comparison. Valid ``mode``\ s are the same
   as for :option:`difference`, above.

.. option:: contains <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Checks whether ``source`` contains all of ``compare``.

.. option:: overlaps <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Checks whether ``source`` overlaps with ``compare``.

.. option:: abuts <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Checks whether ``source`` abuts ``compare`` (that is, whether one begins at
   the same instant the other ends).

.. option:: is-same <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Checks whether ``source`` covers the same span of time as ``compare``.

.. option:: is-shorter <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the duration of ``source``  and ``compare``.

.. option:: is-same-duration <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the duration of ``source``  and ``compare``.

.. option:: is-longer <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the duration of ``source``  and ``compare``.

.. option:: is-before <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the entirety of ``source`` with the start date of ``compare``.

.. option:: starts-before <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the start date of ``source`` with the start date of ``compare``.

.. option:: ends-before <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the end date of ``source`` with the start date of ``compare``.

.. option:: is-during <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Checks whether the entirety of ``source`` lies between the start and end
   dates of ``compare``.

.. option:: starts-during <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Checks whether the start date of ``source`` lies between the start and end
   dates of ``compare``.

.. option:: ends-during <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Checks whether the end date of ``source`` lies between the start and end
   dates of ``compare``.

.. option:: is-after <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the entirety of ``source`` with the end date of ``compare``.

.. option:: starts-after <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the start date of ``source`` with the end date of ``compare``.

.. option:: ends-after <source> <compare>

   :- source: The name of the ``Calends`` value to use.
   :- compare: The ``Calends`` value to compare.

   Compares the end date of ``source`` with the end date of ``compare``.
