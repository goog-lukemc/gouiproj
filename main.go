package main

import (
	"github.com/goog-lukemc/gouidom"
	"github.com/goog-lukemc/gouielement"
	"github.com/goog-lukemc/gouiscreen"
)

var unloadMsg chan int
var done chan struct{}

func main() {
	done = make(chan struct{}, 1)
	unloadMsg = make(chan int, 1)

	go unload()

	app, err := gouidom.NewApp("Reader")
	if err != nil {
		gouidom.CLog("%v", err)
		unloadMsg <- 1 // Unloads WASM with a exit log
	}

	comp := gouielement.NewElementLib(app)

	gouiscreen.DocReader("html/body", comp, nil)

	if app.GetCurrentPath() == "/app/edit" {
		comp.CodeBlock("html/body", app.GetHTMLDocument())
		comp.CodeBlock("html/body", app.GenStyleTemplate())
		comp.CodeBlock("html/body", app.GetAppStyle())
	}

	<-done
}

func unload() {
	for code := range unloadMsg {
		if code == 0 {
			done <- struct{}{}
			gouidom.CLog("Unloading with Exit code:%d", code)
		}
	}
}

// Reader create a readable content
// func Reader(content *gouielement.ReadingData, comp *gouielement.ElementLib, parent string) error {
// 	if content == nil {
// 		content = getDefaultContent(parent, comp)
// 	}
// 	comp.Readable("html/body", content)
// 	return nil
// }

// func getDefaultContent(parent string, comp *gouielement.ElementLib) *gouielement.ReadingData {
// 	c := &gouielement.ReadingData{
// 		Title:    "Welcome to the Go UI framework using Webassembly",
// 		Subtitle: "Calm down!, this project is a WIP",
// 	}

// 	t1 := "Welcome to the goui framework. Have fun!"
// 	c.Content = append(c.Content, comp.IMG(parent, map[string]string{"src": "/image/go-logo_blue.svg", "width": "40", "height": "30"}))
// 	c.Content = append(c.Content, comp.Span(parent, t1))

// 	return c
// }
