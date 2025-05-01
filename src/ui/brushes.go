package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/abdul756/Pixel-Editor/src/utils"
)

func SetupBrushes(app *Appinit) *fyne.Container {
	// Create a radio group for brush types
	brushTypeLabel := widget.NewLabel("Tool Type:")

	// Radio buttons for different brush types
	brushRadio := widget.NewRadioGroup(
		[]string{"Brush", "Eraser"},
		func(value string) {
			switch value {
			case "Brush":
				app.State.BrushType = utils.BrushTypePen
			case "Eraser":
				app.State.BrushType = utils.BrushTypeEraser
			}
			fmt.Println("Tool set to:", value)
		},
	)

	// Set default brush type to Brush
	brushRadio.SetSelected("Brush")

	// Add a label to show the current pixel coordinates
	positionLabel := widget.NewLabel("Position: (0, 0)")

	// Create a container for brush controls
	return container.NewVBox(
		brushTypeLabel,
		brushRadio,
		positionLabel,
	)
}
