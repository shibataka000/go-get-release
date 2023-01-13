package embed

import _ "embed"

//go:embed files/assets.yaml
var embedAssetData []byte
