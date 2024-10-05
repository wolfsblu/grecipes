package templates

import "embed"

//go:embed assets
var AssetsFS embed.FS

//go:embed html
var TemplateFS embed.FS
