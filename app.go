// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	config *Config
	buffer *fltk.TextBuffer
	editor *fltk.TextEditor
	scroll *fltk.Scroll
	view   *fltk.Box
}

func newApp(config *Config) *App {
	app := &App{Window: nil, config: config}
	app.makeMainWindow()
	app.makeWidgets()
	app.Window.End()
	fltk.AddTimeout(0.1, func() { app.onTextChanged() })
	return app
}

func (me *App) makeMainWindow() {
	me.Window = fltk.NewWindow(me.config.Width, me.config.Height)
	if me.config.X > -1 && me.config.Y > -1 {
		me.Window.SetPosition(me.config.X, me.config.Y)
	}
	me.Window.Resizable(me.Window)
	me.Window.SetEventHandler(me.onEvent)
	me.Window.SetLabel(appName)
	addWindowIcon(me.Window, iconSvg)
}

func (me *App) makeWidgets() {
	width := me.Window.W()
	height := me.Window.H()
	var x, y int
	// tile := fltk.NewTile(x, y, width, height) // TODO if Tile added
	vbox := makeVBox(x, y, width, height, pad)
	me.makeMenuBar(vbox, width)
	y += buttonHeight
	me.makeToolBar(vbox, y, width)
	y += buttonHeight
	height -= 2 * buttonHeight
	hbox := makeHBox(x, 0, width, height, pad)
	// tile.resizable(hbox) // TODO if Tile added
	me.makePanels(x, y, width/2, height)
	hbox.End()
	vbox.End()
	// tile.end() // TODO if Tile added
}

func (me *App) makeMenuBar(vbox *fltk.Flex, width int) {
	menuBar := fltk.NewMenuBar(0, 0, width, buttonHeight)
	menuBar.Activate()
	menuBar.AddEx("&File", 0, nil, fltk.SUBMENU)
	menuBar.AddEx("File/&Open", fltk.CTRL+'o', me.onFileOpen,
		fltk.MENU_VALUE|fltk.MENU_DIVIDER)
	menuBar.AddEx("File/&Configure…", 0, me.onFileConfigure,
		fltk.MENU_VALUE|fltk.MENU_DIVIDER)
	menuBar.AddEx("File/&Quit", fltk.CTRL+'q', me.onFileQuit,
		fltk.MENU_VALUE)
	menuBar.AddEx("&Help", 0, nil, fltk.SUBMENU)
	menuBar.Add("Help/&About", me.onHelpAbout)
	vbox.Fixed(menuBar, buttonHeight)
}

func (me *App) makeToolBar(vbox *fltk.Flex, y, width int) {
	hbox := makeHBox(0, y, width, buttonHeight, 0)
	openButton := makeToolbutton(openSvg)
	openButton.SetCallback(func() { me.onFileOpen() })
	hbox.Fixed(openButton, buttonHeight)
	// TODO other toolbuttons
	hbox.End()
	vbox.Fixed(hbox, buttonHeight)
}

func (me *App) makePanels(x, y, width, height int) {
	me.buffer = fltk.NewTextBuffer()
	if me.config.ViewOnLeft {
		me.scroll = fltk.NewScroll(x, y, width, height)
		me.view = fltk.NewBox(fltk.FLAT_BOX, x, y, width, height)
		me.scroll.End()
		x += width
		me.editor = fltk.NewTextEditor(x, y, width, height)
	} else {
		me.editor = fltk.NewTextEditor(x, y, width, height)
		x += width
		me.scroll = fltk.NewScroll(x, y, width, height)
		me.view = fltk.NewBox(fltk.FLAT_BOX, x, y, width, height)
		me.scroll.End()
	}
	me.editor.SetBuffer(me.buffer)
	me.editor.SetCallback(func() { me.onTextChanged() })
	me.editor.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	me.buffer.SetText(defaultText)
}
