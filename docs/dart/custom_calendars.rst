.. _custom-calendars-dart:

.. dart:package:: calends
   :noindex:

.. index::
   pair: custom calendars; Dart

Custom Calendars in Dart
========================

Adding new calendars to Calends is a fairly straightforward process. Extend the
:dart:class:`CalendarDefinition` abstract class, and implement two getters and
three methods. Then, simply construct an instance of your calendar system, and
Calends will do the rest.

Define
------

Extend the :dart:class:`CalendarDefinition` class, implementing the following
methods:

.. dart:class:: CalendarDefinition

   .. dart:method:: get name

      :return: The name of the calendar system.
      :rtype: ``String``

   .. dart:method:: get defaultFormat

      :return: The defalt format of the calendar system.
      :rtype: ``String``

   .. dart:method:: toInternal(dynamic stamp, String format)

      :param stamp: The input stamp. Should support strings at the very minimum.
      :type stamp: ``dynamic``
      :param format: The format string for parsing the input stamp.
      :type format: ``String``
      :return: The parsed internal timestamp.
      :rtype: :dart:class:`TAI64Time`
      :throws CalendsException: when an error occurs

      Converts an input date/time representation to an internal
      :dart:class:`TAI64Time`.

   .. dart:method:: fromInternal(TAI64Time stamp, String format)

      :param stamp: The internal timestamp value.
      :type stamp: :dart:class:`TAI64Time`
      :param format: The format string for formatting the output date.
      :type format: ``String``
      :return: The formatted date/time.
      :rtype: ``String``
      :throws CalendsException: when an error occurs

      Converts an internal :dart:class:`TAI64Time` to a date/time string.

   .. dart:method:: offset(TAI64Time stamp, dynamic offset)

      :param stamp: The internal timestamp value.
      :type stamp: :dart:class:`TAI64Time`
      :param offset: The input offset. Should support strings at the very
                     minimum.
      :type offset: ``dynamic``
      :return: The adjusted internal timestamp.
      :rtype: :dart:class:`TAI64Time`
      :throws CalendsException: when an error occurs

      Adds the given offset to an internal :dart:class:`TAI64Time`.

Registration
------------

Register
::::::::

Once it is registered with the library, your calendar system can be used from
anywhere in your application:

.. dart:method:: CalendarDefinition.register()

   Adds a calendar system to the callback list.

Unregister
::::::::::

When you are done with a calendar system, it is best practice to free up
resources by unregistering it:

.. dart:method:: CalendarDefinition.unregister()

   Removes a calendar system from the callback list.

Check and List
::::::::::::::

.. dart:method:: CalendarDefinition.isRegistered()

   :return: Whether or not the calendar system is currently registered.
   :rtype: ``bool``

   Returns whether or not a calendar system has been registered, yet.

.. dart:method:: CalendarDefinition.listRegistered()

   :return: The sorted list of calendar systems currently registered.
   :rtype: ``List<String>``

   Returns the list of calendar systems currently registered.

Types and Values
----------------

Now we get to the inner workings that make calendar systems function – even the
built-in ones. The majority of the "magic" comes from the
:dart:class:`TAI64Time` object itself, as a reliable way of storing the exact
instants being calculated, and the only way times are handled by the library
itself. A handful of methods provide basic operations that calendar system
developers can use to simplify their conversions (adding and subtracting the
values of other timestamps, and importing/exporting timestamp values from/to
string and numeric types, in particular), and a couple of helpers exclusively
handle adding and removing UTC leap second offsets. As long as you can convert
your dates to/from Unix timestamps in a string or numeric type, the rest is
handled entirely by these helpers in the library itself.

.. dart:class:: TAI64Time

   :dart:class:`TAI64Time` stores a ``TAI64NARUX`` instant in a reliable,
   easily-converted format. Each 9-digit fractional segment is stored in a
   separate 32-bit integer to preserve its value with a very high degree of
   accuracy, without having to rely on string parsing or external
   arbitrary-precision mathematics libraries.

   .. dart:attribute:: Seconds (int)

      The number of TAI seconds since ``CE 1970-01-01 00:00:00 TAI``.

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
         calculated, be sure to convert it using :dart:meth:`utcToTai` before
         returning the result to the rest of the library. And then, of course,
         you'll also need to convert instants from the library back using
         :dart:meth:`taiToUtc` before generating outputs.

   .. dart:attribute:: Nano (int)

      The first 9 digits of the timestamp's fractional component.

   .. dart:attribute:: Atto (int)

      The 10th through 18th digits of the fractional component.

   .. dart:attribute:: Ronto (int)

      The 19th through 27th digits of the fractional component.

   .. dart:attribute:: Udecto (int)

      The 28th through 36th digits of the fractional component.

   .. dart:attribute:: Xindecto (int)

      The 37th through 45th digits of the fractional component.

   .. dart:method:: add(TAI64Time z)

      :param z: The timestamp to add to the current one.
      :type z: :dart:class:`TAI64Time`
      :return: The sum of the two timestamps.
      :rtype: :dart:class:`TAI64Time`

      Calculates the sum of two :dart:class:`TAI64Time` values.

   .. dart:method:: sub(TAI64Time z)

      :param z: The timestamp to subtract from the current one.
      :type z: :dart:class:`TAI64Time`
      :return: The difference of the two timestamps.
      :rtype: :dart:class:`TAI64Time`

      Calculates the difference of two :dart:class:`TAI64Time` values.

   .. dart:method:: toTAI64String()

      :return: The decimal string representation of the current timestamp.
      :rtype: ``String``

      Returns the decimal string representation of the :dart:class:`TAI64Time`
      value.

   .. dart:method:: fromTAI64String(String in)

      :param in: The decimal string representation of a timestamp to calculate.
      :type in: string
      :return: The calculated timestamp.
      :rtype: :dart:class:`TAI64Time`

      Calculates a :dart:class:`TAI64Time` from its decimal string
      representation.

   .. dart:method:: toHex()

      :return: The hexadecimal string representation of the current timestamp.
      :rtype: ``String``

      Returns the hexadecimal string representation of the
      :dart:class:`TAI64Time` value.

   .. dart:method:: fromHex(string in)

      :param in:
         The hexadecimal string representation of a timestamp to calculate.
      :type in: string
      :return: The calculated timestamp.
      :rtype: :dart:class:`TAI64Time`

      Calculates a :dart:class:`TAI64Time` from its hexadecimal string
      representation.

   .. dart:method:: toDouble()

      :return: The floating point representation of the current timestamp.
      :rtype: ``double``

      Returns the ``double`` representation of the :dart:class:`TAI64Time`
      value.

   .. dart:method:: fromDouble(double in)

      :param in:
         The floating point representation of a timestamp to calculate.
      :type in: ``double``
      :return: The calculated timestamp.
      :rtype: :dart:class:`TAI64Time`

      Calculates a :dart:class:`TAI64Time` from its ``double`` representation.

   .. dart:method:: utcToTai()

      :return: The calculated timestamp.
      :rtype: :dart:class:`TAI64Time`

      Removes the UTC leap second offset from a :dart:class:`TAI64Time` value.
      Used when converting from Unix time to TAI time.

   .. dart:method:: taiToUtc()

      :return: The calculated timestamp.
      :rtype: :dart:class:`TAI64Time`

      Adds the UTC leap second offset to a :dart:class:`TAI64Time` value. Used
      when converting from TAI time to Unix time.
