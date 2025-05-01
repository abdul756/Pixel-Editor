package ui

import (
	"fyne.io/fyne/v2/container"
)

func Setup(app *Appinit) {
	// Create the components
	paletteContainer := BuildPalettes(app)
	colorPicker := SetupColorPicker(app)
	brushes := SetupBrushes(app)

	// Create toolbar on the right
	rightToolbar := container.NewVBox(
		colorPicker,
		brushes,
	)

	// Create a simple layout with the canvas
	appLayout := container.NewBorder(
		nil,              // top
		paletteContainer, // bottom
		nil,              // left
		rightToolbar,     // right - now includes brushes
		app.Pixlcanvas,   // center - directly use the canvas for drawing and dragging
	)

	// Set the content
	app.PixlWindow.SetContent(appLayout)
	app.PixlWindow.CenterOnScreen()
}
