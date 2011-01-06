#!/usr/bin/env python

import statstools as stats
from glob import glob as ls
from distutils.core import setup

setup(
   name =          'stats-tools',
   version =       stats.__version__,
   description =   stats.__doc__,
   #long_description=open('README.rst').read(),
   author =        stats.__author__,
   author_email =  stats.__contact__,
   url =           stats.__homepage__,
   download_url =  '%s/tarball/v%s' % (stats.__homepage__, stats.__version__),
   license =       'MIT',
   packages =      ['statstools'],
   scripts =       ls('bin/*'),
   requires =      ['numpy', 'scipy'],
   platforms =     ['Linux', 'Unix', 'Mac OS X', 'Windows'],
   classifiers =   ['Development Status :: 4 - Beta',
                    'License :: OSI Approved :: MIT License',
                    'Intended Audience :: Science/Research',
                    'Topic :: Scientific/Engineering',
                    'Environment :: Console',
                    'Programming Language :: Python',
                    'Operating System :: MacOS',
                    'Operating System :: Microsoft :: Windows',
                    'Operating System :: POSIX',
                    'Operating System :: Unix'
                   ]
)
