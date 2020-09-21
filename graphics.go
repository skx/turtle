package main

import (
	"image"
	"image/color"
	"math"
)

type Graphics struct {
	Image       *image.RGBA
	Pos         Position
	Orientation float64
	Color       color.Color
	Pen         bool
}

type Position struct {
	X, Y float64
}

func NewGraphics(i *image.RGBA, starting Position) (t *Graphics) {
	t = &Graphics{
		Image:       i,
		Pos:         starting,
		Orientation: 0.0,
		Color:       color.Black,
		Pen:         true,
	}

	return
}

func (t *Graphics) Move(x float64, y float64) {
	t.Pos.X = x
	t.Pos.Y = y
}

func (t *Graphics) Forward(dist float64) {
	for i := 0; i < int(dist); i++ {
		if t.Pen {
			t.Image.Set(int(t.Pos.X), int(t.Pos.Y), t.Color)
		}

		x := 1.0 * math.Sin(t.Orientation)
		y := 1.0 * -math.Cos(t.Orientation)

		t.Pos = Position{t.Pos.X + x, t.Pos.Y + y}
	}
}

func (t *Graphics) Turn(radians float64) {
	t.Orientation += radians
}

func (t *Graphics) Direction(radians float64) {
	t.Orientation = radians
}

func (t *Graphics) PenUp() {
	t.Pen = false
}

func (t *Graphics) PenDown() {
	t.Pen = true
}
