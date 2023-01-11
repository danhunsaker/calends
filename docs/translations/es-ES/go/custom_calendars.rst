.. _custom-calendars-go:

.. go:currentpackage:: calends/calendars

.. index::
   pair: custom calendars; Golang

Custom Calendars in Golang
==========================

Adding new calendars to Calends is a fairly straightforward process. Implement
one interface, or its three methods as standalone functions, and then simply
pass them to one of the two registration functions.

Define
------

The interface in question looks like this:

.. go:package:: calends/calendars

.. go:type:: CalendarDefinition

   .. go:function:: func (CalendarDefinition) ToInternal(date interface, format string) (TAI64NARUXTime, error)

      :param date: The input date. Should support :go:type:`string` at the very
                   minimum.
      :type date: :go:type:`interface{}`
      :param format: The format string for parsing the input date.
      :type format: :go:type:`string`
      :return: The parsed internal timestamp.
      :rtype: :go:type:`TAI64NARUXTime`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Converts an input date/time representation to an internal
      :go:type:`TAI64NARUXTime`.

   .. go:function:: func (CalendarDefinition) FromInternal(stamp TAI64NARUXTime, format string) (string, error)

      :param stamp: The internal timestamp value.
      :type stamp: :go:type:`TAI64NARUXTime`
      :param format: The format string for formatting the output date.
      :type format: :go:type:`string`
      :return: The formatted date/time.
      :rtype: :go:type:`string`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Converts an internal :go:type:`TAI64NARUXTime` to a date/time string.

   .. go:function:: func (CalendarDefinition) Offset(stamp TAI64NARUXTime, offset interface) (TAI64NARUXTime, error)

      :param stamp: The internal timestamp value.
      :type stamp: :go:type:`TAI64NARUXTime`
      :param offset: The input offset. Should support :go:type:`string` at the
                     very minimum.
      :type offset: :go:type:`interface{}`
      :return: The adjusted internal timestamp.
      :rtype: :go:type:`TAI64NARUXTime`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Adds the given offset to an internal :go:type:`TAI64NARUXTime`.

Registration
------------

Register
::::::::

Once it is registered with the library, your calendar system can be used from
anywhere in your application. To register a system, pass it to one of the
following two functions:

.. go:function:: func RegisterObject(name string, definition CalendarDefinition, defaultFormat string)

   :param name: The name to register the calendar system under.
   :type name: :go:type:`string`
   :param definition: The calendar system itself.
   :type definition: :go:type:`CalendarDefinition`
   :param defaultFormat: The default format string.
   :type defaultFormat: :go:type:`string`

   Registers a calendar system class, storing ``definition`` as ``name``, and
   saving ``defaultFormat`` for later use while parsing or formatting.

.. go:function:: func RegisterElements(name string, toInternal ToInternal, fromInternal FromInternal, offset Offset, defaultFormat string)

   :param name: The name to register the calendar system under.
   :type name: :go:type:`string`
   :param toInternal: The function for parsing dates into internal timestamps.
   :type toInternal: :go:func:`(CalendarDefinition) ToInternal`
   :param fromInternal: The function for formatting internal timestamps as
                        dates.
   :type fromInternal: :go:func:`(CalendarDefinition) FromInternal`
   :param offset: The function for adding an offset to internal timestamps.
   :type offset: :go:func:`(CalendarDefinition) Offset`
   :param defaultFormat: The default format string.
   :type defaultFormat: :go:type:`string`

   Registers a calendar system from its distinct functions. It does this by
   storing ``toInternal``, ``fromInternal``, and ``offset`` as the elements of
   ``name``, and saving ``defaultFormat`` for later use while parsing or
   formatting.

Unregister
::::::::::

.. go:function:: func Unregister(name string)

   :param name: The name of the calendar system to remove.
   :type name: :go:type:`string`

   Removes a calendar system from the callback list.

Check and List
::::::::::::::

.. go:function:: func Registered(calendar string) bool

   :param name: The calendar system name to check for.
   :type name: :go:type:`string`
   :return: Whether or not the calendar system is currently registered.
   :rtype: :go:type:`bool`

   Returns whether or not a calendar system has been registered, yet.

.. go:function:: func ListRegistered() []string

   :return: The sorted list of calendar systems currently registered.
   :rtype: :go:type:`[]string`

   Returns the list of calendar systems currently registered.

Types and Values
----------------

Now we get to the inner workings that make calendar systems function – even the
built-in ones. The majority of the "magic" comes from the
:go:type:`TAI64NARUXTime` object itself, as a reliable way of storing the exact
instants being calculated, and the only way times are handled by the library
itself. A handful of methods provide basic operations that calendar system
developers can use to simplify their conversions (adding and subtracting the
values of other timestamps, and importing/exporting timestamp values from/to
arbitrary-precision floating point :go:type:`math/big.Float`\ s, in particular),
and a couple of helpers exclusively handle adding and removing UTC leap second
offsets. As long as you can convert your dates to/from Unix timestamps in a
:go:type:`string` or :go:type:`math/big.Float`, the rest is handled entirely by
these helpers in the library itself.

.. go:type:: TAI64NARUXTime

   :param int64 Seconds: The number of TAI seconds since ``CE 1970-01-01
                         00:00:00 TAI``.
   :param uint32 Nano: The first 9 digits of the timestamp's fractional
                       component.
   :param uint32 Atto: The 10th through 18th digits of the fractional component.
   :param uint32 Ronto: The 19th through 27th digits of the fractional
                        component.
   :param uint32 Udecto: The 28th through 36th digits of the fractional component.
   :param uint32 Xindecto: The 37th through 45th digits of the fractional
                        component.

   :go:type:`TAI64NARUXTime` stores a ``TAI64NARUX`` instant in a reliable,
   easy-converted format. Each 9-digit fractional segment is stored in a
   separate 32-bit integer to preserve its value with a very high degree of
   accuracy, without having to rely on string parsing or Golang's
   :go:type:`math/big.*` values.

   .. go:function:: func (TAI64NARUXTime) Add(z TAI64NARUXTime) TAI64NARUXTime

      :param z: The timestamp to add to the current one.
      :type z: :go:type:`TAI64NARUXTime`
      :return: The sum of the two timestamps.
      :rtype: :go:type:`TAI64NARUXTime`

      Calculates the sum of two :go:type:`TAI64NARUXTime` values.

   .. go:function:: func (TAI64NARUXTime) Sub(z TAI64NARUXTime) TAI64NARUXTime

      :param z: The timestamp to subtract from the current one.
      :type z: :go:type:`TAI64NARUXTime`
      :return: The difference of the two timestamps.
      :rtype: :go:type:`TAI64NARUXTime`

      Calculates the difference of two :go:type:`TAI64NARUXTime` values.

   .. go:function:: func (TAI64NARUXTime) String() string

      :return: The decimal string representation of the current timestamp.
      :rtype: :go:type:`string`

      Returns the decimal string representation of the :go:type:`TAI64NARUXTime`
      value.

   .. go:function:: func (TAI64NARUXTime) HexString() string

      :return: The hexadecimal string representation of the current timestamp.
      :rtype: :go:type:`string`

      Returns the hexadecimal string representation of the
      :go:type:`TAI64NARUXTime` value.

   .. go:function:: func (TAI64NARUXTime) Float() Float

      :return: The arbitrary-precision floating point representation of the
               current timestamp.
      :rtype: :go:type:`math/big.(*Float)`

      Returns the :go:type:`math/big.(*Float)` representation of the
      :go:type:`TAI64NARUXTime` value.

   .. go:function:: func (TAI64NARUXTime) MarshalText() ([]byte, error)

      :return: A byte slice containing the marshalled text.
      :rtype: :go:type:`[]byte`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Implements the :go:type:`encoding.TextMarshaler` interface.

   .. go:function:: func (TAI64NARUXTime) UnmarshalText(in []byte) error

      :param in: A byte slice containing the marshalled text.
      :type in: :go:type:`[]byte`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Implements the :go:type:`encoding.TextUnmarshaler` interface.

   .. go:function:: func (TAI64NARUXTime) MarshalBinary() ([]byte, error)

      :return: A byte slice containing the marshalled binary data.
      :rtype: :go:type:`[]byte`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Implements the :go:type:`encoding.BinaryMarshaler` interface.

   .. go:function:: func (TAI64NARUXTime) UnmarshalBinary(in []byte) error

      :param in: A byte slice containing the marshalled binary data.
      :type in: :go:type:`[]byte`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Implements the :go:type:`encoding.BinaryUnmarshaler` interface.

Helpers
-------

.. go:function:: func TAI64NARUXTimeFromDecimalString(in string) TAI64NARUXTime

   :param in: The decimal string representation of a timestamp to calculate.
   :type in: :go:type:`string`
   :return: The calculated timestamp.
   :rtype: :go:type:`TAI64NARUXTime`

   Calculates a :go:type:`TAI64NARUXTime` from its decimal string
   representation.

.. go:function:: func TAI64NARUXTimeFromHexString(in string) TAI64NARUXTime

   :param in: The hexadecimal string representation of a timestamp to calculate.
   :type in: :go:type:`string`
   :return: The calculated timestamp.
   :rtype: :go:type:`TAI64NARUXTime`

   Calculates a :go:type:`TAI64NARUXTime` from its hexadecimal string
   representation.

.. go:function:: func TAI64NARUXTimeFromFloat(in Float) TAI64NARUXTime

   :param in: The arbitrary-precision floating point representation of a
              timestamp to calculate.
   :type in: :go:type:`math/big.Float`
   :return: The calculated timestamp.
   :rtype: :go:type:`TAI64NARUXTime`

   Calculates a :go:type:`TAI64NARUXTime` from its :go:type:`math/big.Float`
   representation.

.. go:function:: func UTCtoTAI(utc TAI64NARUXTime) TAI64NARUXTime

   :param utc: The timestamp to remove the UTC offset from.
   :type utc: :go:type:`TAI64NARUXTime`
   :return: The calculated timestamp.
   :rtype: :go:type:`TAI64NARUXTime`

   Removes the UTC leap second offset from a TAI64NARUXTime value.

.. go:function:: func TAItoUTC(tai TAI64NARUXTime) TAI64NARUXTime

   :param tai: The timestamp to add the UTC offset to.
   :type tai: :go:type:`TAI64NARUXTime`
   :return: The calculated timestamp.
   :rtype: :go:type:`TAI64NARUXTime`

   Adds the UTC leap second offset to a TAI64NARUXTime value.

Errors
------

.. go:type:: ErrUnsupportedInput

   Used to indicate that the input date/time weren't recognized by the calendar
   system, or that the data type is incorrect.

.. go:type:: ErrInvalidFormat

   Indicates that the ``format`` string isn't supported by the calendar system.

.. go:function:: func ErrUnknownCalendar(calendar string) error

   :param in: The name of the unknown calendar system.
   :type in: :go:type:`string`
   :return: Any error that occurs.
   :rtype: :go:type:`error`

   Generates a "calendar not registered" error, including the calendar's actual
   name in the error message.
