package desktop

import (
	"emcs-relay-go/configs"
	"github.com/webview/webview"
)

func Run() {
	w := webview.New(true)
	Web := w
	defer Web.Destroy()
	w.SetTitle("设备中继器")
	w.SetSize(1600, 900, webview.HintNone)
	//indexHtml, err := os.ReadFile("static/index.html")
	//if err != nil {
	//	logger.Log.Debug(err.Error())
	//}
	//html := string(indexHtml)
	//w.SetHtml(html)
	w.Navigate(configs.IndexHtml)
	//w.Navigate("https://hao88.cloud")
	w.Run()
}
