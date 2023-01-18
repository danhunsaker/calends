.. _installation-php:

.. index::
   pair: installation; PHP

Installing Calends for PHP
==========================

For use with PHP, use Composer to install the PHP FFI wrapper:

.. code-block:: bash

   composer install danhunsaker/calends

The post-install script will grab the appropriate ``libcalends`` for your
system, along with the relevant header file. From there, simply update
your ``php.ini`` to load the FFI extension (if not already loaded) and
preload the header file:

.. code-block:: ini

   extension=ffi.so
   ffi.preload=/path/to/your/code/vendor/lib/calends-phpffi.h

If you don't have access to edit your ``php.ini``, ensure the FFI
extension is available and enabled, then manually load the header file in
your code:

.. code-block:: php

   FFI::load(__DIR__ . "/vendor/lib/calends-phpffi.h");
