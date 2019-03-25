.. _installation-c:

.. index::
   pair: installation; C/C++

Installing Calends for C/C++
============================

.. _installation-c-binary:

Binary Install
--------------

For use with C/C++, simply grab the latest version of ``libcalends`` from the
`GitHub Releases page <https://github.com/danhunsaker/calends/releases>`_, and
extract its contents wherever your compiler expects to find ``.h`` and
``.so``/``.dll``/``.dylib`` files. Be sure to grab the correct version for your
architecture!

.. _installation-c-source:

Source Install
--------------

To install from source, you'll need Golang 1.9+ installed to use its compiler.
Clone the repository, build ``libcalends``, then copy the resulting
``.so``/``.dll``/``.dylib`` and ``.h`` files to wherever your C/C++ compiler
expects to find them.

.. code-block:: bash

   # Sample Linux steps:
   mkdir -p $GOPATH/src/github.com/danhunsaker
   cd $GOPATH/src/github.com/danhunsaker
   git clone https://github.com/danhunsaker/calends
   cd calends/libcalends
   go get ../...
   go build -v -i -buildmode=c-shared -o libcalends.so

Adjust the above example commands as needed for your actual development OS.
