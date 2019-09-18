package main

import (
	"gioui.org/ui/app"
)

func main() {
	go func() { // HLmain
		w := app.NewWindow()   // HLeventloop
		for range w.Events() { // HLeventloop
		}
	}() // HLmain
	app.Main() // HLmain
}
