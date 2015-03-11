package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "regexp"
)

var out io.Writer

func init() {
	out = os.Stdout
}

// Does the grepping
func grep(r io.Reader, pmms []PolarizedMultiMatcher, opt *grepOpt) (int, error) {
	buf := bufio.NewReader(r)
	n := 1
	count := 0
	for {
		b, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return 0, err
		}
		line    := string(b)
		matched := true
		for _, pmm := range pmms {
			if pmm.matcher.MatchString(line) != pmm.polar {
				matched = false
				break
			}
		}
		if matched {
			count++
			if opt.optFilename {
				fmt.Fprintf(out, "%s:", opt.filename)
			}
			if opt.optNumber {
				fmt.Fprintf(out, "%d:", n)
			}
			fmt.Fprintln(out, line)
		}
		n++
	}
	return count, nil
}

