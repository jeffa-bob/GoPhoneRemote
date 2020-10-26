package main

import (
	"net/url"
    "github.com/zserge/lorca"
    "github.com/webview/webview"
)

func lorcaMain() {
	// Create UI with data URI
	ui, _ := lorca.New("data:text/html,"+url.PathEscape(formhtml), "", 600, 200)
	defer ui.Close()
	<-ui.Done()
}

func main(){
    webviewMain()
    //lorcaMain()
}

func webviewMain(){
    w := webview.New(true)
    defer w.Destroy()
    w.SetSize(310, 200, webview.HintNone)
    w.Navigate("data:text/html,"+url.PathEscape(formhtml))
     
    w.Run()
}