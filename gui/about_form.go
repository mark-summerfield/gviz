// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package gui

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func ShowAbout(appName, descHtml, iconSvg string) {
	form := makeAboutForm(appName, descHtml, iconSvg)
	form.SetModal()
	form.Show()
}

func makeAboutForm(appName, descHtml, iconSvg string) *fltk.Window {
	const (
		width  = 360
		height = 280
	)
	window := fltk.NewWindow(width, height)
	window.SetLabel(fmt.Sprintf("About — %s", appName))
	AddWindowIcon(window, iconSvg)
	vbox := MakeVBox(0, 0, width, height, Pad)
	view := fltk.NewHelpView(0, 0, width, height-ButtonHeight)
	view.TextFont(fltk.HELVETICA)
	view.SetValue(descHtml)
	y := height - ButtonHeight
	hbox := MakeHBox(0, y, width, ButtonHeight, Pad)
	spacerWidth := (width - ButtonWidth) / 2
	leftSpacer := MakeHBox(0, y, spacerWidth, ButtonHeight, 0)
	leftSpacer.End()
	button := fltk.NewButton(0, 0, ButtonHeight, ButtonWidth, " &Close")
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
