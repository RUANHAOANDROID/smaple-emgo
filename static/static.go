package static

import "embed"

//go:embed index.html static/fonts/* static/js/*.js static/img/* static/css/*.css
var Static embed.FS
