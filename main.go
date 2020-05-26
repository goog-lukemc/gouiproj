package main

import (
	"github.com/goog-lukemc/gouidom"
	"github.com/goog-lukemc/gouielement"
)

var unloadMsg chan int
var done chan struct{}

func main() {
	done = make(chan struct{}, 1)
	unloadMsg = make(chan int, 1)

	go unload()

	app, err := gouidom.NewApp("Wellness")
	if err != nil {
		gouidom.CLog("%v", err)
		unloadMsg <- 1 // Unloads WASM with a exit log
	}

	comp := gouielement.NewElementLib(app)
	comp.Readable()

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
