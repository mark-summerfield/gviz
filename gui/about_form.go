// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package gui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pwiecz/go-fltk"
)

func ShowAbout(appName, descHtml, iconSvg string, textSize int) {
	form := makeAboutForm(appName, descHtml, iconSvg, textSize)
	form.SetModal()
	form.Show()
}

func AboutYear(year int) string {
	y := time.Now().Year()
	if y == year {
		return strconv.Itoa(year)
	} else {
		return fmt.Sprintf("%d-%d", year, y-2000)
	}
}

func makeAboutForm(appName, descHtml, iconSvg string,
	textSize int) *fltk.Window {
	const (
		width  = 440
		height = 320
	)
	window := fltk.NewWindow(width, height)
	window.SetLabel(fmt.Sprintf("About — %s", appName))
	AddWindowIcon(window, iconSvg)
	vbox := MakeVBox(0, 0, width, height, Pad)
	view := fltk.NewHelpView(0, 0, width, height-ButtonHeight)
	view.TextFont(fltk.HELVETICA)
	view.TextSize(textSize)
	view.SetValue(descHtml)
	y := height - ButtonHeight
	hbox := MakeHBox(0, y, width, ButtonHeight, Pad)
	spacerWidth := (width - ReturnButtonWidth) / 2
	leftSpacer := MakeHBox(0, y, spacerWidth, ButtonHeight, 0)
	leftSpacer.End()
	button := fltk.NewReturnButton(0, 0, ButtonHeight, ReturnButtonWidth,
		"&Close")
	button.SetCallback(func() { window.Destroy() })
	righttSpacer := MakeHBox(spacerWidth+ReturnButtonWidth, y, spacerWidth,
		ButtonHeight, 0)
	righttSpacer.End()
	hbox.Fixed(button, ButtonWidth)
	hbox.End()
	vbox.Fixed(hbox, ButtonHeight)
	vbox.End()
	window.End()
	return window
}
