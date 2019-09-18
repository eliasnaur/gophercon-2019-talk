package main

import (
	"image/color"
	"math"
	"time"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/f32"
	"gioui.org/ui/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		ops := new(ui.Ops)
		for e := range w.Events() {
			if e, ok := e.(app.UpdateEvent); ok {
				ops.Reset()

				// START OMIT
				square := f32.Rectangle{Max: f32.Point{X: 500, Y: 500}}
				radius := animateRadius(e.Config.Now(), 250)

				// Position
				ui.TransformOp{}.Offset(f32.Point{ // HLdraw
					X: 100, // HLdraw
					Y: 100, // HLdraw
				}).Add(ops) // HLdraw
				// Color
				paint.ColorOp{Color: color.RGBA{A: 0xff, G: 0xcc}}.Add(ops) // HLdraw
				// Clip corners
				roundRect(ops, 500, 500, radius, radius, radius, radius) // HLdraw
				// Draw
				paint.PaintOp{Rect: square}.Add(ops) // HLdraw
				// Animate
				ui.InvalidateOp{}.Add(ops) // HLdraw

				// Submit operations to the window.
				w.Update(ops) // HLdraw
				// END OMIT
			}
		}
	}()
	app.Main()
}

// START RR OMIT
// https://pomax.github.io/bezierinfo/#circles_cubic.
func roundRect(ops *ui.Ops, width, height, se, sw, nw, ne float32) {
	w, h := float32(width), float32(height)
	const c = 0.55228475 // 4*(sqrt(2)-1)/3
	var b paint.PathBuilder
	b.Init(ops)
	b.Move(f32.Point{X: w, Y: h - se})
	b.Cube(f32.Point{X: 0, Y: se * c}, f32.Point{X: -se + se*c, Y: se}, f32.Point{X: -se, Y: se})
	b.Line(f32.Point{X: sw - w + se, Y: 0})
	b.Cube(f32.Point{X: -sw * c, Y: 0}, f32.Point{X: -sw, Y: -sw + sw*c}, f32.Point{X: -sw, Y: -sw})
	b.Line(f32.Point{X: 0, Y: nw - h + sw})
	b.Cube(f32.Point{X: 0, Y: -nw * c}, f32.Point{X: nw - nw*c, Y: -nw}, f32.Point{X: nw, Y: -nw})
	b.Line(f32.Point{X: w - ne - nw, Y: 0})
	b.Cube(f32.Point{X: ne * c, Y: 0}, f32.Point{X: ne, Y: ne - ne*c}, f32.Point{X: ne, Y: ne})
	b.End()
}

// END RR OMIT

var start = time.Now()

func animateRadius(t time.Time, max float32) float32 {
	dt := t.Sub(start).Seconds()
	radius := math.Abs(math.Sin(dt))
	return float32(radius) * max
}
