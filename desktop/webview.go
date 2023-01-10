package desktop

func Run() {
	//web1()
	web2()
}

// linux win7
//func web1() {
//	w := webview.New(true)
//	Web := w
//	defer Web.Destroy()
//	w.SetTitle("设备中继器")
//	w.SetSize(1600, 900, webview.HintNone)
//	w.Navigate(configs.IndexHtml)
//	w.Run()
//}

// win10 win11 lib:"github.com/jchv/go-webview2"
func web2() {
	//ipv4, err := utils.GetOutboundIP()
	//if err != nil {
	//	panic("未获取到本机IP地址")
	//}
	//w := webview2.New(false)
	////w.SetHtml()
	//if w == nil {
	//	log.Fatalln("Failed to load webview.")
	//}
	//w.SetSize(1400, 900, webview2.HintNone) //先指定大小
	//w.SetTitle("gateflow")
	////w.SetSize(1400, 900, webview2.HintMin) //在固定最小大小
	//defer w.Destroy()
	////w.SetSize(1400, 900, webview2.HintFixed)
	//url := "http://" + ipv4.String() + ":8888/gateflow/"
	//utils.Log.Info(url)
	//w.Navigate(url)
	//w.Run()
}
