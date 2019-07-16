package main

import (
	"gioui.org/ui/app"
)

func init() { // HLinitfunc
	go func() { // HLinitfunc
		w := app.NewWindow(nil)     // HLeventloop
		for e := range w.Events() { // HLeventloop
			if _, ok := e.(app.DrawEvent); ok { // HLeventloop
				w.Draw(nil) // HLeventloop
			} // HLeventloop
		} // HLeventloop
	}() // HLinitfunc
} // HLinitfunc

func main() { // HLmain
	app.Main() // HLmain
} // HLmain
