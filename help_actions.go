// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/mark-summerfield/gviz/gui"
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
	gui.AddWindowIcon(window, iconSvg)
	vbox := gui.MakeVBox(0, 0, width, height, gui.Pad)
	view := fltk.NewHelpView(0, 0, width, height-gui.ButtonHeight)
	view.TextFont(fltk.HELVETICA)
	view.SetValue(aboutHtml())
	y := height - gui.ButtonHeight
	hbox := gui.MakeHBox(0, y, width, gui.ButtonHeight, gui.Pad)
	spacerWidth := (width - gui.ButtonWidth) / 2
	leftSpacer := gui.MakeHBox(0, y, spacerWidth, gui.ButtonHeight, 0)
	leftSpacer.End()
	button := fltk.NewButton(0, 0, gui.ButtonHeight, gui.ButtonWidth,
		" &Close")
	button.SetCallback(func() { window.Destroy() })
	righttSpacer := gui.MakeHBox(spacerWidth+gui.ButtonWidth, y,
		spacerWidth, gui.ButtonHeight, 0)
	righttSpacer.End()
	hbox.Fixed(button, gui.ButtonWidth)
	hbox.End()
	vbox.Fixed(hbox, gui.ButtonHeight)
	vbox.End()
	window.End()
	return window
}
