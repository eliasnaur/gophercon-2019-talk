package main

import (
	"image"
	"image/color"
	"math"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		// START OMIT
		ops := new(op.Ops) // HLops
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				ops.Reset() // HLops

				color := animateColor(e.Now)
				square := image.Rectangle{
					Min: image.Point{X: 50, Y: 50},
					Max: image.Point{X: 500, Y: 500},
				}
				// Add draw operations.
				paint.ColorOp{Color: color}.Add(ops) // HLops
				clip.Rect(square).Add(ops)           // HLops
				paint.PaintOp{}.Add(ops)             // HLops
				// Request immediate redraw.
				op.InvalidateOp{}.Add(ops) // HLops

				// Submit operations.
				e.Frame(ops) // HLops
			}
		}
		// END OMIT
	}()
	app.Main()
}

var start = time.Now()

func animateColor(t time.Time) color.RGBA {
	dt := t.Sub(start).Seconds()
	green := math.Abs(math.Sin(dt))
	return color.RGBA{A: 0xff, G: byte(green * 0xff)}
}
