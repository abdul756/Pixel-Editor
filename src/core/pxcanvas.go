package core

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/abdul756/Pixel-Editor/src/utils"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
	isDragging    bool
}

type PxCanvas struct {
	widget.BaseWidget
	utils.PxCanvasConfig
	renderer    *PxCanvasRenderer
	mouseState  PxCanvasMouseState
	PixelData   image.Image
	appState    *utils.State
	reloadImage bool
	cursorPos   fyne.Position // Track cursor position for highlighting

	// Draggable state
	dragging    bool
	dragStarted bool
	dragPos     fyne.Position
}

// Ensure PxCanvas implements fyne.Draggable and apptype.Brushable
var _ fyne.Draggable = (*PxCanvas)(nil)
var _ utils.Brushable = (*PxCanvas)(nil)

// Ensure PxCanvas implements the desktop mouse interfaces
var _ desktop.Mouseable = (*PxCanvas)(nil)
var _ desktop.Hoverable = (*PxCanvas)(nil)

// No other interface declarations to avoid conflicts with our custom mouse handling

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	x1 := int(pxCanvas.PxCols*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(pxCanvas.PxRows*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	return float32(bounds.Min.X) <= pos.X && pos.X < float32(bounds.Max.X) &&
		float32(bounds.Min.Y) <= pos.Y && pos.Y < float32(bounds.Max.Y)
}

func NewBlankImage(cols, rows int, c color.Color) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, cols, rows))
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			img.Set(j, i, c)
		}
	}
	return img
}

func NewPxCanvas(state *utils.State, config utils.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
		mouseState:     PxCanvasMouseState{},
	}
	// Set the initial canvas color to gray
	pxCanvas.PixelData = NewBlankImage(config.PxCols, config.PxRows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain
	canvasImage.Show()

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < 4; i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{0, 0, 0, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	// Initialize cursor rectangle with default values
	cursorRect := &canvas.Rectangle{
		StrokeColor: color.NRGBA{255, 0, 0, 255},   // Bright red border
		StrokeWidth: 4,                             // Extra thick border
		FillColor:   color.NRGBA{255, 255, 0, 150}, // More opaque yellow fill
	}
	// Hide it initially until mouse moves over canvas
	cursorRect.Hide()

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
		cursorRect:   cursorRect,
	}

	pxCanvas.renderer = renderer
	return renderer
}

func (pxCanvas *PxCanvas) TryPan(previousCoord *fyne.PointEvent, currentCoord *desktop.MouseEvent) {
	if previousCoord != nil && currentCoord.Button == desktop.MouseButtonSecondary {
		pxCanvas.Pan(previousCoord, currentCoord)
	}
}

// Add a method to directly update the canvas position and trigger a refresh
func (pxCanvas *PxCanvas) SetOffset(offset fyne.Position) {
	pxCanvas.CanvasOffset = offset
	if pxCanvas.renderer != nil {
		pxCanvas.renderer.Layout(pxCanvas.Size())
		pxCanvas.Refresh()
	}
}
