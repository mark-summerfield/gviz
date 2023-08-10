// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: Apache-2.0

package gui

import (
	"github.com/pwiecz/go-fltk"
)

const (
	BUTTON_ONE   = 1
	BUTTON_TWO   = 2
	BUTTON_THREE = 3
)

func Ask(title, bodyText, iconSvg string, textSize int, text1,
	text2, text3 string) int {
	result := BUTTON_TWO
	if text3 != "" {
		result = BUTTON_THREE
	}
	form := makeAskForm(title, bodyText, iconSvg, textSize, text1, text2,
		text3, &result)
	form.SetModal()
	form.Show()
	for form.IsShown() {
		fltk.Wait(0.01)
	}
	return result
}

func makeAskForm(title, bodyText, iconSvg string, textSize int, text1,
	text2, text3 string, result *int) *fltk.Window {
	const height = 160
	width := 320
	if text3 != "" {
		width += ButtonWidth / 2
	}
	window := fltk.NewWindow(width, height)
	window.SetLabel(title)
	AddWindowIcon(window, iconSvg)
	vbox := MakeVBox(0, 0, width, height, Pad)
	bodyBox := fltk.NewBox(fltk.FLAT_BOX, 0, 0, width, height-ButtonHeight)
	bodyBox.SetImage(ImageForSvgText(questionSvg, 64))
	bodyBox.SetAlign(fltk.ALIGN_IMAGE_NEXT_TO_TEXT)
	bodyBox.SetLabel(bodyText)
	bodyBox.SetLabelSize(textSize)
	y := height - (ButtonHeight * 3 / 2)
	hbox := MakeHBox(0, y, width, ButtonHeight, Pad)
	var spacerWidth int
	if text3 == "" {
		spacerWidth = (width - (2 * ButtonWidth)) / 2
	} else {
		spacerWidth = (width - (3 * ButtonWidth)) / 2
	}
	leftSpacer := MakeHBox(0, y, spacerWidth, ButtonHeight, 0)
	leftSpacer.End()
	button1 := fltk.NewReturnButton(0, 0, ButtonHeight, ButtonWidth, text1)
	button1.SetCallback(func() { *result = BUTTON_ONE; window.Destroy() })
	button1.TakeFocus()
	button2 := fltk.NewButton(0, 0, ButtonHeight, ButtonWidth, text2)
	button2.SetCallback(func() { *result = BUTTON_TWO; window.Destroy() })
	var button3 *fltk.Button
	if text3 != "" {
		button3 = fltk.NewButton(0, 0, ButtonHeight, ButtonWidth, text3)
		button3.SetCallback(func() {
			*result = BUTTON_THREE
			window.Destroy()
		})
	}
	righttSpacer := MakeHBox(spacerWidth+ButtonWidth, y, spacerWidth,
		ButtonHeight, 0)
	righttSpacer.End()
	hbox.Fixed(button1, ButtonWidth)
	hbox.Fixed(button2, ButtonWidth)
	if button3 != nil {
		hbox.Fixed(button3, ButtonWidth)
	}
	hbox.End()
	vbox.Fixed(hbox, ButtonHeight)
	vbox.End()
	window.End()
	return window
}
