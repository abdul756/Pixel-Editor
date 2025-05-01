package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/abdul756/Pixel-Editor/src/core"
	"github.com/abdul756/Pixel-Editor/src/ui"
	"github.com/abdul756/Pixel-Editor/src/utils"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("PixelEditor")

	// Set a reasonable window size
	pixelWindow.Resize(fyne.NewSize(800, 600))

	state := &utils.State{
		BrushColor:      color.NRGBA{255, 255, 255, 255},
		PaletteSelected: 0,
	}

	pixlCanvasConfig := utils.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		PxCols:       50,                    // Set to 50x50 grid
		PxRows:       50,                    // Set to 50x50 grid
		PxSize:       10,                    // Set pixel size for proper visibility
		CanvasOffset: fyne.NewPos(100, 100), // Initial offset for visibility
	}

	pixlCanvas := core.NewPxCanvas(state, pixlCanvasConfig)

	appInit := ui.Appinit{
		Pixlcanvas: pixlCanvas,
		PixlWindow: pixelWindow,
		State:      state,
		Palettes:   make([]*utils.Palette, 0, 64),
	}

	ui.Setup(&appInit)
	pixelWindow.ShowAndRun()
}
