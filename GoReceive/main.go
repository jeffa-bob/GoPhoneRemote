package main

import (
	"github.com/zserge/lorca"

	"net/url"
)

func lorcaMain() {
	// Create UI with data URI
	ui, _ := lorca.New("data:text/html,"+url.PathEscape(startform), "", 410, 300)
	defer ui.Close()
	<-ui.Done()
}

func main() {
	//webviewMain()
	InitAdb()
	lorcaMain()
}

/*
func webviewMain(){
    w := webview.New(true)
    defer w.Destroy()
    w.SetSize(310, 200, webview.HintNone)
    w.Navigate("data:text/html,"+url.PathEscape(startform))

    w.Run()
}*/
