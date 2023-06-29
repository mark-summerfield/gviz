// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mark-summerfield/gviz/gui"
	"github.com/pwiecz/go-fltk"
)

const (
	rowHeight = 32
	colWidth  = 60
	width     = 200
	height    = 115
)

func (me *App) onConfigure() {
	form := makeConfigForm(me)
	form.SetModal()
	form.Show()
}

func makeConfigForm(app *App) *fltk.Window {
	window := fltk.NewWindow(width, height)
	window.SetLabel("Configure — " + appName)
	gui.AddWindowIcon(window, iconSvg)
	makeWidgets(window, app)
	window.End()
	return window
}

func makeWidgets(window *fltk.Window, app *App) {
	vbox := gui.MakeVBox(0, 0, width, height, gui.Pad)
	hbox := makeScaleRow()
	vbox.Fixed(hbox, rowHeight)
	button := makeLinosRow(app)
	vbox.Fixed(button, rowHeight)
	hbox = makeButtonRow(window)
	vbox.Fixed(hbox, rowHeight)
	vbox.End()
}

func makeScaleRow() *fltk.Flex {
	hbox := gui.MakeHBox(0, 0, width, rowHeight, gui.Pad)
	scaleLabel := gui.MakeAccelLabel(colWidth, rowHeight, "&Scale")
	scaleSpinner := makeScaleSpinner()
	scaleLabel.SetCallback(func() { scaleSpinner.TakeFocus() })
	hbox.Fixed(scaleLabel, colWidth)
	hbox.End()
	scaleSpinner.TakeFocus()
	return hbox
}

func makeScaleSpinner() *fltk.Spinner {
	spinner := fltk.NewSpinner(0, 0, colWidth, rowHeight)
	spinner.SetTooltip("Sets the application's scale.")
	spinner.SetType(fltk.SPINNER_FLOAT_INPUT)
	spinner.SetMinimum(0.5)
	spinner.SetMaximum(3.5)
	spinner.SetStep(0.1)
	spinner.SetValue(float64(fltk.ScreenScale(0)))
	spinner.SetCallback(func() {
		fltk.SetScreenScale(0, float32(spinner.Value()))
	})
	return spinner
}

func makeLinosRow(app *App) *fltk.CheckButton {
	button := fltk.NewCheckButton(0, 0, width, rowHeight,
		"Show &Line Numbers")
	button.SetTooltip("Toggles the editor's line numbers")
	button.SetValue(app.config.Linos)
	button.SetCallback(func() {
		app.config.Linos = button.Value()
		app.onLinosChange()
		app.editor.Redraw()
	})
	return button
}

func makeButtonRow(window *fltk.Window) *fltk.Flex {
	hbox := gui.MakeHBox(0, 0, width, rowHeight, gui.Pad)
	spacerWidth := (width - gui.ButtonWidth) / 2
	leftSpacer := gui.MakeHBox(0, 0, spacerWidth, gui.ButtonHeight, 0)
	leftSpacer.End()
	button := fltk.NewButton(0, 0, gui.ButtonHeight, gui.ButtonWidth,
		"&Close")
	button.SetCallback(func() { window.Destroy() })
	righttSpacer := gui.MakeHBox(spacerWidth+gui.ButtonWidth, 0,
		spacerWidth, gui.ButtonHeight, 0)
	righttSpacer.End()
	hbox.End()
	return hbox
}
