package github

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// Release in GitHub repository
type Release interface {
	Tag() string
	Version() string
	Assets() ([]Asset, error)
	AssetByName(name string) (Asset, error)
	AssetByPlatform(goos string, goarch string) (Asset, error)
}

type release struct {
	client *client
	repo   *repository
	tag    string
	id     int64
}

// Tag return tag name of GitHub release
func (r *release) Tag() string {
	return r.tag
}

// Version return version string, which does not have 'v' prefix
func (r *release) Version() string {
	return strings.TrimLeft(r.Tag(), "v")
}

// Assets return assets in release
func (r *release) Assets() ([]Asset, error) {
	return newAssetLister(r).List()
}

// AssetByName return asset which have specific name
func (r *release) AssetByName(name string) (Asset, error) {
	assets, err := r.Assets()
	if err != nil {
		return nil, err
	}
	for _, a := range assets {
		if a.Name() == name {
			return a, nil
		}
	}
	return nil, fmt.Errorf("no asset found: %s", name)
}

// AssetByPlatform return asset which have specific goos and goarch
func (r *release) AssetByPlatform(goos, goarch string) (Asset, error) {
	if r.hasPredefinedAssetName(goos, goarch) {
		assetName, err := r.predefinedAssetName(goos, goarch)
		if err != nil {
			return nil, err
		}
		return r.AssetByName(assetName)
	}

	assets, err := r.Assets()
	if err != nil {
		return nil, err
	}
	assets, err = filterAssetsByPlatform(assets, goos, goarch)
	if err != nil {
		return nil, err
	}

	switch len(assets) {
	case 0:
		return nil, fmt.Errorf("no asset found")
	case 1:
		return assets[0], nil
	case 2:
		if assets[0].IsExecBinary() && (assets[1].IsArchived() || assets[1].IsCompressed()) {
			return assets[0], nil
		} else if assets[1].IsExecBinary() && (assets[0].IsArchived() || assets[0].IsCompressed()) {
			return assets[1], nil
		}
		fallthrough
	default:
		assetNames := []string{}
		for _, a := range assets {
			assetNames = append(assetNames, a.Name())
		}
		return nil, fmt.Errorf("too many assets found: %v", strings.Join(assetNames, ", "))
	}
}

// hasPredefinedAssetName return true if asset name is predefined in assetNameMap
func (r *release) hasPredefinedAssetName(goos, goarch string) bool {
	_, err := r.predefinedAssetName(goos, goarch)
	return err == nil
}

// predefinedAssetName return predefined asset name in assetNameMap
func (r *release) predefinedAssetName(goos, goarch string) (string, error) {
	key := fmt.Sprintf("%s/%s", r.repo.Owner(), r.repo.Name())
	tmplMap, ok := assetNameMap[key]
	if !ok {
		return "", fmt.Errorf("asset name is not predefined: %s", key)
	}

	key = fmt.Sprintf("%s/%s", goos, goarch)
	tmplStr := ""
	if s, ok := tmplMap["default"]; ok {
		tmplStr = s
	}
	if s, ok := tmplMap[key]; ok {
		tmplStr = s
	}
	if tmplStr == "" {
		return "", fmt.Errorf("unsupported GOOS and GOARCH: %s", key)
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("assetName").Parse(tmplStr)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(buf, struct {
		Tag     string
		Version string
		Goos    string
		Goarch  string
	}{
		Tag:     r.tag,
		Version: r.Version(),
		Goos:    goos,
		Goarch:  goarch,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// filterAssetsByPlatform return assets which have specific goos and goarch
func filterAssetsByPlatform(assets []Asset, goos, goarch string) ([]Asset, error) {
	var result []Asset
	for _, a := range assets {
		if !a.ContainReleaseBinary() {
			continue
		}
		assetGoos, err := a.Goos()
		if err != nil {
			continue
		}
		assetGoarch, err := a.Goarch()
		if err != nil {
			continue
		}
		if assetGoos == goos && assetGoarch == goarch {
			result = append(result, a)
		}
	}
	return result, nil
}
