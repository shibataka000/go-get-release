package github

import (
	"regexp"
)

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

func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.name()))
}

type PatternList []Pattern

func newPatternListFromStringArray(assets, execBinaries []string) (PatternList, error) {
	if len(assets) != len(execBinaries) {
		return nil, newInvalidPatternError()
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
