package ui

import (
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

var windowBg = canvas.NewLinearGradient(
	color.NRGBA{R: 189, G: 121, B: 108, A: 255},
	color.NRGBA{R: 108, G: 67, B: 96, A: 255},
	90,
)
