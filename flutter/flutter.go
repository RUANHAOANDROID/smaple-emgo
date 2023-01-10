package flutter

import "embed"

//go:embed assets/*
//go:embed canvaskit/*
//go:embed icons/*
//go:embed .last_build_id
//go:embed favicon.png
//go:embed flutter_service_worker.js
//go:embed index.html
//go:embed main.dart.js
//go:embed flutter.js
//go:embed manifest.json
//go:embed version.json
var Static embed.FS
