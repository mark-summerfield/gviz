// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	config *Config
	editor *fltk.TextEditor
	buffer *fltk.TextBuffer
	view   *fltk.Box
}

func newApp(config *Config) *App {
	app := &App{Window: nil, config: config}
	app.Window = fltk.NewWindow(config.Width, config.Height)
	if config.X > -1 && config.Y > -1 {
		app.Window.SetPosition(config.X, config.Y)
	}
	app.Window.Resizable(app.Window)
	app.Window.SetEventHandler(app.onEvent)
	app.Window.SetLabel(appName)
	addIcon(app.Window, iconSvg)
	app.addPanels()
	app.Window.End()
	return app
}

func (me *App) addPanels() {
	width := me.Window.W()
	height := me.Window.H()
	x := 0
	y := 0
	// tile := fltk.NewTile(x, 0, width, height) // TODO if Tile added
	vbox := fltk.NewFlex(x, y, width, height)
	vbox.SetType(fltk.COLUMN)
	vbox.SetSpacing(pad)
	menuBar := me.addMenuBar(width)
	vbox.Fixed(menuBar, buttonHeight)
	y += buttonHeight
	// TODO toolbar
	// y += buttonHeight
	height -= 2 * buttonHeight
	hbox := fltk.NewFlex(x, 0, width, height)
	// tile.resizable(hbox) // TODO if Tile added
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	width /= 2
	me.buffer = fltk.NewTextBuffer()
	if me.config.ViewOnLeft {
		me.view = fltk.NewBox(fltk.FLAT_BOX, x, y, width, height)
		x += width
		me.editor = fltk.NewTextEditor(x, y, width, height)
	} else {
		me.editor = fltk.NewTextEditor(x, y, width, height)
		x += width
		me.view = fltk.NewBox(fltk.FLAT_BOX, x, y, width, height)
	}
	me.editor.SetBuffer(me.buffer)
	hbox.End()
	vbox.End()
	// tile.end() // TODO if Tile added
}

func (me *App) addMenuBar(width int) *fltk.MenuBar {
	menuBar := fltk.NewMenuBar(0, 0, width, buttonHeight)
	menuBar.Activate()
	menuBar.AddEx("&File", 0, nil, fltk.SUBMENU)
	menuBar.AddEx("File/&Open", fltk.CTRL+'o', me.onFileOpen,
		fltk.MENU_VALUE)
	// TODO separator
	menuBar.Add("File/&Configure…", me.onFileConfigure)
	// TODO separator
	menuBar.AddEx("File/&Quit", fltk.CTRL+'q', me.onQuit, fltk.MENU_VALUE)
	menuBar.AddEx("&Help", 0, nil, fltk.SUBMENU)
	menuBar.Add("Help/&About", me.onHelpAbout)
	return menuBar
}

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
			fmt.Println("F1")
			return true
		}
	case fltk.CLOSE:
		me.onQuit()
	}
	return false
}

func (me *App) onFileOpen() {
	fmt.Println("onFileOpen") // TODO
}

func (me *App) onFileConfigure() {
	fmt.Println("onFileConfigure") // TODO
}

func (me *App) onQuit() {
	me.config.X = me.Window.X()
	me.config.Y = me.Window.Y()
	me.config.Width = me.Window.W()
	me.config.Height = me.Window.H()
	// TODO:
	// Scale & ViewOnLeft are set in config dialog
	me.config.save()
	me.Window.Destroy()
	fmt.Println("onQuit") // TODO delete
}

func (me *App) onHelpAbout() {
	fmt.Println("onHelpAbout") // TODO
}
