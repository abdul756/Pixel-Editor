package core

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"github.com/abdul756/Pixel-Editor/src/utils"
)

// Default background color - light gray
var DefaultBackgroundColor = color.NRGBA{128, 128, 128, 255}

// Brushable interface implementation
func (pxCanvas *PxCanvas) SetColor(c color.Color, x, y int) {
	if rgba, ok := pxCanvas.PixelData.(*image.RGBA); ok {
		if x >= 0 && y >= 0 && x < pxCanvas.PxCols && y < pxCanvas.PxRows {
			rgba.Set(x, y, c)
			pxCanvas.Refresh()
		}
	}
}

// PaintSwatch paints at the given position using the current brush
func (pxCanvas *PxCanvas) PaintSwatch(p fyne.Position) {
	// Convert from canvas coordinates to pixel coordinates
	pixelX, pixelY := pxCanvas.CanvasToPixel(p)

	// Apply brushing based on brush type
	switch pxCanvas.appState.BrushType {
	case utils.BrushTypePen:
		pxCanvas.paintPixel(pixelX, pixelY)
	case utils.BrushTypeEraser:
		pxCanvas.paintEraser(pixelX, pixelY)
	}
}

// CanvasToPixel converts from canvas position to pixel position
// This is a critical function for cursor position accuracy
func (pxCanvas *PxCanvas) CanvasToPixel(p fyne.Position) (int, int) {
	// Calculate the position relative to the canvas offset
	relativeX := p.X - pxCanvas.CanvasOffset.X
	relativeY := p.Y - pxCanvas.CanvasOffset.Y

	// Convert to pixel coordinates, ensuring we use integer division
	// to get the exact pixel index
	pixelX := int(relativeX) / pxCanvas.PxSize
	pixelY := int(relativeY) / pxCanvas.PxSize

	return pixelX, pixelY
}

// paintPixel paints a single pixel at the given position
func (pxCanvas *PxCanvas) paintPixel(x, y int) {
	// Check if the pixel is within bounds
	if x >= 0 && y >= 0 && x < pxCanvas.PxCols && y < pxCanvas.PxRows {
		pxCanvas.SetColor(pxCanvas.appState.BrushColor, x, y)
		pxCanvas.reloadImage = true
	}
}

// paintEraser erases at the given position (resets to default background color)
func (pxCanvas *PxCanvas) paintEraser(x, y int) {
	if x >= 0 && y >= 0 && x < pxCanvas.PxCols && y < pxCanvas.PxRows {
		pxCanvas.SetColor(DefaultBackgroundColor, x, y)
		pxCanvas.reloadImage = true
	}
}

// Helper function for absolute value (used by mouse dragging)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
