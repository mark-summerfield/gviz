// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func addWindowIcon(window *fltk.Window, svgText string) {
	if image := imageForSvgText(svgText, 0); image != nil {
		window.SetIcons([]*fltk.RgbImage{image})
	}
}

func makeToolbutton(svgText string) *fltk.Button {
	button := fltk.NewButton(0, 0, buttonHeight, buttonHeight, "")
	button.ClearVisibleFocus()
	const size = buttonHeight - pad
	if image := imageForSvgText(svgText, size); image != nil {
		button.SetImage(image)
		button.SetAlign(fltk.ALIGN_IMAGE_BACKDROP)
	}
	return button
}

func imageForSvgText(svgText string, size int) *fltk.RgbImage {
	if svg, err := fltk.NewSvgImageFromString(svgText); err == nil {
		if size != 0 {
			svg.Scale(size, size, true, true)
		}
		return fltk.NewRgbImageFromSvg(svg)
	}
	return nil
}

func makeHBox(x, y, width, height, spacing int) *fltk.Flex {
	return makeVHBox(fltk.ROW, x, y, width, height, spacing)
}

func makeVBox(x, y, width, height, spacing int) *fltk.Flex {
	return makeVHBox(fltk.COLUMN, x, y, width, height, spacing)
}

func makeVHBox(kind fltk.FlexType, x, y, width, height,
	spacing int) *fltk.Flex {
	box := fltk.NewFlex(x, y, width, height)
	box.SetType(kind)
	box.SetSpacing(spacing)
	return box
}
