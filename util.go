// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func addWindowIcon(window *fltk.Window, svgText string) {
	if svg, err := fltk.NewSvgImageFromString(svgText); err == nil {
		icon := fltk.NewRgbImageFromSvg(svg)
		window.SetIcons([]*fltk.RgbImage{icon})
	}
}

func makeToolbutton(svgText string) *fltk.Button {
	button := fltk.NewButton(0, 0, buttonHeight, buttonHeight, "")
	button.ClearVisibleFocus()
	if svg, err := fltk.NewSvgImageFromString(svgText); err == nil {
		image := fltk.NewRgbImageFromSvg(svg)
		// TODO resize to buttonHeight x buttonHeight
		button.SetImage(image)
	}
	//button.SetLabelType(fltk.NO_LABEL)
	return button
}
