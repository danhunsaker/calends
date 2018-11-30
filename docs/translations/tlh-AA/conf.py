# -*- coding: utf-8 -*-

import os
execfile(os.path.abspath('../../conf.py'))

# -- Project information -----------------------------------------------------

project = 'Calends'

# Add any paths that contain templates here, relative to this directory.
templates_path = ['../../.templates']

# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
html_static_path = ['../../.static']

# -- Options for HTMLHelp output ---------------------------------------------

# Output file base name for HTML help builder.
htmlhelp_basename = 'Calendsdoc'


# -- Options for LaTeX output ------------------------------------------------

# Grouping the document tree into LaTeX files. List of tuples
# (source start file, target name, title,
#  author, documentclass [howto, manual, or own class]).
latex_documents = [
    (master_doc, 'Calends.tex', 'Calends: Documented',
     'Dan Hunsaker', 'manual'),
]


# -- Options for manual page output ------------------------------------------

# One entry per manual page. List of tuples
# (source start file, name, description, authors, manual section).
man_pages = [
    (master_doc, 'calends', 'Calends: Documented',
     [author], 1)
]


# -- Options for Texinfo output ----------------------------------------------

# Grouping the document tree into Texinfo files. List of tuples
# (source start file, target name, title, author,
#  dir menu entry, description, category)
texinfo_documents = [
    (master_doc, 'Calends', 'Calends: Documented',
     author, 'Calends', 'One line description of project.',
     'Miscellaneous'),
]
