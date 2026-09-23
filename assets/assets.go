package assets

import "embed"

//go:embed *
var assets embed.FS

func FS() embed.FS {
	return assets
}
