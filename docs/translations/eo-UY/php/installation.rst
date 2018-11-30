.. _installation-php:

.. index::
   pair: installation; PHP

Installing Calends for PHP
==========================

.. TODO ::
   Make this actually work...

   Binary Install
   --------------

   For use with PHP, simply grab the latest version of ``calends-php`` from the
   `GitHub Releases page <https://github.com/danhunsaker/calends/releases>`_,
   and put it in your extension directory. Be sure to grab the correct version
   for your PHP version, OS, and processor!

   Source Install
   --------------

To install from source, you'll need a few prerequisites:

- Go 1.9+
- PHP 7.1+ source code (frequently something like ``php-dev``)
- `Zephir/C <https://github.com/phalcon/zephir/>`_ and its dependencies

Once those are installed, clone the repository, build ``libcalends``, then run
``zephir install`` from the ``php`` subdirectory.

.. code-block:: bash

   # Sample Linux steps:
   mkdir -p $GOPATH/src/github.com/danhunsaker
   cd $GOPATH/src/github.com/danhunsaker
   git clone https://github.com/danhunsaker/calends
   cd calends/libcalends
   go get ../...
   go build -v -i -buildmode=c-shared -o libcalends.so
   cd php
   zephir install

Adjust the above example commands as needed for your actual development OS.

Finally, just add ``extension=calends.so`` to your ``php.ini``, and restart
PHP-FPM and/or your web server.
