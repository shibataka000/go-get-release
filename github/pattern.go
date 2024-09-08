package github

import "regexp"

type Pattern struct {
	re *regexp.Regexp
}

func newPattern(re *regexp.Regexp) Pattern {
	return Pattern{
		re: re,
	}
}

func compilePattern(expr string) (Pattern, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return Pattern{}, err
	}
	return newPattern(re), nil
}
