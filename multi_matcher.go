package main

import (
	"strings"
 )

type MultiMatcher interface {
	MatchString(testee string) bool
}

type PolarizedMultiMatcher struct {
	matcher MultiMatcher
	polar bool
}

type StringMatcher struct {
	str string
}

func (str StringMatcher) MatchString(testee string) bool {
	return strings.Contains( testee, str.str )
}

