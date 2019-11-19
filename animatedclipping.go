package main

import (
	"image/color"
	"math"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		ops := new(op.Ops)
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				ops.Reset()

				// START OMIT
				square := f32.Rectangle{Max: f32.Point{X: 500, Y: 500}}
				radius := animateRadius(e.Config.Now(), 250)

				// Position
				op.TransformOp{}.Offset(f32.Point{ // HLdraw
					X: 100, // HLdraw
					Y: 100, // HLdraw
				}).Add(ops) // HLdraw
				// Color
				paint.ColorOp{Color: color.RGBA{A: 0xff, G: 0xcc}}.Add(ops) // HLdraw
				// Clip corners
				clip.Rect{Rect: square,
					NE: radius, NW: radius, SE: radius, SW: radius}.Op(ops).Add(ops) // HLdraw
				// Draw
				paint.PaintOp{Rect: square}.Add(ops) // HLdraw
				// Animate
				op.InvalidateOp{}.Add(ops) // HLdraw

				// Submit operations to the window.
				e.Frame(ops) // HLdraw
				// END OMIT
			}
		}
	}()
	app.Main()
}

// END RR OMIT

var start = time.Now()

func animateRadius(t time.Time, max float32) float32 {
	dt := t.Sub(start).Seconds()
	radius := math.Abs(math.Sin(dt))
	return float32(radius) * max
}
