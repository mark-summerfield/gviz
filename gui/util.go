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

func MakeAccelLabel(width, height int, label string) *fltk.Button {
	button := fltk.NewButton(0, 0, width, height, label)
	button.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	button.SetBox(fltk.NO_BOX)
	button.ClearVisibleFocus()
	return button
}

func MakeSep(y int, hbox *fltk.Flex) {
	sep := fltk.NewBox(fltk.THIN_DOWN_BOX, 0, y, Pad, ButtonHeight)
	hbox.Fixed(sep, Pad)
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

type MenuItem struct {
	Text     string
	Shortcut int
	Method   func()
	Divider  bool
}

func NewMenuItem(Text string, Shortcut int, Method func(),
	Divider bool) MenuItem {
	return MenuItem{Text, Shortcut, Method, Divider}
}

func MakeMenuItem(menuBar *fltk.MenuBar, item MenuItem) {
	flag := fltk.MENU_VALUE
	if item.Divider {
		flag |= fltk.MENU_DIVIDER
	}
	menuBar.AddEx(item.Text, item.Shortcut, item.Method, flag)
}
