.. _features:

Features in Calends
===================

For a current indication of which of these features are fully implemented at the moment, check `the README <https://github.com/danhunsaker/calends>`_.

* Large range and high precision
    Calends understands dates |2^62| seconds into the future or past, in units
    as small as |10^-45| seconds – that's over 146 billion years into the past
    or future (146 138 512 313 years, 169 days, 10 hours, 5 minutes, and 28
    seconds from ``CE 1970 Jan 01 00:00:00 TAI Gregorian``), at resolutions
    smaller than Planck Time (54×\ |10^-45| seconds, and the smallest meaningful
    duration even on the quantum scale). That encompasses well beyond the
    expected lifespan of the Universe, at resolutions enough to represent
    quantum events.

* Supports date (and time) values in multiple calendar systems
    Supported out of the box are the following (all systems are proleptic –
    extrapolated beyond the officially-defined limits – unless specified
    otherwise):

    \

    * :ref:`Unix time <calendar-system-unix>`
        A count of the number of seconds since ``CE 1970 Jan 01 00:00:00 UTC
        Gregorian``

    * :ref:`TAI64 <calendar-system-tai64>`
        Essentially Unix time plus |2^62|, but using TAI seconds instead of UTC
        seconds, so times can be converted unambiguously (UTC uses leap seconds
        to keep the solar zenith at noon, while TAI is a simple, unadjusted
        count). Calends supports an extended version of this spec, with three
        more components, to encode out to 45 places instead of just 18; this is
        also actually the internal time scale used by Calends itself, which is
        how it can support such a broad range of dates at such a high
        resolution.

        \

        * Automatic calculation of leap second offsets
        * Estimation of undefined past and future leap second insertions
        * Automatic updates for handling leap second insertions

    * :ref:`Gregorian <calendar-system-gregorian>`
        The current international standard calendar system

        \

        * Disconnect from native :go:type:`time.Time` implementation, and its
          limitations

    * :ref:`Julian <calendar-system-julian>`
        The previous version of the Gregorian calendar

    * :ref:`Julian Day Count <calendar-system-jdc>`
        A count of days since ``BCE 4713 Jan 01 12:00:00 UTC Julian
        (proleptic)``

    * :ref:`Hebrew <calendar-system-hebrew>`
        \

    * :ref:`Persian <calendar-system-persian>`
        \

    * :ref:`Chinese <calendar-system-chinese>`
        Several variants

    * :ref:`Meso-American <calendar-system-meso-american>`
        Commonly called Mayan, but used by several cultures in the region

    * :ref:`Discordian <calendar-system-discordian>`
        \

    * :ref:`Stardate <calendar-system-stardate>`
        Yes, the ones from :t:`Star Trek`\ ™; several variants exist

    * :ref:`Dynamic <calendar-system-dynamic>`
        User-defined calendars, which don't require writing any code

* Encodes both time spans and instants in a single interface
    The library treats the time values it encodes as ``[start, end)`` sets (that
    is, the ``start`` point is included in the range, as is every point between
    ``start`` and ``end``, but the ``end`` point itself is _not_ included in the
    range). This allows ``duration`` to accurately be ``end - start`` in all
    cases. (And yes, that also means you can create spans with ``duration <
    0``.)

* Supports calculations and comparisons on spans and instants
    Addition, subtraction, intersection, combination, gap calculation, overlap
    detection, and similar operations are all supported directly on Calends
    values.

* Conversion to/from native date/time types
    While this is possible by using a string representation as an intermediary,
    in either direction, some data and precision is lost in such a conversion.
    Instead, Calends supports conversion to and from such types directly,
    preserving as much data and accuracy as each native type provides.

* Geo-temporally aware
    The library provides methods for passing a location instead of a calendar
    system, and selecting an appropriate calendar based on which was most common
    in that location at that point in time. *(Some guess work is involved in
    this process when parsing dates, so it is still preferred to supply the
    calendar system, if known, when parsing.)*

* Time zone support
    \

* Well-defined interfaces for extending the library
    Add more calendar systems, type conversions, or geo-temporal relationships
    without forking/modifying the library itself.

* Shared library (``.so``/``.dll``/``.dylib``)
    In order to use the library outside of Golang projects, we first need to
    export its functionality in a shared library, which can then be accessed
    from other programming evironments and applications, generally via FFI.

* WebAssembly binary
    In order to use the library in the browser, we first need to export its
    functionality in a WebAssembly (WASM) binary, which can then be accessed
    by JavaScript. (Go currently doesn't support the WASI standard, so the
    functions are registered into the global namespace rather than being
    ``export``\ ed by WebAssembly itself. More on that in the JS docs.)

.. |10^-45| replace:: 10\ :sup:`-45`
.. |10^-20| replace:: 10\ :sup:`-20`
.. |2^62| replace:: 2\ :sup:`62`
