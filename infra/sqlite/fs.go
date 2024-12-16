package sqlite

import (
	"embed"
)

//go:embed migrations
var migrationFS embed.FS
