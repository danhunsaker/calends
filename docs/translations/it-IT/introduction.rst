Introduction
============

As mentioned before, Calends is a library for handling dates and times across
arbitrary calendar systems. But what does that actually mean?

Let's say you're working on an application that uses dates. Pretty much anything
can qualify, really – we use dates in everything, throughout our daily lives.
Scheduling, journaling, historical research, projections into the future, or
just displaying the current date in the UI. Now let's say you want your app to
be used by people all over the globe. The current approach, used for decades, is
to simply use the Gregorian Calendar, which has (partially as a side effect of
this decision) become the default calendar system in use across the globe, for
coordinating, tracking, and preserving events worldwide.

But this decision wasn't made with internationalization and localization in
mind. It was made as a result of practicality, with limited computing
capabilities at the time, and persists mostly as a result of laziness – if the
entire world is already using it anyway, why bother with anything else? It has
also persisted out of ignorance – many people aren't aware there are other
calendars out there in the first place, never mind that several are still in
active use to this day. Properly localizing applications should *include*
adjusting the displayed date to use the preferred calendar system of the user.

Sadly, most of the solutions currently available for handling dates (and times)
in software are purpose-built for a single calendar system, and use APIs
entirely different from those meant to handle dates in others. This makes it
very tricky to build an application that supports more than one calendar system
at a time. Each new calendar system requires hours of work to learn, connect,
and usually abstract to a point where it is usable within the larger
application, and even that's no guarantee the values can be stored or compared
accurately.

That's what Calends set out to solve. It provides a single interface for
interacting with any supported calendar system, and an easy way to extend it to
support others, so that once you've added support for a calendar system once,
you have that support anywhere, without having to rewrite anything to fit your
next application. Additionally, you can take full advantage of any calendar
system implemented by anybody else.

Accept date/time input in any calendar system, perform date/time calculations on
the resulting value, again in any calendar system, and then display the result –
yes, in any calendar system. Dates are stored, internally, in an extremely
accurate value that can track dates out 146 *billion* years into the past or
future, with a resolution of |10^-45| seconds, which is smaller than Planck
Time\ [#fintro1]_. In other words, it should be more than sufficient to record
instants of any duration and resolution desired for any conceived use case.

.. TODO::
   After more calendar systems which are in common use today are implemented,
   talk about how much value they'll add to developers' toolboxes. :)

.. [#fintro1] `Planck Time <https://en.wikipedia.org/wiki/Planck_time>`_ is the
   smallest meaningful unit of time, and is about 54×\ |10^-45| seconds. It
   corresponds to the amount of time it takes a photon (traveling at the speed
   of light through a vaccuum) to traverse one Planck Length, which itself is
   about |10^-20| times the diameter of a proton. Even quantum interactions
   below this scale lose any meaning, and so values below them are considered
   extraneous, in addition to being entirely unmeasurable with current
   technologies and techniques.

.. |10^-45| replace:: 10\ :sup:`-45`
.. |10^-20| replace:: 10\ :sup:`-20`
