package github

import (
	"errors"
	"regexp"
	"slices"
	"text/template"
)

type Pattern struct {
	asset      *regexp.Regexp
	execBinary *template.Template
}

func newPattern(assetPattern string, execBinaryPattern string) (Pattern, error) {
	asset, err := regexp.Compile(assetPattern)
	if err != nil {
		return Pattern{}, err
	}

	execBinary, err := template.New("ExecBinary").Parse(execBinaryPattern)
	if err != nil {
		return Pattern{}, err
	}

	return Pattern{
		asset:      asset,
		execBinary: execBinary,
	}, nil
}

func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.name))
}

func (p Pattern) renderExecBinary(asset Asset) (ExecBinary, error) {
	return ExecBinary{}, nil
}

type PatternList []Pattern

func newPatternList(assetPatterns []string, execBinaryPatterns []string) (PatternList, error) {
	if len(assetPatterns) != len(execBinaryPatterns) {
		return nil, errors.New("")
	}

	patterns := PatternList{}

	for i := range assetPatterns {
		p, err := newPattern(assetPatterns[i], execBinaryPatterns[i])
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, p)
	}

	return patterns, nil
}

func (pl PatternList) find(asset Asset) (Pattern, error) {
	index := slices.IndexFunc(pl, func(p Pattern) bool {
		return p.match(asset)
	})
	if index == -1 {
		return Pattern{}, errors.New("")
	}
	return pl[index], nil
}
