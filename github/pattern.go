package github

import (
	"bytes"
	"regexp"
	"text/template"
)

var (
	DefaultPatterns = map[string]string{}
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
// asset should be a regular expression of GitHub release asset name and compilable by [regexp.Compile].
// execBinary should be a template of executable binary name and parsable by [text/template.Template.Parse].
func newPatternFromString(asset string, execBinary string) (Pattern, error) {
	re, err := regexp.Compile(asset)
	if err != nil {
		return Pattern{}, err
	}

	tmpl, err := template.New(execBinary).Parse(execBinary)
	if err != nil {
		return Pattern{}, err
	}

	return newPattern(re, tmpl), nil
}

// match returns true if pattern matches given asset name.
func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.Name))
}

func (p Pattern) apply(asset Asset) (ExecBinary, error) {
	var b bytes.Buffer
	submatch := p.asset.FindStringSubmatch(asset.Name)
	err := p.execBinary.Execute(&b, submatch)
	if err != nil {
		return ExecBinary{}, err
	}
	return newExecBinary(b.String()), nil
}

// PatternList is a list of [Pattern].
type PatternList []Pattern

// newPatternListFromStringArray returns a new [PatternList] object.
// See [newPatternFromString] for more details.
// The number of assets and the number of execBinaries should be same.
func newPatternListFromStringMap(patterns map[string]string) (PatternList, error) {
	pl := PatternList{}
	for asset, execBinary := range patterns {
		p, err := newPatternFromString(asset, execBinary)
		if err != nil {
			return nil, err
		}
		pl = append(pl, p)
	}
	return pl, nil
}

func find(assets AssetList, patterns PatternList) (Asset, Pattern, error) {
	for _, p := range patterns {
		for _, a := range assets {
			if p.match(a) {
				return a, p, nil
			}
		}
	}
	return Asset{}, Pattern{}, ErrPatternNotMatched
}
