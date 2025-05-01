package core

import (
	"fmt"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// Pan moves the canvas based on mouse movement
// This is the core function handling panning operation
func (pxCanvas *PxCanvas) Pan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previousCoord == nil {
		return
	}

	// Calculate movement delta directly from mouse positions
	deltaX := ev.Position.X - previousCoord.Position.X
	deltaY := ev.Position.Y - previousCoord.Position.Y

	fmt.Printf("Panning: deltaX=%f, deltaY=%f\n", deltaX, deltaY)

	// Apply the movement directly to the canvas offset
	pxCanvas.CanvasOffset = pxCanvas.CanvasOffset.Add(fyne.NewPos(deltaX, deltaY))

	// Force layout update and refresh
	if pxCanvas.renderer != nil {
		pxCanvas.renderer.Layout(pxCanvas.Size())
	}
	pxCanvas.Refresh()
}

// Scale changes the pixel size for zooming
// This function handles both zooming in and zooming out
func (pxCanvas *PxCanvas) Scale(direction int) {
	// Get current pixel size
	currentSize := pxCanvas.PxSize
	newSize := currentSize

	// Calculate new size based on zoom direction
	if direction > 0 {
		// Zoom in - larger pixels
		newSize = currentSize + 1
		if newSize > 40 {
			newSize = 40 // Maximum zoom level
		}
	} else if direction < 0 {
		// Zoom out - smaller pixels
		newSize = currentSize - 1
		if newSize < 2 {
			newSize = 2 // Minimum zoom level
		}
	}

	// Only proceed if the size actually changed
	if newSize != currentSize {
		fmt.Printf("Zooming: %d -> %d\n", currentSize, newSize)

		// Store the old dimensions
		oldWidth := float32(pxCanvas.PxCols * currentSize)
		oldHeight := float32(pxCanvas.PxRows * currentSize)

		// Update the pixel size
		pxCanvas.PxSize = newSize

		// Calculate new dimensions
		newWidth := float32(pxCanvas.PxCols * newSize)
		newHeight := float32(pxCanvas.PxRows * newSize)

		// Adjust offset to zoom toward center
		widthDiff := newWidth - oldWidth
		heightDiff := newHeight - oldHeight

		// Move toward center of canvas
		pxCanvas.CanvasOffset = pxCanvas.CanvasOffset.SubtractXY(
			widthDiff/2,
			heightDiff/2,
		)

		// Force layout update and refresh
		if pxCanvas.renderer != nil {
			pxCanvas.renderer.Layout(pxCanvas.Size())
		}
		pxCanvas.Refresh()
	}
}

// PixelToCanvas converts a pixel position to a canvas position
func (pxCanvas *PxCanvas) PixelToCanvas(x, y int) fyne.Position {
	// Convert pixel coordinates to canvas coordinates
	canvasX := float32(x*pxCanvas.PxSize) + pxCanvas.CanvasOffset.X
	canvasY := float32(y*pxCanvas.PxSize) + pxCanvas.CanvasOffset.Y

	return fyne.NewPos(canvasX, canvasY)
}

// Creates a new RGBA image from the canvas, starting at the given bounds
func (pxCanvas *PxCanvas) PixelSelection(topLeft, bottomRight *image.Point) *image.RGBA {
	bounds := image.Rectangle{
		Min: *topLeft,
		Max: *bottomRight,
	}
	img := image.NewRGBA(bounds)
	// Copy of the pixels in the bounds
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if srcImg, ok := pxCanvas.PixelData.(*image.RGBA); ok {
				img.Set(x, y, srcImg.At(x, y))
			}
		}
	}
	return img
}

// Utility function to get mouse events relative to the canvas
func (pxCanvas *PxCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (int, int) {
	bounds := pxCanvas.Bounds()
	return int(ev.Position.X) - bounds.Min.X, int(ev.Position.Y) - bounds.Min.Y
}
