.. _calendar-system-stardate:

.. index:: Stardate, calendars; Stardate

Stardates
=========

Yes, the ones from :t:`Star Trek`\ ™.

The science fiction universe of :t:`Star Trek`\ ™ introduced a calendar system
which was simultaneously futuristic (for the time) and arbitrary. Over the
decades since its initial use on screen, especially with the growing popularity
of the franchise, the "stardate" has been analyzed and explored and refined into
a number of different variants, each trying to reconcile the arbitrariness of
the on-screen system into something consistent and usable.

This calendar system implementation is an attempt to combine all these variants
into a single system, using the format parameter to select which variant to use.
It was originally ported in 2018 from code by Aaron Chong (2015 version), under
provisions of the MIT License. My thanks for Aaron's use of the MIT License on
the original code, which allowed me to port it cleanly and legally.

Original source: http://rinsanity.weebly.com/files/theme/stardate_public.js

*Calendar Name:*
  ``stardate``

*Supported Input Types:*
  - string
  - integer
  - arbitrary-precision floating point number of seconds

*Supported Format Strings:*
  - ``main``         - **One of the older, more widely-accepted variants.**
    Alternately called the "issue number style" stardate, it's a combined
    TOS/TNG variant, and the one used by Google Calendar. It was originally
    devised by Andrew Main in CE 1994, with revisions made through CE 1997. See
    http://starchive.cs.umanitoba.ca/?stardates/ for the full explanation of
    this variant.
  - ``kennedy``      - In 2006, Richie Kennedy released another combined
    variant, this one designed to have a single continuous count, more like the
    Julian Day Count than Main's issue number system.
  - ``pugh90s``      - Steve Pugh devised 2 separate variants, one of them in
    the 1990s, and the other later on, both focused on the TNG era. They are
    unique in that, for negative stardates, the fractional part increases in the
    opposite direction of the expected one. That is, 15129.999999999 would be
    followed by 15128.000000000 instead of 15129.999999998. The original version
    used an unadjusted Gregorian year as the basis for the duration of a given
    range of stardates, meaning that 0.05 units refer to a larger duration of
    time during a leap year than it would otherwise.
  - ``pughfixed``    - The later of Steve Pugh's systems noted the discrepancy,
    and opted to adjust the year length value to the actual average length of a
    Gregorian year, 365.2425 days. This means 0.05 units are always the same
    duration, but does mean that the Gregorian year doesn't generally start at
    the same point in consecutive stardate ranges.
  - ``schmidt``      - A joint effort between Andreas Schmidt and Graham
    Kennedy, this variant only covers TNG-era stardates, and while it can be
    used proleptically, it ignores the alternate format used prior to TNG. It is
    also virtually identical to ``pugh90s``, but the fractional component
    increases normally for negative stardates.
  - ``guide-equiv``  - One of five variants proposed by TrekGuide.com, this is
    the "out-of-universe equivalent" calculation. It isn't intended to be
    accurate for any use other than personal entertainment.
  - ``guide-tng``    - The second of the five TrekGuide variants, this one is
    the current scale listed for TNG-era stardates, and is show-accurate (or at
    least as close to it as feasible with an entirely arbitrary system). Note,
    however, that it is only accurate for TNG-era dates.
  - ``guide-tos``    - The third variant, then, covers the TOS-era stardates.
    Again, it is only accurate to the TOS era.
  - ``guide-oldtng`` - The fourth variant is no longer displayed on the
    TrekGuide site, and was actually pulled from a previous version of the
    stardates page. It covers the TNG era only, and uses slightly different
    numbers in its calculations than the current approach - specifically, it
    assumes Earth years cover 1000 stardates.
  - ``guide-oldtos`` - Representing the very first set of calculations
    available in archives of the TrekGuide site, the fifth TrekGuide variant
    assumes that 1000 stardates are one Earth year in the TOS era, and
    calculates dates based on that assumption. This variant was replaced within
    seven months of that first archival, after it was noticed that TOS-era
    stardates don't fit a 1000-stardate model.

    .. note::

    	 This calendar system is not yet actually implemented.

  - ``aldrich``      - A proof of concept originally written in C#, this variant
    results in dates very close to those produced by Pugh's and Schmidt's, but
    uses a more simplified calculation to do it.
  - ``red-dragon``   - A system devised by/for the Red Dragon Inn roleplaying
    forum site, it uses a fixed ratio of roughly two and three quarters
    stardates per Earth day. It makes no representations about accuracy outside
    the context of the site itself.
  - ``sto-hynes``    - John Hynes, creator of the Digital Time site, offers a
    calculation for STO\ [#fsd1]_ stardates which appears to be the most accurate
    variant for those interested in generating those. The system doesn't
    represent itself as accurate *outside* the game, but is intentionally
    proleptic.
  - ``sto-academy``  - Based on an online calculator provided by the STO Academy
    game help site, it is only accurate for stardates within the game, and does
    not offer to calculate dates for the rest of the franchise.
  - ``sto-tom``      - Another variant intended only to calculate STO stardates,
    this one was attributed to Major Tom, and hosted as a Wolfram Alpha widget.
  - ``sto-anthodev`` - Another STO variant, hosted on GitHub.

*Offsets:*
  - Must be provided as a string in the format  ``"stardate variant"`` or
    ``"variant stardate"``.

.. [#fsd1] Star Trek™ Online
