package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/abdul756/Pixel-Editor/src/utils"
)

func BuildPalettes(app *Appinit) *fyne.Container {
	canvasPalattes := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Palettes); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		p := utils.NewPalette(app.State, initialColor, i, func(p *utils.Palette) {
			for j := 0; j < len(app.Palettes); j++ {
				app.Palettes[j].Selected = false
				canvasPalattes[j].Refresh()
			}
			app.State.PaletteSelected = p.PaletteIndex
			app.State.BrushColor = p.Color

		})
		if i == 0 {
			p.Selected = true
			app.State.PaletteSelected = 0
			p.Refresh()

		}

		app.Palettes = append(app.Palettes, p)
		canvasPalattes = append(canvasPalattes, p)
	}
	return container.NewGridWrap(fyne.NewSize(20, 20), canvasPalattes...)
}
