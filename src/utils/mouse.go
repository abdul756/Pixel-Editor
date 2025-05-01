package utils

import "fyne.io/fyne/v2/driver/desktop"

func (palette *Palette) MouseDown(ev *desktop.MouseEvent) {
	palette.clickHandler(palette)
	palette.Selected = true
	palette.Refresh()
}

func (palette *Palette) MouseUp(ev *desktop.MouseEvent) {}
