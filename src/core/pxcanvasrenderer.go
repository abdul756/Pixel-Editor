package core

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
	cursorRect   *canvas.Rectangle
}

func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	// Create a slice including the canvas image, border lines, and cursor rect
	objects := make([]fyne.CanvasObject, 0, 6)

	// Add canvas image
	objects = append(objects, renderer.canvasImage)

	// Add 4 border lines
	for i := 0; i < 4; i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}

	// Add cursor rect
	objects = append(objects, renderer.cursorRect)

	return objects
}

func (renderer *PxCanvasRenderer) Destroy() {}

func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	// Position and size the canvas image based on pixel size and offset
	renderer.canvasImage.Move(renderer.pxCanvas.CanvasOffset)
	renderer.canvasImage.Resize(fyne.NewSize(
		float32(renderer.pxCanvas.PxCols*renderer.pxCanvas.PxSize),
		float32(renderer.pxCanvas.PxRows*renderer.pxCanvas.PxSize),
	))

	// Update border to surround the canvas
	renderer.updateBorder()

	// Update the cursor position
	renderer.updateCursor()
}

func (renderer *PxCanvasRenderer) Refresh() {
	// Update image if needed
	if renderer.pxCanvas.reloadImage {
		renderer.canvasImage.Image = renderer.pxCanvas.PixelData
		renderer.canvasImage.Refresh()
		renderer.pxCanvas.reloadImage = false
	}

	// Update border position
	renderer.updateBorder()

	// Update cursor position
	renderer.updateCursor()

	// Refresh all elements
	renderer.canvasImage.Refresh()
	for i := 0; i < 4; i++ {
		renderer.canvasBorder[i].Refresh()
	}
	renderer.cursorRect.Refresh()
}

// updateBorder sets the positions of the border lines
func (renderer *PxCanvasRenderer) updateBorder() {
	offset := renderer.pxCanvas.CanvasOffset
	width := float32(renderer.pxCanvas.PxCols * renderer.pxCanvas.PxSize)
	height := float32(renderer.pxCanvas.PxRows * renderer.pxCanvas.PxSize)

	// Top line
	renderer.canvasBorder[0].Position1 = fyne.NewPos(offset.X, offset.Y)
	renderer.canvasBorder[0].Position2 = fyne.NewPos(offset.X+width, offset.Y)

	// Right line
	renderer.canvasBorder[1].Position1 = fyne.NewPos(offset.X+width, offset.Y)
	renderer.canvasBorder[1].Position2 = fyne.NewPos(offset.X+width, offset.Y+height)

	// Bottom line
	renderer.canvasBorder[2].Position1 = fyne.NewPos(offset.X, offset.Y+height)
	renderer.canvasBorder[2].Position2 = fyne.NewPos(offset.X+width, offset.Y+height)

	// Left line
	renderer.canvasBorder[3].Position1 = fyne.NewPos(offset.X, offset.Y)
	renderer.canvasBorder[3].Position2 = fyne.NewPos(offset.X, offset.Y+height)
}

// updateCursor positions the cursor highlight over the correct pixel
func (renderer *PxCanvasRenderer) updateCursor() {
	// Get cursor position
	cursorPos := renderer.pxCanvas.cursorPos

	// Hide cursor if outside canvas
	if cursorPos.X < 0 || cursorPos.Y < 0 {
		renderer.cursorRect.Hide()
		return
	}

	// Convert to pixel coordinates
	pixelX, pixelY := renderer.pxCanvas.CanvasToPixel(cursorPos)

	// Check if pixel is within bounds
	if pixelX < 0 || pixelY < 0 || pixelX >= renderer.pxCanvas.PxCols || pixelY >= renderer.pxCanvas.PxRows {
		renderer.cursorRect.Hide()
		return
	}

	// Calculate position in canvas coordinates
	pixelSize := float32(renderer.pxCanvas.PxSize)
	xPos := renderer.pxCanvas.CanvasOffset.X + float32(pixelX)*pixelSize
	yPos := renderer.pxCanvas.CanvasOffset.Y + float32(pixelY)*pixelSize

	// Position and size the cursor rectangle
	renderer.cursorRect.Move(fyne.NewPos(xPos, yPos))
	renderer.cursorRect.Resize(fyne.NewSize(pixelSize, pixelSize))

	// Set cursor appearance
	renderer.cursorRect.StrokeWidth = 2
	renderer.cursorRect.StrokeColor = color.NRGBA{255, 0, 0, 255} // Red border
	renderer.cursorRect.FillColor = color.NRGBA{255, 255, 0, 100} // Semi-transparent yellow

	// Make cursor visible
	renderer.cursorRect.Show()
}

// UpdateCursor is called when the mouse moves
func (renderer *PxCanvasRenderer) UpdateCursor(ev *desktop.MouseEvent) {
	if ev == nil {
		// Use stored cursor position if event is nil
		renderer.updateCursor()
		return
	}

	// Update with the new cursor position
	renderer.pxCanvas.cursorPos = ev.Position
	renderer.updateCursor()
}

// newPxCanvasRenderer creates a new canvas renderer for a px canvas
func newPxCanvasRenderer(pxCanvas *PxCanvas) *PxCanvasRenderer {
	// Create a more visible cursor with a contrasting stroke and fill
	cursorRect := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 0, A: 100})
	cursorRect.StrokeWidth = 3
	cursorRect.StrokeColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255} // Bright red stroke
	cursorRect.FillColor = color.NRGBA{R: 255, G: 255, B: 0, A: 100} // Semi-transparent yellow fill

	// Hide cursor initially
	cursorRect.Hide()

	// Create the border lines
	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < 4; i++ {
		canvasBorder[i].StrokeColor = color.RGBA{50, 50, 50, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvas.NewImageFromImage(pxCanvas.PixelData),
		canvasBorder: canvasBorder,
		cursorRect:   cursorRect,
	}

	renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
	renderer.canvasImage.FillMode = canvas.ImageFillContain

	return renderer
}
