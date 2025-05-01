package utils

import (
	"image/color"

	"fyne.io/fyne/v2"
)

// BrushType defines the type of brush being used
type BrushType = int

// Brush Type constants
const (
	BrushTypePen    BrushType = iota // Pen brush for drawing individual pixels
	BrushTypeEraser                  // Eraser brush
)

// Brushable is an interface for objects that can be painted on
type Brushable interface {
	SetColor(c color.Color, x, y int)
	PaintSwatch(p fyne.Position)
}

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

// State contains the active UI information
type State struct {
	BrushColor      color.Color
	BrushType       int
	PaletteSelected int // Used in apptype.go
	SwatchSelected  int // Used in state.go
	FilePath        string
	BrushSize       int
	AreaSize        int // Size of the painted area
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
