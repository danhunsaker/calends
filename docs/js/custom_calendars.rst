.. _custom-calendars-js:

.. js:module:: calends
   :noindex:

.. index::
   pair: custom calendars; JS
   pair: custom calendars; WASM

Custom Calendars in JS/WASM
===========================

Adding new calendars to Calends is a fairly straightforward process. Extend the
:js:class:`CalendarDefinition` abstract class, and implement three methods.
Then, simply construct an instance of your calendar system, and Calends will do
the rest.

Define
------

Extend the :js:class:`CalendarDefinition` class, implementing the following
methods:

.. js:class:: CalendarDefinition

   .. js:attribute:: name

      The name of the calendar system. Can be static or set in the
      :js:meth:`constructor`.

   .. js:attribute:: defaultFormat

      The default/fallback format for the calendar system. Can be static or set
      in the :js:meth:`constructor`.

   .. js:method:: constructor()

      This can do anything you like.

   .. js:method:: toInternal(stamp, format)

      :param stamp: The input. Should support ``string``\ s at the very minimum.
      :type stamp: ``any``
      :param format: The format string for parsing the input date.
      :type format: ``string``
      :return: The parsed internal timestamp.
      :rtype: :js:class:`TAI64Time`
      :throws CalendsException: when an error occurs

      Converts an input date/time representation to an internal
      :js:class:`TAI64Time`.

   .. js:method:: fromInternal(instant, format)

      :param instant: The internal timestamp value.
      :type instant: :js:class:`TAI64Time`
      :param format: The format string for formatting the output date.
      :type format: ``string``
      :return: The formatted date/time.
      :rtype: ``string``
      :throws CalendsException: when an error occurs

      Converts an internal :js:class:`TAI64Time` to a date/time string.

   .. js:method:: offset(instant, offset)

      :param instant: The internal timestamp value.
      :type instant: :js:class:`TAI64Time`
      :param offset: The input offset. Should support ``string``\ s at the very
                     minimum.
      :type offset: ``any``
      :return: The adjusted internal timestamp.
      :rtype: :js:class:`TAI64Time`
      :throws CalendsException: when an error occurs

      Adds the given offset to an internal :js:class:`TAI64Time`.

Registration
------------

Register
::::::::

Once it is registered with the library, your calendar system can be used from
anywhere in your application. To register a system, simply call
:js:meth:`register` on an object of your new class:

.. js:method:: CalendarDefinition.register()

   Registers a calendar system instance with the internal Calends library.

Unregister
::::::::::

The way to unregister a calendar system is to do so manually, using the instance
you created to register it with in the first place:

.. js:method:: CalendarDefinition.unregister()

   Removes a calendar system from the callback list.

Check and List
::::::::::::::

.. js:method:: CalendarDefinition.isRegistered()

   :return: Whether or not the calendar system is currently registered.
   :rtype: ``bool``

   Returns whether or not a calendar system has been registered, yet.

.. js:method:: CalendarDefinition.registered()

   :return: The sorted list of calendar systems currently registered.
   :rtype: ``[string]``

   Returns the list of calendar systems currently registered.

Types and Values
----------------

Now we get to the inner workings that make calendar systems function – even the
built-in ones. The majority of the "magic" comes from the :js:class:`TAI64Time`
object itself, as a reliable way of storing the exact instants being calculated,
and the only way times are handled by the library itself. A handful of methods
provide basic operations that calendar system developers can use to simplify
their conversions (adding and subtracting the values of other timestamps, and
importing/exporting timestamp values from/to string and numeric types, in
particular), and a couple of helpers exclusively handle adding and removing UTC
leap second offsets. As long as you can convert your dates to/from Unix
timestamps in a string or numeric type, the rest is handled entirely by these
helpers in the library itself.

.. js:class:: TAI64Time

   :js:class:`TAI64Time` stores a ``TAI64NARUX`` instant in a reliable,
   easily-converted format. Each 9-digit fractional segment is stored in a
   separate 32-bit integer to preserve its value with a very high degree of
   accuracy, without having to rely on string parsing or external
   arbitrary-precision mathematics libraries.

   .. js:attribute:: seconds

      The number of TAI seconds since ``CE 1970-01-01 00:00:00 TAI``. Should be
      an integer value.

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
         calculated, be sure to convert it using :js:meth:`fromUTC` before
         returning the result to the rest of the library. And then, of course,
         you'll also need to convert instants from the library back using
         :js:meth:`toUTC` before generating outputs.

   .. js:attribute:: nano

      The first 9 digits of the timestamp's fractional component.

   .. js:attribute:: atto

      The 10th through 18th digits of the fractional component.

   .. js:attribute:: ronto

      The 19th through 27th digits of the fractional component.

   .. js:attribute:: udecto

      The 28th through 36th digits of the fractional component.

   .. js:attribute:: xindecto

      The 37th through 45th digits of the fractional component.

   .. js:method:: add(z)

      :param z: The timestamp to add to the current one.
      :type z: :js:class:`TAI64Time`
      :return: The sum of the two timestamps.
      :rtype: :js:class:`TAI64Time`

      Calculates the sum of two :js:class:`TAI64Time` values.

   .. js:method:: sub(z)

      :param z: The timestamp to subtract from the current one.
      :type z: :js:class:`TAI64Time`
      :return: The difference of the two timestamps.
      :rtype: :js:class:`TAI64Time`

      Calculates the difference of two :js:class:`TAI64Time` values.

   .. js:method:: toString()

      :return: The decimal string representation of the current timestamp.
      :rtype: ``string``

      Returns the decimal string representation of the :js:class:`TAI64Time`
      value.

   .. js:method:: fromString(in)

      :param in: The decimal string representation of a timestamp to calculate.
      :type in: ``string``
      :return: The calculated timestamp.
      :rtype: :js:class:`TAI64Time`

      Calculates a :js:class:`TAI64Time` from its decimal string representation.

   .. js:method:: toHex()

      :return: The hexadecimal string representation of the current timestamp.
      :rtype: ``string``

      Returns the hexadecimal string representation of the :js:class:`TAI64Time`
      value.

   .. js:method:: fromHex(in)

      :param in: The hexadecimal string representation of a timestamp to
                 calculate.
      :type in: ``string``
      :return: The calculated timestamp.
      :rtype: :js:class:`TAI64Time`

      Calculates a :js:class:`TAI64Time` from its hexadecimal string
      representation.

   .. js:method:: toNumber()

      :return: The numeric representation of the current timestamp.
      :rtype: ``number``

      Returns the ``number`` representation of the :js:class:`TAI64Time` value.

   .. js:method:: fromNumber(in)

      :param in: The arbitrary-precision floating point representation of a
                 timestamp to calculate.
      :type in: ``number``
      :return: The calculated timestamp.
      :rtype: :js:class:`TAI64Time`

      Calculates a :js:class:`TAI64Time` from its numeric representation.

   .. js:method:: fromUTC()

      :return: The calculated timestamp.
      :rtype: :js:class:`TAI64Time`

      Removes the UTC leap second offset from a TAI64Time value.

   .. js:method:: toUTC()

      :return: The calculated timestamp.
      :rtype: :js:class:`TAI64Time`

      Adds the UTC leap second offset to a TAI64Time value.
