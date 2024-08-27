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

func newPatternFromString(rawAsset, rawExecBinary string) (Pattern, error) {
	asset, err := regexp.Compile(rawAsset)
	if err != nil {
		return Pattern{}, err
	}
	execBinary, err := regexp.Compile(rawExecBinary)
	if err != nil {
		return Pattern{}, err
	}
	return newPattern(asset, execBinary), nil
}

func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.name()))
}

type PatternList []Pattern

func newPatternListFromStringSlice(rawAssets, rawExecBinary []string) (PatternList, error) {
	if len(rawAssets) != len(rawExecBinary) {
		return nil, nil
	}

	patterns := PatternList{}
	for i := range rawAssets {
		pattern, err := newPatternFromString(rawAssets[i], rawExecBinary[i])
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, pattern)
	}
	return patterns, nil
}
