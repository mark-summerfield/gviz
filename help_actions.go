// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func (me *App) onHelpAbout() {
	form := makeAboutForm()
	form.SetModal()
	form.Show()
}

func makeAboutForm() *fltk.Window {
	const (
		width  = 360
		height = 280
	)
	window := fltk.NewWindow(width, height)
	window.SetLabel(fmt.Sprintf("About — %s", appName))
	addWindowIcon(window, iconSvg)
	vbox := makeVBox(0, 0, width, height, pad)
	view := fltk.NewHelpView(0, 0, width, height-buttonHeight)
	view.TextFont(fltk.HELVETICA)
	view.SetValue(aboutHtml())
	y := height - buttonHeight
	hbox := makeHBox(0, y, width, buttonHeight, pad)
	spacerWidth := (width - buttonWidth) / 2
	leftSpacer := makeHBox(0, y, spacerWidth, buttonHeight, 0)
	leftSpacer.End()
	button := fltk.NewButton(0, 0, buttonHeight, buttonWidth, " &Close")
	button.SetCallback(func() { window.Destroy() })
	righttSpacer := makeHBox(spacerWidth+buttonWidth, y, spacerWidth,
		buttonHeight, 0)
	righttSpacer.End()
	hbox.Fixed(button, buttonWidth)
	hbox.End()
	vbox.Fixed(hbox, buttonHeight)
	vbox.End()
	window.End()
	return window
}