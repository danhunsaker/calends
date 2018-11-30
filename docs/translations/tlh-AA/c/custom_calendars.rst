.. _custom-calendars-c:

.. index::
   pair: custom calendars; C/C++

tIgh 'ISjaH qaStaHvIS C/C++
=========================

Adding new calendars to Calends is a fairly straightforward process. Implement a
handful of functions, and then simply pass them to the registration function.

Define
------

'eH Qap je yu' rur ghu'vam:

.. c:function:: TAI64Time Calends_calendar_to_internal_string(char* calendar, char* date, char* format)
                TAI64Time Calends_calendar_to_internal_long_long(char* calendar, long long int date, char* format)
                TAI64Time Calends_calendar_to_internal_double(char* calendar, double date, char* format)
                TAI64Time Calends_calendar_to_internal_tai(char* calendar, TAI64Time date)

   :param calendar: The name of the target calendar system.
   :type calendar: :c:type:`char*`
   :param date: The input date.
   :type date: :c:type:`char*` or :c:type:`long long int` or :c:type:`double`
               or :c:type:`TAI64Time`
   :param format: The format string for parsing the input date.
   :type format: :c:type:`char*`
   :return: The parsed internal timestamp.
   :rtype: :c:type:`TAI64Time`

   Converts an input date/time representation to an internal
   :c:type:`TAI64Time`.

.. c:function:: char* Calends_calendar_from_internal(char* calendar, TAI64Time stamp, char* format)

   :param calendar: The name of the target calendar system.
   :type calendar: :c:type:`char*`
   :param stamp: The internal timestamp value.
   :type stamp: :c:type:`TAI64Time`
   :param format: The format string for formatting the output date.
   :type format: :c:type:`char*`
   :return: The formatted date/time.
   :rtype: :c:type:`char*`

   Converts an internal :c:type:`TAI64Time` to a date/time string.

.. c:function:: TAI64Time Calends_calendar_offset_string(char* calendar, TAI64Time stamp, char* offset)
                TAI64Time Calends_calendar_offset_long_long(char* calendar, TAI64Time stamp, long long int offset)
                TAI64Time Calends_calendar_offset_double(char* calendar, TAI64Time stamp, double offset)
                TAI64Time Calends_calendar_offset_tai(char* calendar, TAI64Time stamp, TAI64Time offset)

   :param calendar: The name of the target calendar system.
   :type calendar: :c:type:`char*`
   :param stamp: The internal timestamp value.
   :type stamp: :c:type:`TAI64Time`
   :param offset: The input offset.
   :type offset: :c:type:`char*` or :c:type:`long long int` or :c:type:`double`
                 or :c:type:`TAI64Time`
   :return: The adjusted internal timestamp.
   :rtype: :c:type:`TAI64Time`

   Adds the given offset to an internal :c:type:`TAI64Time`.

Registration
------------

Register
::::::::

Once it is registered with the library, your calendar system can be used from
anywhere in your application. To register a system, pass it to the following
function:

.. c:function:: void Calends_calendar_register(char* name, char* defaultFormat, Calends_calendar_to_internal_string() to_internal_string, Calends_calendar_to_internal_long_long() to_internal_long_long, Calends_calendar_to_internal_double() to_internal_double, Calends_calendar_to_internal_tai() to_internal_tai, Calends_calendar_from_internal() from_internal, Calends_calendar_offset_string() offset_string, Calends_calendar_offset_long_long() offset_long_long, Calends_calendar_offset_double() offset_double, Calends_calendar_offset_tai() offset_tai)

   :param name: The name to register the calendar system under.
   :type name: :c:type:`char*`
   :param defaultFormat: The default format string.
   :type defaultFormat: :c:type:`char*`
   :param to_internal_string: The calendar parser, for :c:type:`char*` input.
   :type to_internal_string: :c:func:`Calends_calendar_to_internal_string`
   :param to_internal_long_long: The calendar parser, for ``long long int``
                                 input.
   :type to_internal_long_long: :c:func:`Calends_calendar_to_internal_long_long`
   :param to_internal_double: The calendar parser, for :c:type:`double` input.
   :type to_internal_double: :c:func:`Calends_calendar_to_internal_double`
   :param to_internal_tai: The calendar parser, for :c:type:`TAI64Time` input.
   :type to_internal_tai: :c:func:`Calends_calendar_to_internal_tai`
   :param from_internal: The calendar formatter.
   :type from_internal: :c:func:`Calends_calendar_from_internal`
   :param offset_string: The calendar offset calculator, for :c:type:`char*`
                         input.
   :type offset_string: :c:func:`Calends_calendar_offset_string`
   :param offset_long_long: The calendar offset calculator, for ``long long
                            int`` input.
   :type offset_long_long: :c:func:`Calends_calendar_offset_long_long`
   :param offset_double: The calendar offset calculator, for :c:type:`double`
                         input.
   :type offset_double: :c:func:`Calends_calendar_offset_double`
   :param offset_tai: The calendar offset calculator, for :c:type:`TAI64Time`
                      input.
   :type offset_tai: :c:func:`Calends_calendar_offset_tai`

   Registers a calendar system class, storing the collected functions as
   ``name``, and saving ``defaultFormat`` for later use while parsing or
   formatting.

Unregister
::::::::::

.. c:function:: void Calends_calendar_unregister(char* name)

   :param name: The name of the calendar system to remove.
   :type name: :c:type:`char*`

   Removes a calendar system from the callback list.

Check and List
::::::::::::::

.. c:function:: bool Calends_calendar_registered(char* name)

   :param name: The calendar system name to check for.
   :type name: :c:type:`char*`
   :return: Whether or not the calendar system is currently registered.
   :rtype: :c:type:`bool`

   Returns whether or not a calendar system has been registered, yet.

.. c:function:: char* Calends_calendar_list_registered()

   :return: The sorted list of calendar systems currently registered.
   :rtype: :c:type:`char*`

   Returns the list of calendar systems currently registered.

Types and Values
----------------

Now we get to the inner workings that make calendar systems function – even the
built-in ones. The majority of the "magic" comes from the :c:type:`TAI64Time`
struct itself, as a reliable way of storing the exact instants being calculated,
and the only way times are handled by the library itself. A handful of functions
provide basic operations that calendar system developers can use to simplify
their conversions (adding and subtracting the values of other timestamps, and
importing/exporting timestamp values from/to other types, in particular), and a
couple of helpers exclusively handle adding and removing UTC leap second
offsets. As long as you can convert your dates to/from Unix timestamps in a
:c:type:`char*`, :c:type:`long long int`, or :c:type:`double`, the rest is
handled entirely by these helpers in the library itself.

.. c:type:: TAI64Time

   Stores a ``TAI64NAXUR`` instant in a reliable, easy-converted format. Each
   9-digit fractional segment is stored in a separate 32-bit integer to preserve
   its value with a very high degree of accuracy, without having to rely on
   string parsing or external arbitrary-precision math libraries.

   .. c:member:: long long int seconds

        Seconds since ``CE 1970-01-01 00:00:00 TAI``

   .. c:member:: unsigned int nano

        Nanoseconds since the given second

   .. c:member:: unsigned int atto

        Attoseconds since the given nanosecond

   .. c:member:: unsigned int xicto

        Xictoseconds since the given attosecond

   .. c:member:: unsigned int ucto

        Uctoseconds since the given xictosecond

   .. c:member:: unsigned int rocto

        Roctoseconds since the given uctosecond

   .. c:member:: unsigned int padding

        Unused, except to round the value out to the nearest 64 bits

Calculations
------------

.. c:function:: TAI64Time TAI64Time_add(TAI64Time t, TAI64Time z)

   :param t: The current timestamp.
   :type t: :c:type:`TAI64Time`
   :param z: The timestamp to add to the current one.
   :type z: :c:type:`TAI64Time`
   :return: The sum of the two timestamps.
   :rtype: :c:type:`TAI64Time`

   Calculates the sum of two :c:type:`TAI64Time` values.

.. c:function:: TAI64Time TAI64Time_sub(TAI64Time t, TAI64Time z)

   :param t: The current timestamp.
   :type t: :c:type:`TAI64Time`
   :param z: The timestamp to subtract from the current one.
   :type z: :c:type:`TAI64Time`
   :return: The difference of the two timestamps.
   :rtype: :c:type:`TAI64Time`

   Calculates the difference of two :c:type:`TAI64Time` values.

Export
------

.. c:function:: char* TAI64Time_string(TAI64Time t)

   :param t: The current timestamp.
   :type t: :c:type:`TAI64Time`
   :return: The decimal string representation of the current timestamp.
   :rtype: :c:type:`char*`

   Returns the decimal string representation of a :c:type:`TAI64Time` value.

.. c:function:: TAI64Time TAI64Time_from_string(char* in)

   :param in: The decimal string representation of a timestamp to calculate.
   :type in: :c:type:`char*`
   :return: The calculated timestamp.
   :rtype: :c:type:`TAI64Time`

   Calculates a :c:type:`TAI64Time` from its decimal string representation.

.. c:function:: char* TAI64Time_hex_string(TAI64Time t)

   :param t: The current timestamp.
   :type t: :c:type:`TAI64Time`
   :return: The hexadecimal string representation of the current timestamp.
   :rtype: :c:type:`char*`

   Returns the hexadecimal string representation of a :c:type:`TAI64Time` value.

.. c:function:: TAI64Time TAI64Time_from_hex_string(char* in)

   :param in: The hexadecimal string representation of a timestamp to calculate.
   :type in: :c:type:`char*`
   :return: The calculated timestamp.
   :rtype: :c:type:`TAI64Time`

   Calculates a :c:type:`TAI64Time` from its hexadecimal string representation.

.. c:function:: double TAI64Time_double(TAI64Time t)

   :param t: The current timestamp.
   :type t: :c:type:`TAI64Time`
   :return: The arbitrary-precision floating point representation of the
            current timestamp.
   :rtype: :c:type:`double`

   Returns the :c:type:`double` representation of a :c:type:`TAI64Time` value.

.. c:function:: TAI64Time TAI64Time_from_double(double in)

   :param in: The arbitrary-precision floating point representation of a
              timestamp to calculate.
   :type in: :c:type:`double`
   :return: The calculated timestamp.
   :rtype: :c:type:`TAI64Time`

   Calculates a :c:type:`TAI64Time` from its :c:type:`double` representation.

.. c:function:: char* TAI64Time_encode_text(TAI64Time t)

   :param t: The current timestamp.
   :type t: :c:type:`TAI64Time`
   :return: A string containing the encoded text.
   :rtype: :c:type:`char*`

   Encodes a :c:type:`TAI64Time` value as text.

.. c:function:: TAI64Time TAI64Time_decode_text(char* in)

   :param in: A string containing the encoded text.
   :type in: :c:type:`char*`
   :return: The decoded timestamp.
   :rtype: :c:type:`TAI64Time`

   Decodes a :c:type:`TAI64Time` value from text.

.. c:function:: void* TAI64Time_encode_binary(TAI64Time t, int *len)

   :param t: The current timestamp.
   :type t: :c:type:`TAI64Time`
   :param len: Will return the length of the binary data.
   :type len: :c:type:`int*`
   :return: A pointer to the encoded binary data stream.
   :rtype: :c:type:`void*`

   Encodes a :c:type:`TAI64Time` value as a binary data stream.

.. c:function:: TAI64Time TAI64Time_decode_binary(void* in, int len)

   :param in: A pointer to the encoded binary data stream.
   :type in: :c:type:`void*`
   :param len: The length of the binary data.
   :type len: :c:type:`int`
   :return: The decoded timestamp.
   :rtype: :c:type:`TAI64Time`

   Decodes a :c:type:`TAI64Time` value from a binary data stream.

Helpers
-------

.. c:function:: TAI64Time TAI64Time_utc_to_tai(TAI64Time utc)

   :param utc: The timestamp to remove the UTC offset from.
   :type utc: :c:type:`TAI64Time`
   :return: The calculated timestamp.
   :rtype: :c:type:`TAI64Time`

   Removes the UTC leap second offset from a :c:type:`TAI64Time` value.

.. c:function:: TAI64Time TAI64Time_tai_to_utc(TAI64Time tai)

   :param tai: The timestamp to add the UTC offset to.
   :type tai: :c:type:`TAI64Time`
   :return: The calculated timestamp.
   :rtype: :c:type:`TAI64Time`

   Adds the UTC leap second offset to a :c:type:`TAI64Time` value.
