// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3
package main

import (
	"github.com/mark-summerfield/gviz/gui"
	"github.com/pwiecz/go-fltk"
)

const (
	findFormWidth  = 260
	findFormHeight = 120
)

type findForm struct {
	*fltk.Window
	findTextInput         *fltk.Input
	findMatchCaseCheckbox *fltk.CheckButton
	findResult            *findResult
}

type findResult struct {
	findText      *string
	findMatchCase *bool
	ok            bool
}

func newFindForm(findResult *findResult) *findForm {
	findForm := &findForm{findResult: findResult}
	findForm.Window = fltk.NewWindow(findFormWidth, findFormHeight)
	findForm.SetLabel("Find — " + appName)
	gui.AddWindowIcon(findForm.Window, iconSvg)
	findForm.makeWidgets()
	findForm.End()
	return findForm
}

func (me *findForm) makeWidgets() {
	vbox := gui.MakeVBox(0, 0, findFormWidth, findFormHeight, gui.Pad)
	hbox := me.makeFindTextRow()
	vbox.Fixed(hbox, rowHeight)
	me.makeMatchCaseRow()
	vbox.Fixed(me.findMatchCaseCheckbox, rowHeight)
	hbox = me.makeFindButtonRow()
	vbox.Fixed(hbox, rowHeight)
	vbox.End()
}

func (me *findForm) makeFindTextRow() *fltk.Flex {
	hbox := gui.MakeHBox(0, 0, findFormWidth, rowHeight, gui.Pad)
	findLabel := gui.MakeAccelLabel(colWidth, rowHeight, "Find &Text")
	me.findTextInput = fltk.NewInput(0, 0, findFormWidth-colWidth,
		rowHeight)
	me.findTextInput.SetValue(*me.findResult.findText)
	hbox.Fixed(findLabel, colWidth)
	hbox.End()
	findLabel.SetCallback(func() { me.findTextInput.TakeFocus() })
	me.findTextInput.TakeFocus()
	return hbox
}

func (me *findForm) makeMatchCaseRow() {
	me.findMatchCaseCheckbox = fltk.NewCheckButton(0, 0,
		findFormWidth, rowHeight, "Case &Sensitive")
	me.findMatchCaseCheckbox.SetValue(*me.findResult.findMatchCase)
}

func (me *findForm) makeFindButtonRow() *fltk.Flex {
	hbox := gui.MakeHBox(0, 0, findFormWidth, rowHeight, gui.Pad)
	spacerWidth := (findFormWidth - gui.ButtonWidth) / 2
	leftSpacer := gui.MakeHBox(0, 0, spacerWidth, gui.ButtonHeight, 0)
	leftSpacer.End()
	findButton := fltk.NewReturnButton(0, 0, gui.ButtonHeight,
		gui.ButtonWidth, "&Find")
	findButton.SetCallback(func() {
		*me.findResult.findText = me.findTextInput.Value()
		*me.findResult.findMatchCase = me.findMatchCaseCheckbox.Value()
		me.findResult.ok = true
		me.Destroy()
	})
	closeButton := fltk.NewButton(0, 0, gui.ButtonHeight, gui.ButtonWidth,
		"&Close")
	closeButton.SetCallback(func() { me.Destroy() })
	righttSpacer := gui.MakeHBox(spacerWidth+gui.ButtonWidth, 0,
		spacerWidth, gui.ButtonHeight, 0)
	righttSpacer.End()
	hbox.End()
	return hbox
}
