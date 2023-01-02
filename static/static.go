package static

import (
	"embed"
)

//go:embed index.html manifest.json version.json .last_build_id flutter.js flutter_service_worker.js main.dart.js assets/* canvaskit/* icons/*
var FlutterWeb embed.FS
