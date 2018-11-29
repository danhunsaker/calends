.. _calendar-systems:

.. index:: calendars

Calendar Systems
================

The systems listed here are the built-in ones. This list is expected to grow
significantly over time, as more and more calendar systems are added. But you
can also add custom systems to your apps without waiting for us to add them
ourselves – instructions for that are in the Custom Calendars section, below.

Throughout these documents, the term ``TAITime`` is used to refer generically to
the ``TAI64NAXURTime``, ``TAI64Time``, or ``TAITime`` type. The exact form of
this name you'll see most often varies by programming language, and is covered
in much more detail in the Custom Calendars section.

.. toctree::
   :maxdepth: 2
   :glob:

   systems/*

.. _custom-calendars:

.. index:: custom calendars

Custom Calendars
================

As with every other aspect of Calends, the custom calendar system support uses
the same basic flow in every language, with minor variations in each to account
for the differences those languages introduce. As with every other aspect of
Calends, though, we've opted to document each language's unique approaches
separately, so you don't have to do any mental conversions yourself.

.. note::
	 Custom calendars are considered an advanced feature, so most users woun't be
	 using anything detailed here. It can be nice to know how these things work
	 under the hood, though, for those interested in that. Select your language,
	 below, and dig right in!

.. toctree::
   :maxdepth: 2

   Golang <go/custom_calendars.rst>
   C/C++ <c/custom_calendars.rst>
   PHP <php/custom_calendars.rst>
