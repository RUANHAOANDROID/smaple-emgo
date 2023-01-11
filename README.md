# smaple-emgo
go 简单的示例包括了：
- webview desktop
- gin web server
- gin static flutter web 
- http client
- udp 
- log
- sqlite gorm

## 构建

因为使用了sqlite3需要开启CGO CGO_ENABLED=1

### 静态文件服务
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
//go:embed manifest.json  
//go:embed version.json 
var Static embed.FS 

