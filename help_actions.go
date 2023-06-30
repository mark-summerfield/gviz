// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mark-summerfield/gviz/gui"
)

func (me *App) onHelpAbout() {
	descHtml := gui.DescHtml(appName, Version,
		"An application for editing and viewing GraphViz diagrams.", // desc
		"https://github.com/mark-summerfield/gviz",                  // url
		"Mark Summerfield", gui.AboutYear(2023)) // author, year
	gui.ShowAbout(appName, descHtml, iconSvg, me.config.TextSize-1)
}
