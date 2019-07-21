package main

import (
	"gioui.org/ui/app"
)

func init() { // HLinitfunc
	go func() { // HLinitfunc
		w := app.NewWindow(nil) // HLeventloop
		for range w.Events() {  // HLeventloop
		} // HLeventloop
	}() // HLinitfunc
} // HLinitfunc

func main() { // HLmain
	app.Main() // HLmain
} // HLmain
