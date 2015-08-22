package stats

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"

	flag "github.com/ogier/pflag"
)

const programVersion = "0.0.1"

var (
	ignore = flag.BoolP("ignore", "i", false, "ignore invalid numbers")
	behead = flag.BoolP("behead", "b", false,
		"remove the first line (head) from calculations. Useful to ignore column names")
	separator = flag.StringP("separator", "s", " ",
		"define the SEPARATOR to use instead of whitespace for column separator")
	column  = flag.IntP("column", "c", 1, "calculate stats based on the specified COLUMN")
	version = flag.BoolP("version", "v", false, "print version information and exit")
)

func fail(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func calculate(s *Stats) {
	if len(flag.Args()) == 0 {
		parse("<stdin>", os.Stdin, s)
	}

	for _, filename := range flag.Args() {
		if filename == "-" {
			parse("<stdin>", os.Stdin, s)
			continue
		}
		file, err := os.Open(filename)
		if err != nil {
			fail("%s\n", err.Error())
		}
		parse(filename, file, s)
	}
}

func parse(filename string, input *os.File, s *Stats) {
	r := csv.NewReader(input)
	sep, _ := utf8.DecodeRuneInString(*separator)
	r.Comma = sep
	var line int64
	for {
		line += 1
		record, err := r.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			fail("An error occurred while reading the file %s: %+v\n", filename, err)
		}
		if line == 1 && *behead {
			continue
		}
		if *column > len(record) {
			fail("Invalid column number: %d", *column)
		}
		value, err := strconv.ParseFloat(record[*column-1], 64)
		if err != nil {
			if *ignore {
				continue
			} else {
				fail("Invalid number found in file %s at line %d: %s\n", filename, line, record[*column-1])
			}
		}
		s.Update(value)
	}
}

func Tool(desc string, outputter func(s *Stats)) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s\n\n", desc)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nWith no FILE, or when FILE is -, read standard input.")
	}
	flag.Parse()

	if *version {
		fmt.Printf("%s %s\n", os.Args[0], programVersion)
		return
	}

	s := NewStats()
	calculate(s)
	outputter(s)
}
