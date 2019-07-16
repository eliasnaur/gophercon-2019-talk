package main

import (
	"image"
	"image/color"

	"gioui.org/ui"
	"gioui.org/ui/app"
	"gioui.org/ui/draw"
	"gioui.org/ui/f32"
	"gioui.org/ui/input"
	"gioui.org/ui/pointer"
)

// START QUEUE OMIT
func init() {
	go func() {
		w := app.NewWindow(nil)
		button := new(Button)
		ops := new(ui.Ops) // HLops
		for e := range w.Events() {
			if _, ok := e.(app.DrawEvent); ok {
				ops.Reset()
				queue := w.Queue() // HLqueue
				button.Layout(queue, ops)
				w.Draw(ops)
			}
		}
	}()
}

// END QUEUE OMIT

func main() {
	app.Main()
}

type Button struct {
	pressed bool
}

// START OMIT
func (b *Button) Layout(queue input.Queue, ops *ui.Ops) {
	for _, e := range queue.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}

	col := color.RGBA{A: 0xff, R: 0xff}
	if b.pressed {
		col = color.RGBA{A: 0xff, G: 0xff}
	}
	pointer.RectAreaOp{ // HLevent
		Size: image.Point{X: 500, Y: 500}, // HLevent
	}.Add(ops) // HLevent
	pointer.HandlerOp{Key: b}.Add(ops) // HLevent
	drawSquare(ops, col)
}

// END OMIT

func drawSquare(ops *ui.Ops, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{X: 500, Y: 500},
	}
	draw.ColorOp{Color: color}.Add(ops)
	draw.DrawOp{Rect: square}.Add(ops)
}
