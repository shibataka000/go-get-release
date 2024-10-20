package github

import (
	"bytes"
	"regexp"
	"strconv"
	"text/template"
)

var (
	// DefaultPatterns is recommended patterns. This is overwritten for each platform.
	DefaultPatterns = map[string]string{}
)

// Pattern represents a pair of regular expression of GitHub release asset name and template of executable binary name.
// This is used to select the appropriate one from GitHub release assets and determine an executable binary name.
type Pattern struct {
	// asset is a regular expression of GitHub release asset name.
	// This is used to select the appropriate one from GitHub release assets and used as input data to determine an executable binary name.
	asset *regexp.Regexp
	// execBinary is a template of executable binary name.
	// This is used to determine an executable binary name.
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
func newPatternFromString(asset string, execBinary string) (Pattern, error) {
	a, err := regexp.Compile(asset)
	if err != nil {
		return Pattern{}, err
	}

	b, err := template.New("").Parse(execBinary)
	if err != nil {
		return Pattern{}, err
	}

	return newPattern(a, b), nil
}

// match returns true if regular expression of GitHub release asset name in pattern matches given GitHub release asset name.
func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.Name))
}

// priority returns a literal prefix length of regular expression of GitHub release asset name as priority of pattern.
// Pattern with bigger priority is prioritized over pattern with smaller priority.
func (p Pattern) priority() int {
	prefix, _ := p.asset.LiteralPrefix()
	return len(prefix)
}

// execute applies a template of executable binary name to values of capturing group in regular expression of GitHub release asset name and returns [ExecBinary] object.
func (p Pattern) execute(asset Asset) (ExecBinary, error) {
	data := map[string]string{}
	submatch := p.asset.FindStringSubmatch(asset.Name)

	// Construct data from capturing group.
	for i := range submatch {
		data[strconv.Itoa(i)] = submatch[i]
	}

	// Construct data from named capturing group.
	for _, name := range p.asset.SubexpNames() {
		index := p.asset.SubexpIndex(name)
		if index >= 0 && index < len(submatch) {
			data[name] = submatch[index]
		}
	}

	// Apply a template to data.
	var b bytes.Buffer
	if err := p.execBinary.Execute(&b, data); err != nil {
		return ExecBinary{}, err
	}

	return newExecBinary(b.String()), nil
}

// PatternList is a list of [Pattern].
type PatternList []Pattern

// newPatternListFromStringMap returns a new [PatternList] object.
// Keys of map should be regular expressions of GitHub release asset name.
// Values of map should be templates of executable binary name.
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
			if p.match(a) && priority < p.priority() {
				foundAsset, foundPattern, priority = a, p, p.priority()
			}
		}
	}
	if priority == 0 {
		return Asset{}, Pattern{}, ErrPatternNotMatched
	}
	return foundAsset, foundPattern, nil
}
