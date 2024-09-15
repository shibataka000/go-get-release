package github

import (
	"errors"
	"regexp"
	"slices"
	"text/template"
)

// Pattern represents a pair of regular expression of GitHub release asset name and template of executable binary name.
type Pattern struct {
	asset      *regexp.Regexp
	execBinary *template.Template
}

// newPattern returns a new [Pattern] object.
func newPattern(asset *regexp.Regexp, execBinary *template.Template) Pattern {
	return Pattern{
		asset:      asset,
		execBinary: execBinary,
	}
}

// newPatternFromString returns a new [Pattern] object.
// asset must be a regular expression of GitHub release asset name and compilable by [regexp.Compile].
// execBinary must be a template of executable binary name and parsable by [text/template.Template.Parse].
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

// mustNewPatternFromString is like [newPatternFromString] but panics if arguments cannot be parsed.
func mustNewPatternFromString(asset string, execBinary string) Pattern {
	p, err := newPatternFromString(asset, execBinary)
	if err != nil {
		panic(err)
	}
	return p
}

// match returns true if pattern matches given asset name.
func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.name))
}

func (p Pattern) apply(asset Asset) (ExecBinary, error) {
	return ExecBinary{}, nil
}

// PatternList is a list of [Pattern].
type PatternList []Pattern

// newPatternListFromStringArray returns a new [PatternList] object.
// See [newPatternFromString] for more details.
// The number of assets and the number of execBinaries must be same.
func newPatternListFromStringArray(assets []string, execBinaries []string) (PatternList, error) {
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

// find returns pattern which match given asset first.
func (pl PatternList) find(asset Asset) (Pattern, error) {
	index := slices.IndexFunc(pl, func(p Pattern) bool {
		return p.match(asset)
	})
	if index == -1 {
		return Pattern{}, errors.New("")
	}
	return pl[index], nil
}
