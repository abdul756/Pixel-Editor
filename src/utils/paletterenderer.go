package utils

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PaletteRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Palette
}

func (renderer *PaletteRenderer) MinSize() fyne.Size {
	return fyne.NewSize(30, 30)
}

func (renderer *PaletteRenderer) Layout(size fyne.Size) {
	renderer.objects[0].Resize(size)
}
func (renderer *PaletteRenderer) Refresh() {
	renderer.Layout(fyne.NewSize(30, 30))
	renderer.square.FillColor = renderer.parent.Color
	renderer.square.StrokeWidth = 1
	renderer.square.StrokeColor = color.NRGBA{0, 0, 0, 255} // Black border for all palettes

	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{255, 255, 255, 255} // White for selected
	}

	renderer.objects[0] = &renderer.square
	canvas.Refresh(renderer.parent)
}

func (renderer *PaletteRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}
func (renderer *PaletteRenderer) Destroy() {}
