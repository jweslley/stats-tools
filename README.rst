stats-tools - A set of command-line statistics tools
====================================================

.. contents::

Installation
------------

Dependencies
````````````
 * numpy
 * scipy

On Debian systems, you may install them by running::

  sudo apt-get install python-numpy python-scipy

Installing stats-tools
``````````````````````

Install ``stats-tools`` by using ``easy_install``::

  sudo easy_install stats-tools

Or checkout from git repository::

  git clone git://github.com/jweslley/stats-tools.git
  cd stats-tools
  sudo python setup.py install


Utilities
---------

 * **min** - Calculate the minimum of a number sequence
 * **max** - Calculate the maximum of a number sequence
 * **mean** - Calculate the mean of a number sequence
 * **median** - Calculate the median of a number sequence
 * **std** - Calculate the standard deviation of a number sequence
 * **var** - Calculate the variance of a number sequence
 * **sum** - Calculate the sum of a number sequence
 * **stats** - Output a summary table including mean, median, mininum, maximum, standard deviation and variance of a number sequence
 * **summary** - Output a summary table including minimum, lower quartile, median, upper quartile, maximum of a number sequence
 * **fivenum** - Calculate Tukey's five number summary (minimum, lower-hinge, median, upper-hinge, maximum) of a number sequence based on 1.5 times the interquartile distance


Usage
-----

All utilities take as input a file in table format to perform some calculation based on it. A tipical input file is shown below::

  1 2 4
  3 5 4
  6 4 6
  4 5 6
  9 12 16

Considering this input file, let's call it ``example1.dat``, you can calculate some statistics like:

The ``max`` value on the first column::

  max example1.dat

The ``min`` value on the second column::

  min -c 2 example1.dat

You still can use negative column numbers to start counting from the right. Thus, the ``sum`` of the values on last column::

  sum -c -1 example1.dat

If the input file's columns are separated by another character instead of whitespace characters (space, tab, newline, return, formfeed), like CSV files, you can use the ``-s`` option to denote this. The next example outputs a statistical ``summary`` about the second column of the following file (``example2.dat``)::

  "A",10,12
  "A",11,14
  "B",5,8
  "B",6,10
  "A",10.5,13
  "B",7,11

Calculating the summary::

  summary -c 2 -s , example2.dat

Commonly, data files may contain a head, i.e., the first line describes the columns, something like the ``example3.dat`` file showed below::

  Year,Make,Model,Description,Price
  1997,Ford,E350,"ac abs moon",3000.00
  1999,Chevy,"Venture ""Extended Edition""","",4900.00
  1999,Chevy,"Venture ""Extended Edition, Very Large""","",5000.00
  1996,Jeep,Grand Cherokee,"MUST SELL!air, moon roof, loaded",4799.00

The ``-b`` option remove the first line from calculations. In this case, the mean price of the cars is given by::

  mean -b -s, -c-1 test/example3.dat


Piping data
```````````

All ``stats-tools`` read data from standard input if no file is passed to them. The following command calculates the max value on the second column containing the word ``bar`` in the file ``foo.dat``::

  grep bar foo.dat | max -c 2


Bugs and Feedback
-----------------

If you discover any bugs or have some idea, feel free to create an issue on GitHub:
  
`<http://github.com/jweslley/stats-tools/issues>`_


License
-------

MIT license. Copyright (c) 2011 Jonhnny Weslley <http://jonhnnyweslley.net>

See the LICENSE file provided with the source distribution for full details.
