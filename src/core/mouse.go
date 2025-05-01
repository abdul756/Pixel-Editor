package core

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// Scrolled handles mouse wheel events for zooming
func (pxCanvas *PxCanvas) Scrolled(event *fyne.ScrollEvent) {
	// Call Scale with the scroll direction for zooming
	pxCanvas.Scale(int(event.Scrolled.DY))
}

// MouseDown handles the mouse down event
func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	// Store position for later use
	pxCanvas.mouseState.previousCoord = &fyne.PointEvent{
		Position:         ev.Position,
		AbsolutePosition: ev.AbsolutePosition,
	}

	// Set dragging state
	pxCanvas.mouseState.isDragging = true

	// Only paint with left button
	if ev.Button == desktop.MouseButtonPrimary {
		pixelX, pixelY := pxCanvas.CanvasToPixel(ev.Position)
		// Check if click is within canvas
		if pixelX >= 0 && pixelY >= 0 && pixelX < pxCanvas.PxCols && pixelY < pxCanvas.PxRows {
			pxCanvas.PaintSwatch(ev.Position)
		}
	}

	pxCanvas.Refresh()
}

// MouseUp handles the mouse up event
func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {
	// Clear dragging state
	pxCanvas.mouseState.isDragging = false
	pxCanvas.mouseState.previousCoord = nil
	pxCanvas.Refresh()
}

// MouseMoved handles the mouse moved event
func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	// Always update cursor position
	pxCanvas.cursorPos = ev.Position

	// If not dragging, just update cursor
	if !pxCanvas.mouseState.isDragging || pxCanvas.mouseState.previousCoord == nil {
		pxCanvas.Refresh()
		return
	}

	// Handle based on mouse button
	if ev.Button == desktop.MouseButtonPrimary {
		// Draw with left button if within canvas
		pixelX, pixelY := pxCanvas.CanvasToPixel(ev.Position)
		if pixelX >= 0 && pixelY >= 0 && pixelX < pxCanvas.PxCols && pixelY < pxCanvas.PxRows {
			pxCanvas.PaintSwatch(ev.Position)
		}
	} else if ev.Button == desktop.MouseButtonSecondary {
		// Pan with right button - direct call to Pan
		pxCanvas.Pan(pxCanvas.mouseState.previousCoord, ev)
	}

	// Update previous position for next move
	pxCanvas.mouseState.previousCoord = &fyne.PointEvent{
		Position:         ev.Position,
		AbsolutePosition: ev.AbsolutePosition,
	}
}

// Implement Cursor for the desktop.Cursor interface
func (pxCanvas *PxCanvas) Cursor() desktop.Cursor {
	return desktop.DefaultCursor
}

func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {
	pxCanvas.cursorPos = ev.Position
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseOut() {
	pxCanvas.cursorPos = fyne.Position{X: -1, Y: -1}
	pxCanvas.Refresh()
}

// GetCursorPixel returns the pixel coordinates under cursor
func (pxCanvas *PxCanvas) GetCursorPixel() (int, int) {
	return pxCanvas.CanvasToPixel(pxCanvas.cursorPos)
}

// Dragged implements the Draggable interface
func (pxCanvas *PxCanvas) Dragged(event *fyne.DragEvent) {
	// Direct panning with drag event
	pxCanvas.CanvasOffset = pxCanvas.CanvasOffset.Add(event.Dragged)
	pxCanvas.Refresh()
}

// DragEnd implements the Draggable interface
func (pxCanvas *PxCanvas) DragEnd() {
	// Nothing needed here
}
