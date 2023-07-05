// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package gui

import (
	"fmt"
	"strconv"
	"time"
)

func ShowAbout(appName, descHtml, iconSvg string, textSize int) {
	form := MakeInfoForm("About", appName, descHtml, iconSvg, 440, 300,
		textSize, false)
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
