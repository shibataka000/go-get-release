package github

import "regexp"

type Pattern struct {
	asset      *regexp.Regexp
	execBinary *regexp.Regexp
}

func newPattern(asset, execBinary *regexp.Regexp) Pattern {
	return Pattern{
		asset:      asset,
		execBinary: execBinary,
	}
}

func newPatternFromString(rawPattern string) (Pattern, error) {
	return Pattern{}, nil
}

func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.name()))
}

type PatternList []Pattern

func newPatternListFromStringSlice(rawPatterns []string) (PatternList, error) {
	patterns := PatternList{}
	for _, rawPattern := range rawPatterns {
		pattern, err := newPatternFromString(rawPattern)
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, pattern)
	}
	return patterns, nil
}
