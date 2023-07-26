// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mark-summerfield/gviz/gui"
	"github.com/pwiecz/go-fltk"
)

type configForm struct {
	*fltk.Window
	width  int
	height int
}

func newConfigForm(app *App) configForm {
	form := configForm{width: 240, height: 300}
	form.Window = fltk.NewWindow(form.width, form.height)
	form.Window.SetLabel("Configure — " + appName)
	gui.AddWindowIcon(form.Window, getEmbStr(iconSvg))
	form.makeWidgets(app)
	form.Window.End()
	return form
}

func (me *configForm) makeWidgets(app *App) {
	vbox := gui.MakeVBox(0, 0, me.width, me.height, gui.Pad)
	hbox := me.makeScaleRow()
	vbox.Fixed(hbox, rowHeight)
	hbox = me.makeTextSizeRow(app)
	vbox.Fixed(hbox, rowHeight)
	button := me.makeLinosRow(app)
	vbox.Fixed(button, rowHeight)
	button = me.makeStandardToolbarRow(app)
	vbox.Fixed(button, rowHeight)
	button = me.makeExtraShapesRow(app)
	vbox.Fixed(button, rowHeight)
	button = me.makeViewOnLeftRow(app)
	vbox.Fixed(button, rowHeight)
	button = me.makeFormatRow(app)
	vbox.Fixed(button, rowHeight)
	hbox = me.makeButtonRow()
	vbox.Fixed(hbox, rowHeight)
	vbox.End()
}

func (me *configForm) makeScaleRow() *fltk.Flex {
	hbox := gui.MakeHBox(0, 0, me.width, rowHeight, gui.Pad)
	scaleLabel := gui.MakeAccelLabel(colWidth, rowHeight, "&Scale")
	scaleSpinner := me.makeScaleSpinner()
	scaleLabel.SetCallback(func() { scaleSpinner.TakeFocus() })
	hbox.Fixed(scaleLabel, colWidth)
	hbox.End()
	scaleSpinner.TakeFocus()
	return hbox
}

func (me *configForm) makeScaleSpinner() *fltk.Spinner {
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

func (me *configForm) makeTextSizeRow(app *App) *fltk.Flex {
	hbox := gui.MakeHBox(0, 0, me.width, rowHeight, gui.Pad)
	sizeLabel := gui.MakeAccelLabel(gui.LabelWidth, gui.ButtonHeight,
		"&Text Size")
	spinner := fltk.NewSpinner(0, 0, gui.LabelWidth, gui.ButtonHeight)
	spinner.SetTooltip("Set the size of the editor's text; you may need " +
		"to quit and rerun for this to take effect")
	spinner.SetType(fltk.SPINNER_INT_INPUT)
	spinner.SetMinimum(10)
	spinner.SetMaximum(20)
	spinner.SetValue(float64(app.config.TextSize))
	spinner.SetCallback(func() {
		size := int(spinner.Value())
		app.config.TextSize = size
		app.editor.SetTextSize(size)
		app.editor.SetLinenumberSize(size - 1)
		fltk.AddTimeout(tinyTimeout, func() { app.editor.Redraw() })
	})
	sizeLabel.SetCallback(func() { spinner.TakeFocus() })
	hbox.Fixed(sizeLabel, gui.LabelWidth)
	hbox.End()
	return hbox
}

func (me *configForm) makeLinosRow(app *App) *fltk.CheckButton {
	button := fltk.NewCheckButton(0, 0, me.width, rowHeight,
		"Show &Line Numbers")
	button.SetTooltip("Toggles the editor's line numbers")
	button.SetValue(app.config.Linos)
	button.SetCallback(func() {
		app.config.Linos = button.Value()
		app.onLinosChange()
		fltk.AddTimeout(tinyTimeout, func() { app.editor.Redraw() })
	})
	return button
}

func (me *configForm) makeStandardToolbarRow(app *App) *fltk.CheckButton {
	button := fltk.NewCheckButton(0, 0, me.width, rowHeight,
		"Sho&w Standard Toolbar")
	button.SetTooltip("Toggles the standard toolbar")
	button.SetValue(app.config.ShowStandardToolbar)
	button.SetCallback(func() {
		app.config.ShowStandardToolbar = button.Value()
		app.onToggleStandardToolbar(true)
	})
	return button
}

func (me *configForm) makeExtraShapesRow(app *App) *fltk.CheckButton {
	button := fltk.NewCheckButton(0, 0, me.width, rowHeight,
		"Show E&xtra Shapes Toolbar")
	button.SetTooltip("Toggles the extra shapes toolbar")
	button.SetValue(app.config.ShowExtraShapesToolbar)
	button.SetCallback(func() {
		app.config.ShowExtraShapesToolbar = button.Value()
		app.onToggleExtraShapesToolbar(true)
	})
	return button
}

func (me *configForm) makeViewOnLeftRow(app *App) *fltk.CheckButton {
	button := fltk.NewCheckButton(0, 0, me.width, rowHeight,
		"&View on Left")
	button.SetTooltip("if checked the view is on the left and the editor " +
		"is on the right. If changed, quit and rerun for the setting to " +
		"take effect")
	button.SetValue(app.config.ViewOnLeft)
	button.SetCallback(func() {
		app.config.ViewOnLeft = button.Value()
	})
	return button
}

func (me *configForm) makeFormatRow(app *App) *fltk.CheckButton {
	button := fltk.NewCheckButton(0, 0, me.width, rowHeight,
		"&Format on Save")
	button.SetTooltip("if checked the dot text will automatically " +
		"be saved in canonical format")
	button.SetValue(app.config.AutoFormat)
	button.SetCallback(func() {
		app.config.AutoFormat = button.Value()
	})
	return button
}

func (me *configForm) makeButtonRow() *fltk.Flex {
	hbox := gui.MakeHBox(0, 0, me.width, rowHeight, gui.Pad)
	spacerWidth := (me.width - gui.ButtonWidth) / 2
	leftSpacer := gui.MakeHBox(0, 0, spacerWidth, gui.ButtonHeight, 0)
	leftSpacer.End()
	button := fltk.NewButton(0, 0, gui.ButtonHeight, gui.ButtonWidth,
		"&Close")
	button.SetCallback(func() { me.Window.Destroy() })
	righttSpacer := gui.MakeHBox(spacerWidth+gui.ButtonWidth, 0,
		spacerWidth, gui.ButtonHeight, 0)
	righttSpacer.End()
	hbox.End()
	return hbox
}
