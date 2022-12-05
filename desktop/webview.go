package desktop

import (
	"emcs-relay-go/configs"
	"github.com/jchv/go-webview2"
	"log"
)

func Run() {
	web2()
}

// linux win7
func web1() {
	//w := webview.New(true)
	//Web := w
	//defer Web.Destroy()
	//w.SetTitle("设备中继器")
	//w.SetSize(1600, 900, webview.HintNone)
	//w.Navigate(configs.IndexHtml)
	//w.Run()
}

// win10 win11
func web2() {
	w := webview2.New(true)
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetSize(1400, 900, webview2.HintFixed)
	w.Navigate(configs.IndexHtml)
	w.Run()
}
