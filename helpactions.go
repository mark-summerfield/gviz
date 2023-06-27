// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

const (
	width       = 360
	height      = 280
	buttonWidth = labelWidth + (labelWidth / 2)
)

func (me *App) onHelpAbout() {
	form := makeAboutForm()
	form.SetModal()
	form.Show()
}

func makeAboutForm() *fltk.Window {
	window := fltk.NewWindow(width, height)
	window.Resizable(window)
	window.SetLabel(fmt.Sprintf("About — %s", appName))
	addWindowIcon(window, iconSvg)
	vbox := makeVBox(0, 0, width, height, pad)
	view := fltk.NewHelpView(0, 0, width, height-buttonHeight)
	view.TextFont(fltk.HELVETICA)
	view.SetValue(aboutHtml())
	hbox := makeHBox(0, height-buttonHeight, width, buttonHeight, pad)
	button := fltk.NewButton(0, 0, buttonHeight, buttonWidth, " &Close")
	if image := imageForSvgText(closeSvg,
		toolbuttonIconSize); image != nil {
		button.SetImage(image)
		button.SetAlign(fltk.ALIGN_IMAGE_NEXT_TO_TEXT)
	}
	button.SetCallback(func() { window.Destroy() })
	hbox.Fixed(button, buttonWidth)
	hbox.End()
	vbox.Fixed(hbox, buttonHeight)
	vbox.End()
	window.End()
	return window
}
