// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package gui

import (
	"github.com/pwiecz/go-fltk"
)

func AddWindowIcon(window *fltk.Window, svgText string) {
	if image := ImageForSvgText(svgText, 0); image != nil {
		window.SetIcons([]*fltk.RgbImage{image})
	}
}

func ImageForSvgText(svgText string, size int) *fltk.RgbImage {
	if svg, err := fltk.NewSvgImageFromString(svgText); err == nil {
		if size != 0 {
			svg.Scale(size, size, true, true)
		}
		return fltk.NewRgbImageFromSvg(svg)
	}
	return nil
}

func MakeToolbutton(svgText string) *fltk.Button {
	button := fltk.NewButton(0, 0, ButtonHeight, ButtonHeight, "")
	button.ClearVisibleFocus()
	if image := ImageForSvgText(svgText, ToolbuttonIconSize); image != nil {
		button.SetImage(image)
		button.SetAlign(fltk.ALIGN_IMAGE_BACKDROP)
	}
	return button
}

func MakeHBox(x, y, width, height, spacing int) *fltk.Flex {
	return makeVHBox(fltk.ROW, x, y, width, height, spacing)
}

func MakeVBox(x, y, width, height, spacing int) *fltk.Flex {
	return makeVHBox(fltk.COLUMN, x, y, width, height, spacing)
}

func makeVHBox(kind fltk.FlexType, x, y, width, height,
	spacing int) *fltk.Flex {
	box := fltk.NewFlex(x, y, width, height)
	box.SetType(kind)
	box.SetSpacing(spacing)
	return box
}
