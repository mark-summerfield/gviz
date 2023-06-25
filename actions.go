// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func (me *App) onEvent(event fltk.Event) bool {
	key := fltk.EventKey()
	switch fltk.EventType() {
	case fltk.SHORTCUT:
		if key == fltk.ESCAPE {
			return true // ignore
		}
	case fltk.KEY:
		switch key {
		case fltk.HELP, fltk.F1:
			return true // ignore
		}
	case fltk.CLOSE:
		me.onFileQuit()
	}
	return false
}

func (me *App) onFileOpen() {
	fmt.Println("onFileOpen") // TODO
}

func (me *App) onFileConfigure() {
	fmt.Println("onFileConfigure") // TODO
}

func (me *App) onFileQuit() {
	me.config.X = me.Window.X()
	me.config.Y = me.Window.Y()
	me.config.Width = me.Window.W()
	me.config.Height = me.Window.H()
	// TODO:
	// Scale & ViewOnLeft are set in config dialog
	me.config.save()
	me.Window.Destroy()
	fmt.Println("onFileQuit") // TODO delete
}

func (me *App) onHelpAbout() {
	fmt.Println("onHelpAbout") // TODO
}
