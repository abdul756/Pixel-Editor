package utils

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Palette struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	PaletteIndex int
	clickHandler func(palette *Palette)
}

func (p *Palette) SetColor(c color.Color) {
	p.Color = c
	p.Refresh()
}

func NewPalette(state *State, color color.Color, PaletteIndex int, clickHandler func(palette *Palette)) *Palette {
	palette := &Palette{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		PaletteIndex: PaletteIndex,
	}
	palette.ExtendBaseWidget(palette)

	return palette
}

// Tapped handles mouse click event
func (palette *Palette) Tapped(_ *fyne.PointEvent) {
	palette.clickHandler(palette)
}

func (palette *Palette) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(palette.Color)
	objects := []fyne.CanvasObject{square}

	return &PaletteRenderer{
		square:  *square,
		objects: objects,
		parent:  palette,
	}
}
