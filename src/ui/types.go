package ui

import (
	"fyne.io/fyne/v2"
	"github.com/abdul756/Pixel-Editor/src/core"
	"github.com/abdul756/Pixel-Editor/src/utils"
)

type Appinit struct {
	Pixlcanvas *core.PxCanvas
	PixlWindow fyne.Window
	State      *utils.State
	Palettes   []*utils.Palette
}
