package application

import (
	"embed"
)

//go:embed templates/*.html
var EmbedFS embed.FS
