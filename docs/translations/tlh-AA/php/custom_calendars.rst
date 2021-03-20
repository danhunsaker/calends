.. _custom-calendars-php:

.. php:namespace:: Calends

.. index::
   pair: custom calendars; PHP

Custom Calendars in PHP
==========================

'ISjaH chu' chel Calends mIw fairly straightforward. Implement
one of two interfaces, and then simply pass it to the registration method.

define
------

The more common interface looks like this:

.. php:interface:: CalendarInterface

   .. php:staticmethod:: ToInternal(mixed $date, string $format): TAITime

      :param $date: The input date. Should support strings at the very minimum.
      :type $date: ``mixed``
      :param $format: The format string for parsing the input date.
      :type $format: ``string``
      :return: parse internal timestamp.
      :rtype: :php:class:`TAITime`
      :throws CalendsException: when an error occurs

      representation input date ghap poH bIDameH internal
      :php:class:`TAITime`.

   .. php:staticmethod:: FromInternal(TAITime $stamp, string $format): string

      :param $stamp: The internal timestamp value.
      :type $stamp: :php:class:`TAITime`
      :param $format: The format string for formatting the output date.
      :type $format: ``string``
      :return: format date ghap poH.
      :rtype: ``string``
      :throws CalendsException: when an error occurs

      Converts an internal :php:class:`TAITime` to a date/time string.

   .. php:staticmethod:: Offset(TAITime $stamp, mixed $offset): TAITime

      :param $stamp: The internal timestamp value.
      :type $stamp: :php:class:`TAITime`
      :param $offset: The input offset. Should support strings at the very
                      minimum.
      :type $offset: ``mixed``
      :return: lIS internal timestamp.
      :rtype: :php:class:`TAITime`
      :throws CalendsException: when an error occurs

      Adds the given offset to an internal :php:class:`TAITime`.

The other is virtually identical, but uses object instances instead of static
class methods:

.. php:interface:: CalendarObjectInterface

   .. php:method:: ToInternal(mixed $date, string $format): TAITime

      :param $date: The input date. Should support strings at the very minimum.
      :type $date: ``mixed``
      :param $format: The format string for parsing the input date.
      :type $format: ``string``
      :return: parse internal timestamp.
      :rtype: :php:class:`TAITime`
      :throws CalendsException: when an error occurs

      representation input date ghap poH bIDameH internal
      :php:class:`TAITime`.

   .. php:method:: FromInternal(TAITime $stamp, string $format): string

      :param $stamp: The internal timestamp value.
      :type $stamp: :php:class:`TAITime`
      :param $format: The format string for formatting the output date.
      :type $format: ``string``
      :return: format date ghap poH.
      :rtype: ``string``
      :throws CalendsException: when an error occurs

      Converts an internal :php:class:`TAITime` to a date/time string.

   .. php:method:: Offset(TAITime $stamp, mixed $offset): TAITime

      :param $stamp: The internal timestamp value.
      :type $stamp: :php:class:`TAITime`
      :param $offset: The input offset. Should support strings at the very
                      minimum.
      :type $offset: ``mixed``
      :return: lIS internal timestamp.
      :rtype: :php:class:`TAITime`
      :throws CalendsException: when an error occurs

      Adds the given offset to an internal :php:class:`TAITime`.

registration
------------

Suy qachmey
::::::::

wa'logh Suy qachmey 'oH je be'nI''a'wI', Datu', laH yIlo' 'ISjaH pat vo'
anywhere neH application. pat Suy qachmey, 'oH juS toblu'
Qap:

.. php:class:: Calends

.. php:staticmethod:: calendarRegister(string $name, string $defaultFormat, mixed $calendar)

   :param $name: The name to register the calendar system under.
   :type $name: ``string``
   :param $defaultFormat: The default format string.
   :type $defaultFormat: ``string``
   :param $calendar: The calendar system itself.
   :type $calendar: :php:interface:`CalendarInterface` or
                    :php:interface:`CalendarObjectInterface`

   Registers a calendar system class or object, storing ``$calendar`` as
   ``$name``, and saving ``$defaultFormat`` for later use while parsing or
   formatting.

unregister
::::::::::

.. php:staticmethod:: calendarUnregister(string $name)

   :param $name: The name of the calendar system to remove.
   :type $name: ``string``

   ['ISjaH pat vo' callback tetlh. ghaHDaq teq BERNARDO.

check 'ej tetlh
::::::::::::::

.. php:staticmethod:: calendarRegistered(string $name): bool

   :param $name: The calendar system name to check for.
   :type $name: ``string``
   :return: chaq pagh wej currently Suy qachmey 'ISjaH pat.
   :rtype: ``bool``

   chegh chaq pagh wej Suy qachmey 'ISjaH pat, 'ach.

.. php:staticmethod:: calendarListRegistered(): array

   :return: currently Suy qachmey Segh tetlh 'ISjaH pat.
   :rtype: ``[string]``

   'ISjaH pat currently Suy qachmey tetlh chegh.

Segh lo'laHghach je
----------------

DaH inner workings San 'ISjaH pat Qap – wej qaSpu'bogh
built-in wa'. The majority of the "magic" comes from the :php:class:`TAITime`
object itself, as a reliable way of storing the exact instants being calculated,
'ej DeS neH mIw poH pong be'nI''a'wI', Datu' narghtaHvIS 'oH. A handful of methods
basic yo'SeH laH DanoHmeH 'ISjaH pat developers simplify 'e' DuHIvDI'
conversions (chel 'ej latlh timestamps nIv boqHa' je
importing/exporting timestamp values from/to string and numeric types, in
particular), and a couple of helpers exclusively handle adding and removing UTC
leap second offsets. As long as you can convert your dates to/from Unix
timestamps in a string or numeric type, the rest is handled entirely by these
helpers in the library itself.

.. php:class:: TAITime

   :php:class:`TAITime` stores a ``TAI64NAXUR`` instant in a reliable,
   easily-converted format. Each 9-digit fractional segment is stored in a
   separate 32-bit integer to preserve its value with a very high degree of
   accuracy, without having to rely on string parsing or external
   arbitrary-precision mathematics libraries.

   .. php:attr:: seconds (float)

      The number of TAI seconds since ``CE 1970-01-01 00:00:00 TAI``. Should be an integer value; the ``float`` type is used, here, only to be able to hold a full signed 64-bit integer value regardless of architecture.

   .. php:attr:: nano (integer)

      The first 9 digits of the timestamp's fractional component.

   .. php:attr:: atto (integer)

      The 10th through 18th digits of the fractional component.

   .. php:attr:: xicto (integer)

      The 19th through 27th digits of the fractional component.

   .. php:attr:: ucto (integer)

      The 28th through 36th digits of the fractional component.

   .. php:attr:: rocto (integer)

      The 37th through 45th digits of the fractional component.

   .. php:method:: add(TAITime $z): TAITime

      :param $z: The timestamp to add to the current one.
      :type $z: :php:class:`TAITime`
      :return: sum timestamps cha'.
      :rtype: :php:class:`TAITime`

      Calculates the sum of two :php:class:`TAITime` values.

   .. php:method:: sub(TAITime $z): TAITime

      :param $z: The timestamp to subtract from the current one.
      :type $z: :php:class:`TAITime`
      :return: difference timestamps cha'.
      :rtype: :php:class:`TAITime`

      Calculates the difference of two :php:class:`TAITime` values.

   .. php:method:: toString(): string

      :return: decimal SIrgh representation Qu'mey potlh timestamp.
      :rtype: ``string``

      Returns the decimal string representation of the :php:class:`TAITime`
      value.

      .. Note::

         :php:class:`TAITime` also implements :php:meth:`!__toString`, so you
         can use that instead of calling this function directly, if you prefer.

   .. php:method:: fromString(string $in): TAITime

      :param $in: The decimal string representation of a timestamp to calculate.
      :type $in: string
      :return: SIm timestamp.
      :rtype: :php:class:`TAITime`

      Calculates a :php:class:`TAITime` from its decimal string representation.

   .. php:method:: toHex(): string

      :return: hexadecimal SIrgh representation Qu'mey potlh timestamp.
      :rtype: ``string``

      Returns the hexadecimal string representation of the :php:class:`TAITime`
      value.

   .. php:method:: fromHex(string $in):TAITime

      :param $in: The hexadecimal string representation of a timestamp to calculate.
      :type $in: string
      :return: SIm timestamp.
      :rtype: :php:class:`TAITime`

      Calculates a :php:class:`TAITime` from its hexadecimal string
      representation.

   .. php:method:: toNumber(): float

      :return: The numeric representation of the current timestamp.
      :rtype: ``float``

      Returns the ``float`` representation of the :php:class:`TAITime` value.

   .. php:method:: fromNumber(numeric $in): TAITime

      :param $in: The arbitrary-precision floating point representation of a
                 Qo'noS timestamp SIm.
      :type $in: ``integer`` or ``float``
      :return: SIm timestamp.
      :rtype: :php:class:`TAITime`

      Calculates a :php:class:`TAITime` from its numeric (``integer`` or
      ``float``) representation.

   .. php:method:: fromUTC(): TAITime

      :return: SIm timestamp.
      :rtype: :php:class:`TAITime`

      Removes the UTC leap second offset from a TAITime value.

   .. php:method:: toUTC(): TAITime

      :return: SIm timestamp.
      :rtype: :php:class:`TAITime`

      Adds the UTC leap second offset to a TAITime value.
