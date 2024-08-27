package github

import (
	"fmt"
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

func newPatternListFromStringSlice(assets, execBinary []string) (PatternList, error) {
	if len(assets) != len(execBinary) {
		return nil, fmt.Errorf("number of asset patterns and exec binary patterns are not same")
	}

	patterns := PatternList{}
	for i := range assets {
		pattern, err := newPatternFromString(assets[i], execBinary[i])
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, pattern)
	}
	return patterns, nil
}
