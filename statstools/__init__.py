"A set of command-line statistics tools"

VERSION = (0, 1, 0)

__author__ = "Jonhnny Weslley"
__contact__ = "jw@jonhnnyweslley.net"
__homepage__ = "http://github.com/jweslley/stats-tools"
__version__ = ".".join(map(str, VERSION))

import fileinput, sys, optparse, numpy

class Tool:

  usage = '[OPTIONS] [FILENAME]...'
  description = ''
  version = __version__
  format = ''

  def __init__(self, function):
    self.function = function

  def parse_options(self):
    """
      Define and parse `optparse` options for command-line usage.
    """
    parser = optparse.OptionParser(
               usage = "%%prog %s" % self.usage,
               description = self.description,
               version = "%%prog %s" % self.version)

    parser.add_option("-c", "--column", type="int", dest="column", metavar="COLUMN",
                       help="calculate stats based on the specified COLUMN")
    parser.add_option("-s", "--separator", dest="separator", metavar="SEPARATOR",
                       help="use SEPARATOR instead of whitespace for column separator")
    parser.add_option("-b", "--behead", dest="behead", default=False, action="store_true", 
                       help="remove the first line (head) from calculations. Useful to ignore column names.")

    (options, sys.argv[1:]) = parser.parse_args()
    self.behead = options.behead
    self.separator = options.separator
    self.column = options.column if options.column else 1
    self.column = self.column if self.column < 0 else self.column - 1

  def load_data(self):
    self.data = []
    for line in fileinput.input():
      if self.behead and fileinput.isfirstline(): continue
      columns = line.split(self.separator) if self.separator else line.split()
      self.data.append(float(columns[self.column]))

  def print_results(self):
    if len(self.data) > 0:
      formatter = self.format or ('\n'.join(map(lambda stat: '%s: %%s' % stat, self.function)) \
                    if isinstance(self.function, list) else '%s')
      print formatter % self.calculate()

  def calculate(self):
    if isinstance(self.function, str):
      return getattr(numpy, self.function)(self.data)
    elif isinstance(self.function, list):
      return tuple([getattr(numpy, stat)(self.data) for stat in self.function])
    elif callable(self.function):
      return self.function(self.data)
    else:
      raise 'Undefined'

  def run(self):
    self.parse_options()
    self.load_data()
    self.print_results()

def run(function, description, format=''):
  tool = Tool(function)
  tool.description = description
  tool.format = format
  tool.run()
