package github

import (
	"bytes"
	"regexp"
	"slices"
	"strconv"
	"text/template"
)

var (
	// DefaultPatterns is recommended patterns. This is overwritten for each platform.
	DefaultPatterns = map[string]string{}
)

// Pattern represents a pair of regular expression of GitHub release asset name and template of executable binary name.
// This is used to select an appropriate one from GitHub release assets and determine an executable binary name.
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

// match returns true if regular expression of GitHub release asset name matches given GitHub release asset name.
func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.name()))
}

// priority returns a literal prefix length of regular expression of GitHub release asset name as priority of pattern.
// Pattern with higher priority is prioritized over pattern with lower priority.
func (p Pattern) priority() int {
	prefix, _ := p.asset.LiteralPrefix()
	return len(prefix)
}

// execute applies a template of executable binary name to values of capturing group in regular expression of GitHub release asset name and returns [ExecBinary] object.
func (p Pattern) execute(asset Asset) (ExecBinary, error) {
	data := map[string]string{}
	submatch := p.asset.FindStringSubmatch(asset.name())

	for i := range submatch {
		data[strconv.Itoa(i)] = submatch[i]
	}

	for _, name := range p.asset.SubexpNames() {
		index := p.asset.SubexpIndex(name)
		if index >= 0 && index < len(submatch) {
			data[name] = submatch[index]
		}
	}

	var b bytes.Buffer
	if err := p.execBinary.Execute(&b, data); err != nil {
		return ExecBinary{}, err
	}

	return newExecBinary(b.String()), nil
}

// newPatternArrayFromStringMap returns a new array of [Pattern] objects.
// Map's keys should be regular expressions of GitHub release asset name and values should be templates of executable binary name.
func newPatternArrayFromStringMap(patterns map[string]string) ([]Pattern, error) {
	ps := []Pattern{}
	for asset, execBinary := range patterns {
		p, err := newPatternFromString(asset, execBinary)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}

// find [Asset] and [Pattern] which match and returns them.
// Pattern with higher priority is prioritized over pattern with lower priority.
func find(assets []Asset, patterns []Pattern) (Asset, Pattern, error) {
	cloned := slices.Clone(patterns)
	slices.SortFunc(cloned, func(p1, p2 Pattern) int {
		return p2.priority() - p1.priority()
	})

	for _, p := range cloned {
		for _, a := range assets {
			if p.match(a) {
				return a, p, nil
			}
		}
	}

	return Asset{}, Pattern{}, ErrNoPatternAndAssetMatched
}
