package mux

import "regexp"

// regexpCache stores compiled regular expressions to avoid repeated compilation.
var regexpCache = make(map[string]*regexp.Regexp)

func getOrCompileRegexp(pattern string) (*regexp.Regexp, error) {
	if re, ok := regexpCache[pattern]; ok {
		return re, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	regexpCache[pattern] = re
	return re, nil
}
