package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lusingander/colorpicker"
)

func SetupColorPicker(app *Appinit) *fyne.Container {
	picker := colorpicker.New(200 /* height */, colorpicker.StyleHue /* Style */)
	picker.SetOnChanged(func(c color.Color) {
		app.State.BrushColor = c
		app.Palettes[app.State.PaletteSelected].SetColor(c)

	})

	return container.NewVBox(picker)
}

// you can use it just like any other Fyne widget
