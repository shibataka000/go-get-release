package github

import (
	"bytes"
	"regexp"
	"text/template"
)

var (
	// DefaultPatterns is recommended patterns. This is overwrite for each platform.
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
	a, err := regexp.Compile(asset)
	if err != nil {
		return Pattern{}, err
	}

	b, err := template.New("ExecBinaryName").Parse(execBinary)
	if err != nil {
		return Pattern{}, err
	}

	return newPattern(a, b), nil
}

// match returns true if pattern matches given asset name.
func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.Name))
}

// execute applies a pattern to given asset and returns [ExecBinary] object.
func (p Pattern) execute(asset Asset) (ExecBinary, error) {
	submatch := p.asset.FindSubmatch([]byte(asset.Name))
	data := map[string]string{}
	for _, name := range p.asset.SubexpNames() {
		index := p.asset.SubexpIndex(name)
		if index >= 0 && index < len(submatch) {
			data[name] = string(submatch[index])
		}
	}

	var b bytes.Buffer
	err := p.execBinary.Execute(&b, data)
	if err != nil {
		return ExecBinary{}, err
	}
	return newExecBinary(b.String()), nil
}

// priority returns priority.
// Pattern with bigger priority is prioritized over pattern with smaller priority.
func (p Pattern) priority(asset Asset) int {
	if p.match(asset) {
		return len(p.asset.String())
	}
	return 0
}

// PatternList is a list of [Pattern].
type PatternList []Pattern

// newPatternListFromStringMap returns a new [PatternList] object.
// Keys of map should be regular expressions of GitHub release asset name and compilable by [regexp.Compile].
// Values of map should be templates of executable binary name and parsable by [text/template.Template.Parse].
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

// find [Asset] which matches [Pattern] with biggest priority and returns [Asset] and [Pattern].
func find(assets AssetList, patterns PatternList) (Asset, Pattern, error) {
	var foundAsset Asset
	var foundPattern Pattern
	var priority = 0

	for _, p := range patterns {
		for _, a := range assets {
			if p.match(a) && priority < p.priority(a) {
				foundAsset, foundPattern, priority = a, p, p.priority(a)
			}
		}
	}
	if priority == 0 {
		return Asset{}, Pattern{}, ErrPatternNotMatched
	}
	return foundAsset, foundPattern, nil
}
