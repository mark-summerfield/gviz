// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

const (
	ASK_CANCELED = 0
	ASK_OK       = 1
	ASK_YES      = 1
	ASK_CANCEL   = 2
	ASK_NO       = 2
)

func askYesNo(title, bodyText string) int {
	return ask(title, bodyText, "&Yes", "&No")
}

// func askOkCancel(title, bodyText string) int {
// 	return ask(title, bodyText, "&OK", "&Cancel")
// }

func ask(title, bodyText, text1, text2 string) int {
	result := ASK_CANCELED
	form := makeAskForm(title, bodyText, text1, text2, &result)
	form.SetModal()
	form.Show()
	for form.IsShown() {
		fltk.Wait(0.01)
	}
	return result
}

func makeAskForm(title, bodyText, text1, text2 string,
	result *int) *fltk.Window {
	const (
		width  = 320
		height = 140
	)
	window := fltk.NewWindow(width, height)
	window.SetLabel(fmt.Sprintf("%s — %s", title, appName))
	addWindowIcon(window, iconSvg)
	vbox := makeVBox(0, 0, width, height, pad)
	bodyBox := fltk.NewBox(fltk.FLAT_BOX, 0, 0, width, height-buttonHeight)
	bodyBox.SetImage(imageForSvgText(questionSvg, 64))
	bodyBox.SetAlign(fltk.ALIGN_IMAGE_NEXT_TO_TEXT)
	bodyBox.SetLabel(bodyText)
	y := height - buttonHeight
	hbox := makeHBox(0, y, width, buttonHeight, pad)
	spacerWidth := (width - (2 * buttonWidth)) / 2
	leftSpacer := makeHBox(0, y, spacerWidth, buttonHeight, 0)
	leftSpacer.End()
	button1 := fltk.NewButton(0, 0, buttonHeight, buttonWidth, text1)
	button1.SetCallback(func() { *result = ASK_YES; window.Destroy() })
	button1.TakeFocus()
	button2 := fltk.NewButton(0, 0, buttonHeight, buttonWidth, text2)
	button2.SetCallback(func() { *result = ASK_NO; window.Destroy() })
	righttSpacer := makeHBox(spacerWidth+buttonWidth, y, spacerWidth,
		buttonHeight, 0)
	righttSpacer.End()
	hbox.Fixed(button1, buttonWidth)
	hbox.Fixed(button2, buttonWidth)
	hbox.End()
	vbox.Fixed(hbox, buttonHeight)
	vbox.End()
	window.End()
	return window
}
