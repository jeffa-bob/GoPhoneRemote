package main

import (
	"github.com/webview/webview"
	"github.com/zserge/lorca"
	"net/url"
)

var port int32;

func lorcaMain() {
	// Create UI with data URI
	ui, _ := lorca.New("data:text/html,"+url.PathEscape(startform), "", 410, 300)

	for _, device := range ListDevices() {
		ui.Eval("AddItemToDeviceList(" + device.Model + ", " + device.Serial + ");")
	}
	defer ui.Close()
	ui.Bind("ListDevices", ListDevices)
	ui.Bind("StartServer", StartServer)
	ui.Bind("TapSpot", TapSpot)
	ui.Bind("GetTemp", GetTemp)
	<-ui.Done()
}

func main() {
	InitAdb()
	//ListDevices()
	webviewMain()
	// lorcaMain()
}


func webviewMain() {
	ui := webview.New(true)
	defer ui.Destroy()
	ui.SetTitle("GoPhoneRemote")
	ui.SetSize(310, 200, webview.HintNone)
	ui.Navigate("data:text/html," + url.PathEscape(startform))

	ui.Bind("ListDevices", ListDevices)
	ui.Bind("StartServer", StartServer)
	ui.Bind("TapSpot", TapSpot)
	ui.Bind("GetTemp", GetTemp)

	ui.Run()
}

func StartServer(serial string, selport int32) {
	port = selport
	GenerateCode()
	SelectDevice(serial)
	//TapSpot(2560/2, 1600/2)
}
