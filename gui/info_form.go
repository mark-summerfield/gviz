// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package gui

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func MakeInfoForm(title, appName, descHtml, iconSvg string, width, height,
	textSize int, resizable bool) *fltk.Window {
	window := fltk.NewWindow(width, height)
	if resizable {
		window.Resizable(window)
	}
	window.SetLabel(fmt.Sprintf("%s — %s", title, appName))
	AddWindowIcon(window, iconSvg)
	vbox := MakeVBox(0, 0, width, height, Pad)
	view := fltk.NewHelpView(0, 0, width, height-ButtonHeight)
	view.TextFont(fltk.HELVETICA)
	view.TextSize(textSize)
	view.SetValue(descHtml)
	y := height - ButtonHeight
	hbox := MakeHBox(0, y, width, ButtonHeight, Pad)
	spacerWidth := (width - ReturnButtonWidth) / 2
	leftSpacer := MakeHBox(0, y, spacerWidth, ButtonHeight, 0)
	leftSpacer.End()
	button := fltk.NewReturnButton(0, 0, ButtonHeight, ReturnButtonWidth,
		"&Close")
	button.SetCallback(func() { window.Destroy() })
	righttSpacer := MakeHBox(spacerWidth+ReturnButtonWidth, y, spacerWidth,
		ButtonHeight, 0)
	righttSpacer.End()
	hbox.Fixed(button, ButtonWidth)
	hbox.End()
	vbox.Fixed(hbox, ButtonHeight)
	vbox.End()
	window.End()
	return window
}
