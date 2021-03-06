package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"regexp"

	"github.com/koron/gomigemo/embedict"
	"github.com/koron/gomigemo/migemo"
)

const version = "0.1.0"

const separator = " "

var flag_n = flag.Bool("n", false, "print line number with output lines")
var flag_H = flag.Bool("H", false, "print the filename for each match")

type grepOpt struct {
	optNumber   bool
	optFilename bool
	filename    string
}

func main() {
	st := _main()
	os.Exit(st)
}

func _main() int {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "multigrep v%s\n\nUsage: multigrep [options] pattern [files...]\n", version)
		flag.PrintDefaults()
	}
	var dictPath = flag.String("d", "", "Alternate location to dictionary")

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return 2
	}

	var dict migemo.Dict
	var err error
	if *dictPath == "" {
		dict, err = embedict.Load()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
	} else {
		dict, err = migemo.Load(*dictPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
	}

	res := make( []PolarizedMultiMatcher, 0, 10 )
	patterns := strings.Split(strings.Trim(flag.Arg(0), " "), separator)
	for _, pat := range patterns {

		polar := true
		if pat[0] == '!' {
			polar = false
			pat = pat[1:]
		}

		var re MultiMatcher 
		var err error
		if len(pat) > 2 {
			switch pat[0:2] {
			case "r:": // Regexp
				re, err = regexp.Compile(pat[2:])
			case "i:": // Ignorecase Regexp
				re, err = regexp.Compile("(?i)" + pat[2:])
			case "m:": // Migemo
				re, err = migemo.Compile(dict, pat[2:])
			case "s:": // String Contains
				re = StringMatcher{ str: pat[2:] }
			default:
				re, err = regexp.Compile("(?i)" + pat)
			}
		} else {
			re, err = regexp.Compile("(?i)" + pat)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}

		res = append( res, PolarizedMultiMatcher{ matcher: re, polar: polar } )
	}

	opt := &grepOpt{
		optNumber:   *flag_n,
		optFilename: *flag_H || flag.NArg() > 2,
	}

	total := 0
	// If there's only one arg, then we need to match against the input
	if flag.NArg() == 1 {
		opt.filename = "stdin"

		if total, err = grep(os.Stdin, res, opt); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}

	} else {
		// More than one arg. We must be searching against a file
		for _, arg := range flag.Args()[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 2
			}
			defer f.Close()

			opt.filename = arg
			var count int
			if count, err = grep(f, res, opt); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 2
			}
			total += count
		}
	}

	if total == 0 {
		return 1
	}

	return 0
}
