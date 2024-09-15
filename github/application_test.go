package github

// func TestApplicationServiceFindAsset(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		repoFullName  string
// 		tag           string
// 		assetPatterns []string
// 		asset         Asset
// 	}{}

// 	ctx := context.Background()
// 	app := NewApplicationService(
// 		NewAssetRepository(ctx, githubToken),
// 		NewExecBinaryRepository(),
// 	)

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			require := require.New(t)
// 			asset, err := app.FindAsset(ctx, tt.repoFullName, tt.tag, tt.assetPatterns)
// 			require.NoError(err)
// 			require.Equal(tt.asset, asset)
// 		})
// 	}
// }

// func TestApplicationServiceFindExecBinary(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		asset              Asset
// 		assetPatterns      []string
// 		execBinaryPatterns []string
// 		execBinary         ExecBinary
// 	}{}

// 	ctx := context.Background()
// 	app := NewApplicationService(
// 		NewAssetRepository(ctx, githubToken),
// 		NewExecBinaryRepository(),
// 	)

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			require := require.New(t)
// 			execBinary, err := app.FindExecBinary(tt.asset, tt.assetPatterns, tt.execBinaryPatterns)
// 			require.NoError(err)
// 			require.Equal(tt.execBinary, execBinary)
// 		})
// 	}
// }

// func TestApplicationServiceInstall(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		asset              Asset
// 		assetPatterns      []string
// 		execBinaryPatterns []string
// 		execBinary         ExecBinary
// 	}{}

// 	ctx := context.Background()
// 	app := NewApplicationService(
// 		NewAssetRepository(ctx, githubToken),
// 		NewExecBinaryRepository(),
// 	)

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			require := require.New(t)
// 			execBinary, err := app.FindExecBinary(tt.asset, tt.assetPatterns, tt.execBinaryPatterns)
// 			require.NoError(err)
// 			require.Equal(tt.execBinary, execBinary)
// 		})
// 	}
// }
