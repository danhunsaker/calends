.. _custom-calendars-go:

.. go:currentpackage:: calends/calendars

.. index::
   pair: custom calendars; Golang

Custom Calendars in Golang
==========================

'ISjaH chu' chel Calends mIw fairly straightforward. Implement
one interface, or its three methods as standalone functions, and then simply
pass them to one of the two registration functions.

define
------

The interface in question looks like this:

.. go:package:: calends/calendars

.. go:type:: CalendarDefinition

   .. go:function:: func (CalendarDefinition) ToInternal(date interface, format string) (TAI64NAXURTime, error)

      :param date: The input date. Should support :go:type:`string` at the very
                   minimum.
      :type date: :go:type:`interface{}`
      :param format: format SIrgh input date parsing.
      :type format: :go:type:`string`
      :return: parse internal timestamp.
      :rtype: :go:type:`TAI64NAXURTime`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      representation input date ghap poH bIDameH internal
      :go:type:`TAI64NAXURTime`.

   .. go:function:: func (CalendarDefinition) FromInternal(stamp TAI64NAXURTime, format string) (string, error)

      :param stamp: internal timestamp lo'laHghach.
      :type stamp: :go:type:`TAI64NAXURTime`
      :param format: format SIrgh output date formatting.
      :type format: :go:type:`string`
      :return: format date ghap poH.
      :rtype: :go:type:`string`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Converts an internal :go:type:`TAI64NAXURTime` to a date/time string.

   .. go:function:: func (CalendarDefinition) Offset(stamp TAI64NAXURTime, offset interface) (TAI64NAXURTime, error)

      :param stamp: internal timestamp lo'laHghach.
      :type stamp: :go:type:`TAI64NAXURTime`
      :param offset: input offset. Should support :go:type:`string` at the
                     very minimum.
      :type offset: :go:type:`interface{}`
      :return: lIS internal timestamp.
      :rtype: :go:type:`TAI64NAXURTime`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Adds the given offset to an internal :go:type:`TAI64NAXURTime`.

registration
------------

Suy qachmey
::::::::

wa'logh Suy qachmey 'oH je be'nI''a'wI', Datu', laH yIlo' 'ISjaH pat vo'
anywhere neH application. To register a system, pass it to one of the
following two functions:

.. go:function:: func RegisterObject(name string, definition CalendarDefinition, defaultFormat string)

   :param name: 'ISjaH pat bopummeH Suy qachmey pong.
   :type name: :go:type:`string`
   :param definition: The calendar system itself.
   :type definition: :go:type:`CalendarDefinition`
   :param defaultformat: default format Surgh.
   :type defaultFormat: :go:type:`string`

   Registers a calendar system class, storing ``definition`` as ``name``, and
   saving ``defaultFormat`` for later use while parsing or formatting.

.. go:function:: func RegisterElements(name string, toInternal ToInternal, fromInternal FromInternal, offset Offset, defaultFormat string)

   :param name: 'ISjaH pat bopummeH Suy qachmey pong.
   :type name: :go:type:`string`
   :param toInternal: The function for parsing dates into internal timestamps.
   :type toInternal: :go:func:`(CalendarDefinition) ToInternal`
   :param fromInternal: The function for formatting internal timestamps as
                        dates.
   :type fromInternal: :go:func:`(CalendarDefinition) FromInternal`
   :param offset: The function for adding an offset to internal timestamps.
   :type offset: :go:func:`(CalendarDefinition) Offset`
   :param defaultformat: default format Surgh.
   :type defaultFormat: :go:type:`string`

   Registers a calendar system from its distinct functions. It does this by
   storing ``toInternal``, ``fromInternal``, and ``offset`` as the elements of
   ``name`` 'ej ``defaultFormat`` later lo' poStaHvIS parsing toD pagh
   formatting.

unregister
::::::::::

.. go:function:: func Unregister(name string)

   :param name: 'ISjaH pat teq pong.
   :type name: :go:type:`string`

   ['ISjaH pat vo' callback tetlh. ghaHDaq teq BERNARDO.

check 'ej tetlh
::::::::::::::

.. go:function:: func Registered(calendar string) bool

   :param name: 'ISjaH pat pong check.
   :type name: :go:type:`string`
   :return: chaq pagh wej currently Suy qachmey 'ISjaH pat.
   :rtype: :go:type:`bool`

   chegh chaq pagh wej Suy qachmey 'ISjaH pat, 'ach.

.. go:function:: func ListRegistered() []string

   :return: currently Suy qachmey Segh tetlh 'ISjaH pat.
   :rtype: :go:type:`[]string`

   'ISjaH pat currently Suy qachmey tetlh chegh.

Segh lo'laHghach je
----------------

DaH inner workings San 'ISjaH pat Qap – wej qaSpu'bogh
built-in wa'. The majority of the "magic" comes from the
:go:type:`TAI64NAXURTime` object itself, as a reliable way of storing the exact
instants being calculated, and the only way times are handled by the library
itself. A handful of methods provide basic operations that calendar system
developers can use to simplify their conversions (adding and subtracting the
values of other timestamps, and importing/exporting timestamp values from/to
arbitrary-precision floating point :go:type:`math/big.Float`\ s, in particular),
and a couple of helpers exclusively handle adding and removing UTC leap second
offsets. Hoch nI' law' dates laH bIDameH SoH ghap vo' unix timestamps neH
:go:type:`string` or :go:type:`math/big.Float`, the rest is handled entirely by
these helpers in the library itself.

.. go:type:: TAI64NAXURTime

   :param int64 Seconds: The number of TAI seconds since ``CE 1970-01-01
                         00:00:00 TAI``.
   :param uint32 Nano: The first 9 digits of the timestamp's fractional
                       component.
   :param uint32 Atto: The 10th through 18th digits of the fractional component.
   :param uint32 Xicto: The 19th through 27th digits of the fractional
                        component.
   :param uint32 Ucto: The 28th through 36th digits of the fractional component.
   :param uint32 Rocto: The 37th through 45th digits of the fractional
                        component.

   :go:type:`TAI64NAXURTime` stores a ``TAI64NAXUR`` instant in a reliable,
   easy-converted format. Each 9-digit fractional segment is stored in a
   separate 32-bit integer to preserve its value with a very high degree of
   accuracy, without having to rely on string parsing or Golang's
   :go:type:`math/big.*` values.

   .. go:function:: func (TAI64NAXURTime) Add(z TAI64NAXURTime) TAI64NAXURTime

      :param z: chel wa' Qu'mey potlh timestamp.
      :type z: :go:type:`TAI64NAXURTime`
      :return: sum timestamps cha'.
      :rtype: :go:type:`TAI64NAXURTime`

      Calculates the sum of two :go:type:`TAI64NAXURTime` values.

   .. go:function:: func (TAI64NAXURTime) Sub(z TAI64NAXURTime) TAI64NAXURTime

      :param z: boqHa' vo' wa' Qu'mey potlh timestamp.
      :type z: :go:type:`TAI64NAXURTime`
      :return: difference timestamps cha'.
      :rtype: :go:type:`TAI64NAXURTime`

      Calculates the difference of two :go:type:`TAI64NAXURTime` values.

   .. go:function:: func (TAI64NAXURTime) String() string

      :return: decimal SIrgh representation Qu'mey potlh timestamp.
      :rtype: :go:type:`string`

      Returns the decimal string representation of the :go:type:`TAI64NAXURTime`
      value.

   .. go:function:: func (TAI64NAXURTime) HexString() string

      :return: hexadecimal SIrgh representation Qu'mey potlh timestamp.
      :rtype: :go:type:`string`

      Returns the hexadecimal string representation of the
      :go:type:`TAI64NAXURTime` value.

   .. go:function:: func (TAI64NAXURTime) Float() Float

      :return: arbitrary-precision 'ej Dunbogh jom lang representation
               Qu'mey potlh timestamp.
      :rtype: :go:type:`math/big.(*Float)`

      Returns the :go:type:`math/big.(*Float)` representation of the
      :go:type:`TAI64NAXURTime` value.

   .. go:function:: func (TAI64NAXURTime) MarshalText() ([]byte, error)

      :return: A byte slice containing the marshalled text.
      :rtype: :go:type:`[]byte`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Implements the :go:type:`encoding.TextMarshaler` interface.

   .. go:function:: func (TAI64NAXURTime) UnmarshalText(in []byte) error

      :param in: A byte slice containing the marshalled text.
      :type in: :go:type:`[]byte`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Implements the :go:type:`encoding.TextUnmarshaler` interface.

   .. go:function:: func (TAI64NAXURTime) MarshalBinary() ([]byte, error)

      :return: A byte slice containing the marshalled binary data.
      :rtype: :go:type:`[]byte`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Implements the :go:type:`encoding.BinaryMarshaler` interface.

   .. go:function:: func (TAI64NAXURTime) UnmarshalBinary(in []byte) error

      :param in: A byte slice containing the marshalled binary data.
      :type in: :go:type:`[]byte`
      :return: Any error that occurs.
      :rtype: :go:type:`error`

      Implements the :go:type:`encoding.BinaryUnmarshaler` interface.

helper
-------

.. go:function:: func TAI64NAXURTimeFromDecimalString(in string) TAI64NAXURTime

   :param in: decimal SIrgh representation timestamp SIm.
   :type in: :go:type:`string`
   :return: SIm timestamp.
   :rtype: :go:type:`TAI64NAXURTime`

   Calculates a :go:type:`TAI64NAXURTime` from its decimal string
   representation.

.. go:function:: func TAI64NAXURTimeFromHexString(in string) TAI64NAXURTime

   :param in: hexadecimal SIrgh representation timestamp SIm.
   :type in: :go:type:`string`
   :return: SIm timestamp.
   :rtype: :go:type:`TAI64NAXURTime`

   Calculates a :go:type:`TAI64NAXURTime` from its hexadecimal string
   representation.

.. go:function:: func TAI64NAXURTimeFromFloat(in Float) TAI64NAXURTime

   :param in: arbitrary-precision 'ej Dunbogh jom lang representation
              Qo'noS timestamp SIm.
   :type in: :go:type:`math/big.Float`
   :return: SIm timestamp.
   :rtype: :go:type:`TAI64NAXURTime`

   Calculates a :go:type:`TAI64NAXURTime` from its :go:type:`math/big.Float`
   representation.

.. go:function:: func UTCtoTAI(utc TAI64NAXURTime) TAI64NAXURTime

   :param utc: timestamp UTC offset vo' teq.
   :type utc: :go:type:`TAI64NAXURTime`
   :return: SIm timestamp.
   :rtype: :go:type:`TAI64NAXURTime`

   Removes the UTC leap second offset from a TAI64NAXURTime value.

.. go:function:: func TAItoUTC(tai TAI64NAXURTime) TAI64NAXURTime

   :param tai: timestamp UTC offset chel.
   :type tai: :go:type:`TAI64NAXURTime`
   :return: SIm timestamp.
   :rtype: :go:type:`TAI64NAXURTime`

   Adds the UTC leap second offset to a TAI64NAXURTime value.

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
