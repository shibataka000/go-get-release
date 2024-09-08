package github

import "regexp"

// Pattern is a base of [AssetPattern] and [ExecBinaryPattern].
type Pattern struct {
	re *regexp.Regexp
}

// newPattern returns a new [Pattern] object.
func newPattern(re *regexp.Regexp) Pattern {
	return Pattern{
		re: re,
	}
}

// compilePattern compiles expr as regular expression and returns a new [Pattern] object.
func compilePattern(expr string) (Pattern, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return Pattern{}, err
	}
	return newPattern(re), nil
}
