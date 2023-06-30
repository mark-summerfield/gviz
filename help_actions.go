// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mark-summerfield/gviz/gui"
)

func (me *App) onHelpAbout() {
	descHtml := gui.DescHtml(appName, Version, description, url, author,
		gui.AboutYear(2023))
	gui.ShowAbout(appName, descHtml, iconSvg, me.config.TextSize-1)
}
