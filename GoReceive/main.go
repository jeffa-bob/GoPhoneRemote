package main

import (
	"github.com/webview/webview"
	"github.com/zserge/lorca"
	"net/url"
)

func lorcaMain() {
	// Create UI with data URI
	ui, _ := lorca.New("data:text/html,"+url.PathEscape(startform), "", 410, 300)

	for _, device := range ListDevices() {
		ui.Eval("AddItemToDeviceList(" + device.Model + ", " + device.Serial + ");")
	}
	defer ui.Close()
	<-ui.Done()
}

func main() {
	InitAdb()
	//ListDevices()
	webviewMain()
	// lorcaMain()
}

func webviewMain() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("GoPhoneRemote")
	w.SetSize(310, 200, webview.HintNone)
	w.Navigate("data:text/html," + url.PathEscape(startform))

	w.Bind("ListDevices", ListDevices)

	w.Run()
}
