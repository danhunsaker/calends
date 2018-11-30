.. _custom-calendars-c:

.. index::
   pair: custom calendars; C/C++

tIgh 'ISjaH qaStaHvIS C/C++
=========================

'ISjaH chu' chel Calends mIw fairly straightforward. implement
Qap, handful 'ej ghIq simply chaH Qap registration.

define
------

'eH Qap je yu' rur ghu'vam:

.. c:function:: TAI64Time Calends_calendar_to_internal_string(char* calendar, char* date, char* format)
                TAI64Time Calends_calendar_to_internal_long_long(char* calendar, long long int date, char* format)
                TAI64Time Calends_calendar_to_internal_double(char* calendar, double date, char* format)
                TAI64Time Calends_calendar_to_internal_tai(char* calendar, TAI64Time date)

   :param calendar: DoS 'ISjaH pat pong.
   :type calendar: :c:type:`char*`
   :param date: The input date.
   :type date: :c:type:`char*` or :c:type:`long long int` or :c:type:`double`
               or :c:type:`TAI64Time`
   :param format: format SIrgh input date parsing.
   :type format: :c:type:`char*`
   :return: parse internal timestamp.
   :rtype: :c:type:`TAI64Time`

   representation input date ghap poH bIDameH internal
   :c:type:`TAI64Time`.

.. c:function:: char* Calends_calendar_from_internal(char* calendar, TAI64Time stamp, char* format)

   :param calendar: DoS 'ISjaH pat pong.
   :type calendar: :c:type:`char*`
   :param stamp: internal timestamp lo'laHghach.
   :type stamp: :c:type:`TAI64Time`
   :param format: format SIrgh output date formatting.
   :type format: :c:type:`char*`
   :return: format date ghap poH.
   :rtype: :c:type:`char*`

   internal bIDameH :c:type:`TAI64Time` date ghap poH Surgh.

.. c:function:: TAI64Time Calends_calendar_offset_string(char* calendar, TAI64Time stamp, char* offset)
                TAI64Time Calends_calendar_offset_long_long(char* calendar, TAI64Time stamp, long long int offset)
                TAI64Time Calends_calendar_offset_double(char* calendar, TAI64Time stamp, double offset)
                TAI64Time Calends_calendar_offset_tai(char* calendar, TAI64Time stamp, TAI64Time offset)

   :param calendar: DoS 'ISjaH pat pong.
   :type calendar: :c:type:`char*`
   :param stamp: internal timestamp lo'laHghach.
   :type stamp: :c:type:`TAI64Time`
   :param offset: input offset.
   :type offset: :c:type:`char*` or :c:type:`long long int` or :c:type:`double`
                 or :c:type:`TAI64Time`
   :return: lIS internal timestamp.
   :rtype: :c:type:`TAI64Time`

   nob offset chel internal :c:type:`TAI64Time`.

registration
------------

Suy qachmey
::::::::

wa'logh Suy qachmey 'oH je be'nI''a'wI', Datu', laH yIlo' 'ISjaH pat vo'
anywhere neH application. pat Suy qachmey, 'oH juS toblu'
Qap:

.. c:function:: void Calends_calendar_register(char* name, char* defaultFormat, Calends_calendar_to_internal_string() to_internal_string, Calends_calendar_to_internal_long_long() to_internal_long_long, Calends_calendar_to_internal_double() to_internal_double, Calends_calendar_to_internal_tai() to_internal_tai, Calends_calendar_from_internal() from_internal, Calends_calendar_offset_string() offset_string, Calends_calendar_offset_long_long() offset_long_long, Calends_calendar_offset_double() offset_double, Calends_calendar_offset_tai() offset_tai)

   :param name: 'ISjaH pat bopummeH Suy qachmey pong.
   :type name: :c:type:`char*`
   :param defaultformat: default format Surgh.
   :type defaultFormat: :c:type:`char*`
   :param to_internal_string: parser 'ISjaH, :c:type:`char *` input.
   :type to_internal_string: :c:func:`Calends_calendar_to_internal_string`
   :param to_internal_long_long: 'ISjaH parser, ``long long int``
                                 input.
   :type to_internal_long_long: :c:func:`Calends_calendar_to_internal_long_long`
   :param to_internal_double: parser 'ISjaH, :c:type:`double`, cha'logh vaj input.
   :type to_internal_double: :c:func:`Calends_calendar_to_internal_double`
   :param to_internal_tai: parser 'ISjaH, :c:type:`TAI64Time` input.
   :type to_internal_tai: :c:func:`Calends_calendar_to_internal_tai`
   :param from_internal: 'ISjaH formatter.
   :type from_internal: :c:func:`Calends_calendar_from_internal`
   :param offset_string: calculator, offset 'ISjaH :c:type:`char *`
                         input.
   :type offset_string: :c:func:`Calends_calendar_offset_string`
   :param offset_long_long: calculator, offset 'ISjaH ``long long
                            int`` input.
   :type offset_long_long: :c:func:`Calends_calendar_offset_long_long`
   :param offset_double: calculator, offset 'ISjaH :c:type:`double`, cha'logh vaj
                         input.
   :type offset_double: :c:func:`Calends_calendar_offset_double`
   :param offset_tai: calculator, offset 'ISjaH :c:type:`TAI64Time``
                      input.
   :type offset_tai: :c:func:`Calends_calendar_offset_tai`

   'ISjaH pat Segh, boS Qap je ngevwI' Suy qachmey
   ``name`` 'ej ``defaultFormat`` later lo' poStaHvIS parsing toD pagh
   formatting.

unregister
::::::::::

.. c:function:: void Calends_calendar_unregister(char* name)

   :param name: 'ISjaH pat teq pong.
   :type name: :c:type:`char*`

   ['ISjaH pat vo' callback tetlh. ghaHDaq teq BERNARDO.

check 'ej tetlh
::::::::::::::

.. c:function:: bool Calends_calendar_registered(char* name)

   :param name: 'ISjaH pat pong check.
   :type name: :c:type:`char*`
   :return: chaq pagh wej currently Suy qachmey 'ISjaH pat.
   :rtype: :c:type:`bool`

   chegh chaq pagh wej Suy qachmey 'ISjaH pat, 'ach.

.. c:function:: char* Calends_calendar_list_registered()

   :return: currently Suy qachmey Segh tetlh 'ISjaH pat.
   :rtype: :c:type:`char*`

   'ISjaH pat currently Suy qachmey tetlh chegh.

Segh lo'laHghach je
----------------

DaH inner workings San 'ISjaH pat Qap – wej qaSpu'bogh
built-in wa'. vo' "magic" majority :c:type:`TAI64Time`
struct narghtaHvIS 'oH, law' reliable mIwvam'e' pup instants SIm ngevwI'.
'ej DeS neH mIw poH pong be'nI''a'wI', Datu' narghtaHvIS 'oH. Qap handful
basic yo'SeH laH DanoHmeH 'ISjaH pat developers simplify 'e' DuHIvDI'
conversions (chel 'ej latlh timestamps nIv boqHa' je
DughajmoH ghap timestamp lo'laHghach vo' ghap latlh Segh exporting je bIH), 'ej
chel 'ej utc Sup cha'DIch teq exclusively DeS helpers couple
offsets. Hoch nI' law' dates laH bIDameH SoH ghap vo' unix timestamps neH
:c:type:`char *`, :c:type:`long long int`, pagh :c:type:`double`, cha'logh vaj, meqchaj
DeS entirely pong be'nI''a'wI', Datu' narghtaHvIS 'oH helpers.

.. c:type:: TAI64Time

   '' tai64naxur '' instant neH reliable, ngeD-bIDameH format ngevwI'. Hoch
   ngevwI' 9-digit fractional segment neH chev 32-bit integer choq
   jen qechmeyDaj Huj accuracy. ghaH Hutlh vay' wuv lo'laHghach
   SIrgh parse pagh external arbitrary-precision math be'nI''a'wI', Datu'.

   .. c:member:: long long int seconds

        cha'DIch qaSchoH ``CE 1970-01-01 00:00:00 TAI``

   .. c:member:: unsigned int nano

        nanoseconds qaSchoH nob cha'DIch

   .. c:member:: unsigned int atto

        attoseconds qaSchoH nob nanosecond

   .. c:member:: unsigned int xicto

        xictoseconds qaSchoH nob attosecond

   .. c:member:: unsigned int ucto

        uctoseconds qaSchoH nob xictosecond

   .. c:member:: unsigned int rocto

        roctoseconds qaSchoH nob uctosecond

   .. c:member:: unsigned int padding

        unused, jIQongqa'laHbe' lo'laHghach round bits nearest 64

calculation
------------

.. c:function:: TAI64Time TAI64Time_add(TAI64Time t, TAI64Time z)

   :param t: The current timestamp.
   :type t: :c:type:`TAI64Time`
   :param z: chel wa' Qu'mey potlh timestamp.
   :type z: :c:type:`TAI64Time`
   :return: sum timestamps cha'.
   :rtype: :c:type:`TAI64Time`

   sum cha' SIm :c:type:`TAI64Time` lo'laHghach.

.. c:function:: TAI64Time TAI64Time_sub(TAI64Time t, TAI64Time z)

   :param t: Qu'mey potlh timestamp.
   :type t: :c:type:`TAI64Time`
   :param z: boqHa' vo' wa' Qu'mey potlh timestamp.
   :type z: :c:type:`TAI64Time`
   :return: difference timestamps cha'.
   :rtype: :c:type:`TAI64Time`

   difference cha' SIm :c:type:`TAI64Time` lo'laHghach.

export
------

.. c:function:: char* TAI64Time_string(TAI64Time t)

   :param t: Qu'mey potlh timestamp.
   :type t: :c:type:`TAI64Time`
   :return: decimal SIrgh representation Qu'mey potlh timestamp.
   :rtype: :c:type:`char*`

   decimal SIrgh representation chegh :c:type:`TAI64Time` lo'laHghach.

.. c:function:: TAI64Time TAI64Time_from_string(char* in)

   :param in: decimal SIrgh representation timestamp SIm.
   :type in: :c:type:`char*`
   :return: SIm timestamp.
   :rtype: :c:type:`TAI64Time`

   SIm :c:type:`TAI64Time` vo' SIrgh decimal representation.

.. c:function:: char* TAI64Time_hex_string(TAI64Time t)

   :param t: Qu'mey potlh timestamp.
   :type t: :c:type:`TAI64Time`
   :return: hexadecimal SIrgh representation Qu'mey potlh timestamp.
   :rtype: :c:type:`char*`

   hexadecimal SIrgh representation chegh :c:type:`TAI64Time` lo'laHghach.

.. c:function:: TAI64Time TAI64Time_from_hex_string(char* in)

   :param in: hexadecimal SIrgh representation timestamp SIm.
   :type in: :c:type:`char*`
   :return: SIm timestamp.
   :rtype: :c:type:`TAI64Time`

   SIm :c:type:`TAI64Time` vo' SIrgh hexadecimal representation.

.. c:function:: double TAI64Time_double(TAI64Time t)

   :param t: Qu'mey potlh timestamp.
   :type t: :c:type:`TAI64Time`
   :return: arbitrary-precision 'ej Dunbogh jom lang representation
            Qu'mey potlh timestamp.
   :rtype: :c:type:`double`

   :c:type:`double` representation chegh :c:type:`TAI64Time` lo'laHghach.

.. c:function:: TAI64Time TAI64Time_from_double(double in)

   :param in: arbitrary-precision 'ej Dunbogh jom lang representation
              Qo'noS timestamp SIm.
   :type in: :c:type:`double`
   :return: SIm timestamp.
   :rtype: :c:type:`TAI64Time`

   SIm :c:type:`TAI64Time` vo' :c:type:`double` representation.

.. c:function:: char* TAI64Time_encode_text(TAI64Time t)

   :param t: Qu'mey potlh timestamp.
   :type t: :c:type:`TAI64Time`
   :return: SIrgh encode bIngDaq ghItlh leghlu' ngaS.
   :rtype: :c:type:`char*`

   encodes :c:type:`TAI64Time` lo'laHghach je bIngDaq ghItlh leghlu'.

.. c:function:: TAI64Time TAI64Time_decode_text(char* in)

   :param in: SIrgh encode bIngDaq ghItlh leghlu' ngaS.
   :type in: :c:type:`char*`
   :return: decode timestamp.
   :rtype: :c:type:`TAI64Time`

   decodes :c:type:`TAI64Time` lo'laHghach vo' bIngDaq ghItlh leghlu'.

.. c:function:: void* TAI64Time_encode_binary(TAI64Time t, int *len)

   :param t: Qu'mey potlh timestamp.
   :type t: :c:type:`TAI64Time`
   :param len: De' binary 'ab jIchegh.
   :type len: :c:type:`int*`
   :return: pointer encode binary De' stream.
   :rtype: :c:type:`void*`

   encodes :c:type:`TAI64Time` lo'laHghach je binary De' stream.

.. c:function:: TAI64Time TAI64Time_decode_binary(void* in, int len)

   :param in: pointer encode binary De' stream.
   :type in: :c:type:`void*`
   :param len: 'ab De' binary.
   :type len: :c:type:`int`
   :return: decode timestamp.
   :rtype: :c:type:`TAI64Time`

   decodes :c:type:`TAI64Time` lo'laHghach vo' binary De' stream.

helper
-------

.. c:function:: TAI64Time TAI64Time_utc_to_tai(TAI64Time utc)

   :param utc: timestamp UTC offset vo' teq.
   :type utc: :c:type:`TAI64Time`
   :return: SIm timestamp.
   :rtype: :c:type:`TAI64Time`

   UTC Sup cha'DIch offset vo' teq :c:type:`TAI64Time` lo'laHghach.

.. c:function:: TAI64Time TAI64Time_tai_to_utc(TAI64Time tai)

   :param tai: timestamp UTC offset chel.
   :type tai: :c:type:`TAI64Time`
   :return: SIm timestamp.
   :rtype: :c:type:`TAI64Time`

   UTC Sup cha'DIch offset chel :c:type:`TAI64Time` lo'laHghach.
