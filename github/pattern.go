package github

import (
	"regexp"
)

type Pattern struct {
	asset      *regexp.Regexp
	execBinary *regexp.Regexp
}

// newPattern returns a new pattern object.
func newPattern(asset, execBinary *regexp.Regexp) Pattern {
	return Pattern{
		asset:      asset,
		execBinary: execBinary,
	}
}

// newPatternFromString returns a new pattern object.
func newPatternFromString(asset, execBinary string) (Pattern, error) {
	assetRegexp, err := regexp.Compile(asset)
	if err != nil {
		return Pattern{}, err
	}
	execBinaryRegexp, err := regexp.Compile(execBinary)
	if err != nil {
		return Pattern{}, err
	}
	return newPattern(assetRegexp, execBinaryRegexp), nil
}

// match returns true if pattern matches given asset.
func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.name()))
}

// PatternList is a list of patterns.
type PatternList []Pattern

// newPatternListFromStringArray returns a new list of patterns.
// The number of asset patterns and the number of exec binary patterns in arguments must be same.
func newPatternListFromStringArray(assets, execBinaries []string) (PatternList, error) {
	if len(assets) != len(execBinaries) {
		return nil, ErrUnpairablePattern
	}

	patterns := PatternList{}
	for i := range assets {
		pattern, err := newPatternFromString(assets[i], execBinaries[i])
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, pattern)
	}
	return patterns, nil
}
