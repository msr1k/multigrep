package main

import (
	"strings"
 )

type MultiMatcher interface {
	MatchString(pattern string) bool
}

type PolarizedMultiMatcher struct {
	matcher MultiMatcher
	polar bool
}

type StringMatcher struct {
	str string
}

func (str StringMatcher) MatchString(pattern string) bool {
	return strings.Contains( str.str, pattern )
}

