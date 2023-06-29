// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mark-summerfield/gviz/gui"
)

func (me *App) onHelpAbout() {
	gui.ShowAbout(appName, aboutHtml(), iconSvg, me.config.TextSize-1)
}
