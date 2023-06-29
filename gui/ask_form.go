// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package gui

import (
	"github.com/pwiecz/go-fltk"
)

const (
	CANCELED = 0
	OK       = 1
	YES      = 1
	CANCEL   = 2
	NO       = 2
)

func YesNo(title, bodyText, iconSvg string) int {
	return ask(title, bodyText, iconSvg, "&Yes", "&No")
}

// func OkCancel(title, bodyText, iconSvg string) int {
// 	return ask(title, bodyText, iconSvg, "&OK", "&Cancel")
// }

func ask(title, bodyText, iconSvg, text1, text2 string) int {
	result := CANCELED
	form := makeAskForm(title, bodyText, iconSvg, text1, text2, &result)
	form.SetModal()
	form.Show()
	for form.IsShown() {
		fltk.Wait(0.01)
	}
	return result
}

func makeAskForm(title, bodyText, iconSvg, text1, text2 string,
	result *int) *fltk.Window {
	const (
		width  = 320
		height = 140
	)
	window := fltk.NewWindow(width, height)
	window.SetLabel(title)
	AddWindowIcon(window, iconSvg)
	vbox := MakeVBox(0, 0, width, height, Pad)
	bodyBox := fltk.NewBox(fltk.FLAT_BOX, 0, 0, width, height-ButtonHeight)
	bodyBox.SetImage(ImageForSvgText(questionSvg, 64))
	bodyBox.SetAlign(fltk.ALIGN_IMAGE_NEXT_TO_TEXT)
	bodyBox.SetLabel(bodyText)
	y := height - ButtonHeight
	hbox := MakeHBox(0, y, width, ButtonHeight, Pad)
	spacerWidth := (width - (2 * ButtonWidth)) / 2
	leftSpacer := MakeHBox(0, y, spacerWidth, ButtonHeight, 0)
	leftSpacer.End()
	button1 := fltk.NewReturnButton(0, 0, ButtonHeight, ButtonWidth, text1)
	button1.SetCallback(func() { *result = YES; window.Destroy() })
	button1.TakeFocus()
	button2 := fltk.NewButton(0, 0, ButtonHeight, ButtonWidth, text2)
	button2.SetCallback(func() { *result = NO; window.Destroy() })
	righttSpacer := MakeHBox(spacerWidth+ButtonWidth, y, spacerWidth,
		ButtonHeight, 0)
	righttSpacer.End()
	hbox.Fixed(button1, ButtonWidth)
	hbox.Fixed(button2, ButtonWidth)
	hbox.End()
	vbox.Fixed(hbox, ButtonHeight)
	vbox.End()
	window.End()
	return window
}
