package main

import (
	"image/color"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/draw"
	"gioui.org/ui/f32"
)

func init() {
	go func() {
		w := app.NewWindow(nil)
		// START OMIT
		ops := new(ui.Ops)
		for e := range w.Events() {
			if e, ok := e.(app.DrawEvent); ok {
				ops.Reset()

				square := f32.Rectangle{
					Min: f32.Point{X: 50, Y: 50},
					Max: f32.Point{X: 500, Y: 500},
				}
				draw.ColorOp{Color: color.RGBA{A: 0xff, G: 0xff}}.Add(ops)
				draw.DrawOp{Rect: square}.Add(ops)
				ui.InvalidateOp{}.Add(ops)

				w.Draw(ops)
			}
		}
		// END OMIT
	}()
}

func main() {
	app.Main()
}
