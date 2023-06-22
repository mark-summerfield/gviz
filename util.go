// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func addIcon(window *fltk.Window, svgText string) {
	if svg, err := fltk.NewSvgImageFromString(svgText); err == nil {
		icon := fltk.NewRgbImageFromSvg(svg)
		window.SetIcons([]*fltk.RgbImage{icon})
	}
}
