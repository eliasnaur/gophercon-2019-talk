package main

import (
	"image/color"
	"math"
	"time"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/draw"
	"gioui.org/ui/f32"
)

func init() {
	go func() {
		w := app.NewWindow(nil)
		// START OMIT
		ops := new(ui.Ops) // HLops
		for e := range w.Events() {
			if e, ok := e.(app.DrawEvent); ok {
				ops.Reset() // HLops

				color := animateColor(e.Config.Now())
				square := f32.Rectangle{
					Min: f32.Point{X: 50, Y: 50},
					Max: f32.Point{X: 500, Y: 500},
				}
				// Add draw operations.
				draw.ColorOp{Color: color}.Add(ops) // HLops
				draw.DrawOp{Rect: square}.Add(ops)  // HLops
				// Request immediate redraw.
				ui.InvalidateOp{}.Add(ops) // HLops

				// Submit operations.
				w.Draw(ops) // HLops
			}
		}
		// END OMIT
	}()
}

func main() {
	app.Main()
}

var start = time.Now()

func animateColor(t time.Time) color.RGBA {
	dt := t.Sub(start).Seconds()
	green := math.Abs(math.Sin(dt))
	return color.RGBA{A: 0xff, G: byte(green * 0xff)}
}
