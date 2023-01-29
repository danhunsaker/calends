.. _installation-js:

.. index::
   pair: installation; JS
   pair: installation; WASM

Installing Calends for JS/WASM
==============================

For use with JS, use ``npm`` (or your preferred package manager)::

   .. code-block:: bash

      npm install -s calends

This will pull in the JS wrapper package as well as the corresponding WASM
binary.

For use on the server, that's pretty much it. The library takes care of the
rest.

For use on the web, you'll need to ensure the WASM is accessible to the server,
next to the library itself. The easiest way to ensure this is to pull in
``calends.js`` directly via ``<script>`` tag, but if you use a package to
compile/minify/etc your JS dependencies, you'll need to configure that package
to include ``calends.wasm`` alongside your script(s). Here's an example for
``webpack``::

   .. code-block:: javascript

      const CopyPlugin = require("copy-webpack-plugin");

      // ...

      module.exports = {
         // ...
         plugins: [
            new CopyPlugin({
               patterns: [
                  { from: "node_modules/calends/calends.wasm",
                     to: "[name][ext]" },
               ],
            }),
         ],
      };

It's not clean, but until Go compiles compliant WASM binaries, it's the best we
can do right now, since we can't use `WebAssembly ESM Integration
<https://github.com/WebAssembly/esm-integration/tree/main/proposals/esm-integration>`
yet. Once it defines ``export``\ s correctly, we can drop much of the JS wrapper
and focus purely on translating bare functions into full classes exclusively.
