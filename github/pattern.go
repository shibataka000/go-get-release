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

func newPattern(asset *regexp.Regexp, execBinary *template.Template) Pattern {
	return Pattern{
		asset:      asset,
		execBinary: execBinary,
	}
}

func newPatternFromString(asset string, execBinary string) (Pattern, error) {
	re, err := regexp.Compile(asset)
	if err != nil {
		return Pattern{}, err
	}

	tmpl, err := template.New("ExecBinary").Parse(execBinary)
	if err != nil {
		return Pattern{}, err
	}

	return newPattern(re, tmpl), nil
}

func mustNewPatternFromString(asset string, execBinary string) Pattern {
	p, err := newPatternFromString(asset, execBinary)
	if err != nil {
		panic(err)
	}
	return p
}

func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.name))
}

func (p Pattern) renderExecBinary(asset Asset) (ExecBinary, error) {
	return ExecBinary{}, nil
}

type PatternList []Pattern

func newPatternList(assets []string, execBinaries []string) (PatternList, error) {
	if len(assets) != len(execBinaries) {
		return nil, errors.New("")
	}

	patterns := PatternList{}

	for i := range assets {
		p, err := newPatternFromString(assets[i], execBinaries[i])
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
