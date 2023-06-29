// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package gui

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func ShowAbout(appName, descHtml, iconSvg string, textSize int) {
	form := makeAboutForm(appName, descHtml, iconSvg, textSize)
	form.SetModal()
	form.Show()
}

func makeAboutForm(appName, descHtml, iconSvg string,
	textSize int) *fltk.Window {
	const (
		width  = 440
		height = 320
	)
	window := fltk.NewWindow(width, height)
	window.Resizable(window)
	window.SetLabel(fmt.Sprintf("About — %s", appName))
	AddWindowIcon(window, iconSvg)
	vbox := MakeVBox(0, 0, width, height, Pad)
	view := fltk.NewHelpView(0, 0, width, height-ButtonHeight)
	view.TextFont(fltk.HELVETICA)
	view.TextSize(textSize)
	view.SetValue(descHtml)
	y := height - ButtonHeight
	hbox := MakeHBox(0, y, width, ButtonHeight, Pad)
	spacerWidth := (width - ButtonWidth) / 2
	leftSpacer := MakeHBox(0, y, spacerWidth, ButtonHeight, 0)
	leftSpacer.End()
	button := fltk.NewReturnButton(0, 0, ButtonHeight, ButtonWidth,
		"&Close")
	button.SetCallback(func() { window.Destroy() })
	righttSpacer := MakeHBox(spacerWidth+ButtonWidth, y, spacerWidth,
		ButtonHeight, 0)
	righttSpacer.End()
	hbox.Fixed(button, ButtonWidth)
	hbox.End()
	vbox.Fixed(hbox, ButtonHeight)
	vbox.End()
	window.End()
	return window
}
