.. _custom-calendars-php:

.. php:namespace:: Calends

.. index::
   pair: custom calendars; PHP

Custom Calendars in PHP
==========================

Adding new calendars to Calends is a fairly straightforward process. Extend the
:php:class:`CalendarDefinition` abstract class, and implement three methods.
Then, simply construct an instance of your calendar system, and Calends will do
the rest.

Define
------

Extend the :php:class:`CalendarDefinition` class, implementing the following
methods:

.. php:class:: CalendarDefinition

   .. php:method:: toInternal(mixed $date, string $format): TAITime

      :param $date: The input date. Should support strings at the very minimum.
      :type $date: ``mixed``
      :param $format: The format string for parsing the input date.
      :type $format: ``string``
      :return: The parsed internal timestamp.
      :rtype: :php:class:`TAITime`
      :throws CalendsException: when an error occurs

      Converts an input date/time representation to an internal
      :php:class:`TAITime`.

   .. php:method:: fromInternal(TAITime $stamp, string $format): string

      :param $stamp: The internal timestamp value.
      :type $stamp: :php:class:`TAITime`
      :param $format: The format string for formatting the output date.
      :type $format: ``string``
      :return: The formatted date/time.
      :rtype: ``string``
      :throws CalendsException: when an error occurs

      Converts an internal :php:class:`TAITime` to a date/time string.

   .. php:method:: offset(TAITime $stamp, mixed $offset): TAITime

      :param $stamp: The internal timestamp value.
      :type $stamp: :php:class:`TAITime`
      :param $offset: The input offset. Should support strings at the very
                      minimum.
      :type $offset: ``mixed``
      :return: The adjusted internal timestamp.
      :rtype: :php:class:`TAITime`
      :throws CalendsException: when an error occurs

      Adds the given offset to an internal :php:class:`TAITime`.

Registration
------------

Register
::::::::

Once it is registered with the library, your calendar system can be used from
anywhere in your application. To register a system, simply construct an
instance:

.. code-block:: php

   $customCalendars[] = new MyCalendarSystem('example', 'yyyy-mm-dd@HH-MM-SS');

The first argument is the calendar system's name. The second is a default
format string which will be passed to your calendar system whenever users leave
the format string blank or unset with :php:class:`Calends` methods.

Unregister
::::::::::

There are two ways to unregister a calendar system once it's no longer needed.
The first is to simply destruct the instance you created to register it. For
that reason, it's important to store all your calendar systems to variables
rather than simply constructing them in place. Well, that and the fact you need
the calendar system object to persist in order to handle requests from the rest
of the library.

The other way to unregister a calendar system is to do so manually, using the
instance you created to register it in the first place:

.. php:method:: unregister()

   Removes a calendar system from the callback list.

Check and List
::::::::::::::

.. php:staticmethod:: isRegistered(string $name): bool

   :param $name: The calendar system name to check for.
   :type $name: ``string``
   :return: Whether or not the calendar system is currently registered.
   :rtype: ``bool``

   Returns whether or not a calendar system has been registered, yet.

.. php:staticmethod:: listRegistered(): array

   :return: The sorted list of calendar systems currently registered.
   :rtype: ``[string]``

   Returns the list of calendar systems currently registered.

Types and Values
----------------

Now we get to the inner workings that make calendar systems function – even the
built-in ones. The majority of the "magic" comes from the :php:class:`TAITime`
object itself, as a reliable way of storing the exact instants being calculated,
and the only way times are handled by the library itself. A handful of methods
provide basic operations that calendar system developers can use to simplify
their conversions (adding and subtracting the values of other timestamps, and
importing/exporting timestamp values from/to string and numeric types, in
particular), and a couple of helpers exclusively handle adding and removing UTC
leap second offsets. As long as you can convert your dates to/from Unix
timestamps in a string or numeric type, the rest is handled entirely by these
helpers in the library itself.

.. php:class:: TAITime

   :php:class:`TAITime` stores a ``TAI64NARUX`` instant in a reliable,
   easily-converted format. Each 9-digit fractional segment is stored in a
   separate 32-bit integer to preserve its value with a very high degree of
   accuracy, without having to rely on string parsing or external
   arbitrary-precision mathematics libraries.

   .. php:attr:: seconds (float)

      The number of TAI seconds since ``CE 1970-01-01 00:00:00 TAI``. Should be
      an integer value; the ``float`` type is used, here, only to be able to
      hold a full signed 64-bit integer value regardless of architecture.

      .. note:: TAI vs UTC

         You may have noticed that a TAI64Time object stores times in ``TAI
         seconds``, not ``Unix seconds``, with a timezone offset of ``TAI``
         rather than ``UTC``. This distinction is **very important** as it will
         affect internal calculations and comparisons to mix the two up. TAI
         time is very similar to Unix time (itself based on UTC time), with one
         major difference. While Unix/UTC seconds include the insertion and
         removal of "leap seconds" to keep the solar zenith at local noon (which
         is useful for day-to-day living and planning), TAI seconds are a
         continuous count, unconcerned with dates whatsoever. Indeed, the only
         reason a date was given in the description above was to make it easier
         for human readers to know exactly when ``0 TAI`` took place.

         In other words, once you have a Unix timestamp of your instant
         calculated, be sure to convert it using :php:meth:`fromUTC` before
         returning the result to the rest of the library. And then, of course,
         you'll also need to convert instants from the library back using
         :php:meth:`toUTC` before generating outputs.

   .. php:attr:: nano (integer)

      The first 9 digits of the timestamp's fractional component.

   .. php:attr:: atto (integer)

      The 10th through 18th digits of the fractional component.

   .. php:attr:: ronto (integer)

      The 19th through 27th digits of the fractional component.

   .. php:attr:: udecto (integer)

      The 28th through 36th digits of the fractional component.

   .. php:attr:: xindecto (integer)

      The 37th through 45th digits of the fractional component.

   .. php:method:: add(TAITime $z): TAITime

      :param $z: The timestamp to add to the current one.
      :type $z: :php:class:`TAITime`
      :return: The sum of the two timestamps.
      :rtype: :php:class:`TAITime`

      Calculates the sum of two :php:class:`TAITime` values.

   .. php:method:: sub(TAITime $z): TAITime

      :param $z: The timestamp to subtract from the current one.
      :type $z: :php:class:`TAITime`
      :return: The difference of the two timestamps.
      :rtype: :php:class:`TAITime`

      Calculates the difference of two :php:class:`TAITime` values.

   .. php:method:: toString(): string

      :return: The decimal string representation of the current timestamp.
      :rtype: ``string``

      Returns the decimal string representation of the :php:class:`TAITime`
      value.

      .. Note::

         :php:class:`TAITime` also implements :php:meth:`!__toString`, so you
         can use that instead of calling this function directly, if you prefer.

   .. php:method:: fromString(string $in): TAITime

      :param $in: The decimal string representation of a timestamp to calculate.
      :type $in: string
      :return: The calculated timestamp.
      :rtype: :php:class:`TAITime`

      Calculates a :php:class:`TAITime` from its decimal string representation.

   .. php:method:: toHex(): string

      :return: The hexadecimal string representation of the current timestamp.
      :rtype: ``string``

      Returns the hexadecimal string representation of the :php:class:`TAITime`
      value.

   .. php:method:: fromHex(string $in):TAITime

      :param $in: The hexadecimal string representation of a timestamp to calculate.
      :type $in: string
      :return: The calculated timestamp.
      :rtype: :php:class:`TAITime`

      Calculates a :php:class:`TAITime` from its hexadecimal string
      representation.

   .. php:method:: toNumber(): float

      :return: The numeric representation of the current timestamp.
      :rtype: ``float``

      Returns the ``float`` representation of the :php:class:`TAITime` value.

   .. php:method:: fromNumber(numeric $in): TAITime

      :param $in: The arbitrary-precision floating point representation of a
                 timestamp to calculate.
      :type $in: ``integer`` or ``float``
      :return: The calculated timestamp.
      :rtype: :php:class:`TAITime`

      Calculates a :php:class:`TAITime` from its numeric (``integer`` or
      ``float``) representation.

   .. php:method:: fromUTC(): TAITime

      :return: The calculated timestamp.
      :rtype: :php:class:`TAITime`

      Removes the UTC leap second offset from a TAITime value.

   .. php:method:: toUTC(): TAITime

      :return: The calculated timestamp.
      :rtype: :php:class:`TAITime`

      Adds the UTC leap second offset to a TAITime value.
