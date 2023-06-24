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
	buffer *fltk.TextBuffer
	editor *fltk.TextEditor
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
	addWindowIcon(app.Window, iconSvg)
	app.addWidgets()
	app.Window.End()
	return app
}

func (me *App) addWidgets() {
	width := me.Window.W()
	height := me.Window.H()
	var x, y int
	// tile := fltk.NewTile(x, y, width, height) // TODO if Tile added
	vbox := makeVBox(x, y, width, height, pad)
	me.addMenuBar(vbox, width)
	y += buttonHeight
	me.addToolBar(vbox, y, width)
	y += buttonHeight
	height -= 2 * buttonHeight
	hbox := makeHBox(x, 0, width, height, pad)
	// tile.resizable(hbox) // TODO if Tile added
	me.addPanels(x, y, width/2, height)
	hbox.End()
	vbox.End()
	// tile.end() // TODO if Tile added
}

func (me *App) addMenuBar(vbox *fltk.Flex, width int) {
	menuBar := fltk.NewMenuBar(0, 0, width, buttonHeight)
	menuBar.Activate()
	menuBar.AddEx("&File", 0, nil, fltk.SUBMENU)
	menuBar.AddEx("File/&Open", fltk.CTRL+'o', me.onFileOpen,
		fltk.MENU_VALUE|fltk.MENU_DIVIDER)
	menuBar.AddEx("File/&Configure…", 0, me.onFileConfigure,
		fltk.MENU_VALUE|fltk.MENU_DIVIDER)
	menuBar.AddEx("File/&Quit", fltk.CTRL+'q', me.onQuit, fltk.MENU_VALUE)
	menuBar.AddEx("&Help", 0, nil, fltk.SUBMENU)
	menuBar.Add("Help/&About", me.onHelpAbout)
	vbox.Fixed(menuBar, buttonHeight)
}

func (me *App) addToolBar(vbox *fltk.Flex, y, width int) {
	hbox := makeHBox(0, y, width, buttonHeight, 0)
	openButton := makeToolbutton(openSvg)
	openButton.SetCallback(func() { me.onFileOpen() })
	hbox.Fixed(openButton, buttonHeight)
	// TODO other toolbuttons
	hbox.End()
	vbox.Fixed(hbox, buttonHeight)
}

func (me *App) addPanels(x, y, width, height int) {
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
	me.buffer.SetText(defaultText)
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
			return true // ignore
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
