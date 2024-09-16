package github

import (
	"bytes"
	"regexp"
	"text/template"
)

var (
	defaultPatterns = PatternList{
		mustNewPatternFromString("trivy_0.53.0_Linux-64bit.tar.gz", "trivy"),
		mustNewPatternFromString("argocd-linux-amd64", "argocd"),
		mustNewPatternFromString("kubectl-argo-rollouts-linux-amd64", "kubectl-argo-rollouts"),
		mustNewPatternFromString("argo-linux-amd64.gz", "argo"),
		mustNewPatternFromString("pack-v0.34.2-linux.tgz", "pack"),
		mustNewPatternFromString("gh_2.52.0_linux_amd64.tar.gz", "gh"),
		mustNewPatternFromString("buildx-v0.15.1.linux-amd64", "docker-buildx"),
		mustNewPatternFromString("docker-compose-linux-x86_64", "docker-compose"),
		mustNewPatternFromString("sops-v3.9.0.linux.amd64", "sops"),
		mustNewPatternFromString("dockle_0.4.14_Linux-64bit.tar.gz", "dockle"),
		mustNewPatternFromString("istioctl-1.22.2-linux-amd64.tar.gz", "istioctl"),
		mustNewPatternFromString("yq_linux_amd64.tar.gz", "yq"),
		mustNewPatternFromString("conftest_0.53.0_Linux_x86_64.tar.gz", "conftest"),
		mustNewPatternFromString("gator-v3.16.3-linux-amd64.tar.gz", "gator"),
		mustNewPatternFromString("opa_linux_amd64", "opa"),
		mustNewPatternFromString("protoc-27.2-linux-x86_64.zip", "protoc"),
		mustNewPatternFromString("snyk-linux", "snyk"),
		mustNewPatternFromString("starship-x86_64-unknown-linux-gnu.tar.gz", "starship"),
	}
)

func DefaultAssetPatterns() []string {
	patterns := []string{}
	for _, p := range defaultPatterns {
		patterns = append(patterns, p.asset.String())
	}
	return patterns
}

func DefaultExecBinaryPatterns() []string {
	patterns := []string{}
	for _, p := range defaultPatterns {
		patterns = append(patterns, p.execBinary.Name())
	}
	return patterns
}

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

	tmpl, err := template.New(execBinary).Parse(execBinary)
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
	var b bytes.Buffer
	submatch := p.asset.FindStringSubmatch(asset.name)
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
// The number of assets and the number of execBinaries must be same.
func newPatternListFromStringArray(assets []string, execBinaries []string) (PatternList, error) {
	if len(assets) != len(execBinaries) {
		return nil, ErrUnpairablePattern
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
