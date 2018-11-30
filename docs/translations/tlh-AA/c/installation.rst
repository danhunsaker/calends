.. _installation-c:

.. index::
   pair: installation; C/C++

Calends jom C/C++
============================

binary jom
--------------

Qo'noS lo' C/C++, simply latest version ``libcalends`` grab vo'
github Releases page <https://github.com/danhunsaker/calends/releases>_, 'ej
'a ghIH extract wherever ``.h`` tu' pIH compiler je
``.so``/``.dll`` teywI'. be lugh version grab
architecture!

Hal jom
--------------

jom vo' Hal, golang 1.9 + compiler lo' jom bImejnIS.
repository clone, ``libcalends`` qach, vaj ghot'e' copy
teywI' ``.so``/``.dll``, ``.h`` 'ej wherever pIH C/C++ compiler
chaH vItu'.

.. code-block:: bash

   # vIlopQo' jay' sample linux mIw:
   mkdir -p $GOPATH/src/github.com/danhunsaker
   cd $GOPATH/src/github.com/danhunsaker
   git clone https://github.com/danhunsaker/calends
   cd calends/libcalends
   go get ../...
   go build -v -i -buildmode=c-shared -o libcalends.so

wovbe' example ra' lIS Hoch nIS actual development os.
