package main

import (
	"github.com/zserge/lorca"
	"net/url"
)

func main() {
	// Create UI with data URI
	ui, _ := lorca.New("data:text/html,"+url.PathEscape(`
 <html>
  <head><title>Hello</title></head>
  <body><h1>Hello, world!</h1></body>
 </html>
 `), "", 600, 200)
	defer ui.Close()
	// Create a GoLang function callable from JS
	ui.Bind("hello", func() string { return "World!" })
	// Call above `hello` function then log to the JS console
	ui.Eval("hello().then( (x) => { console.log(x) })")
	// Wait until UI window is closed
	<-ui.Done()
}
