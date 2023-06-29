// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"os"

	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/gviz/gui"
	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	config          *Config
	filename        string
	dirty           bool
	editor          *fltk.TextEditor
	buffer          *fltk.TextBuffer
	highlightBuffer *fltk.TextBuffer
	textStyles      []fltk.StyleTableEntry
	scroll          *fltk.Scroll
	view            *fltk.Box
	zoomLevel       float64
}

func newApp(config *Config) *App {
	app := &App{Window: nil, config: config, zoomLevel: 1}
	app.makeMainWindow()
	app.makeWidgets()
	app.Window.End()
	if len(os.Args) > 1 && gong.FileExists(os.Args[1]) {
		fltk.AddTimeout(0.1, func() { app.loadFile(os.Args[1]) })
	} else if config.LastFile != "" && gong.FileExists(config.LastFile) {
		fltk.AddTimeout(0.1, func() { app.loadFile(config.LastFile) })
	} else {
		fltk.AddTimeout(0.1, func() {
			app.onTextChanged(false)
			app.dirty = false
		})
	}
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
	gui.AddWindowIcon(me.Window, iconSvg)
}

func (me *App) makeWidgets() {
	width := me.Window.W()
	height := me.Window.H()
	var x, y int
	vbox := gui.MakeVBox(x, y, width, height, gui.Pad)
	me.makeMenuBar(vbox, width)
	y += gui.ButtonHeight
	me.makeToolBar(vbox, y, width)
	y += gui.ButtonHeight
	height -= 2 * gui.ButtonHeight
	tile := fltk.NewTile(x, y, width, height)
	me.makePanels(x, y, width/2, height)
	tile.End()
	vbox.End()
}

func (me *App) makeMenuBar(vbox *fltk.Flex, width int) {
	menuBar := fltk.NewMenuBar(0, 0, width, gui.ButtonHeight)
	menuBar.Activate()
	menuBar.AddEx("&File", 0, nil, fltk.SUBMENU)
	menuBar.AddEx("File/&New", fltk.CTRL+'n', me.onFileNew,
		fltk.MENU_VALUE)
	menuBar.AddEx("File/&Open", fltk.CTRL+'o', me.onFileOpen,
		fltk.MENU_VALUE)
	menuBar.AddEx("File/&Save", fltk.CTRL+'s', me.onFileSave,
		fltk.MENU_VALUE)
	menuBar.AddEx("File/Save &As…", 0, me.onFileSaveAs,
		fltk.MENU_VALUE)
	menuBar.AddEx("File/&Export…", 0, me.onFileExport,
		fltk.MENU_VALUE|fltk.MENU_DIVIDER)
	menuBar.AddEx("File/&Configure…", 0, me.onConfigure,
		fltk.MENU_VALUE|fltk.MENU_DIVIDER)
	menuBar.AddEx("File/&Quit", fltk.CTRL+'q', me.onFileQuit,
		fltk.MENU_VALUE)
	menuBar.AddEx("&Edit", 0, nil, fltk.SUBMENU)
	// TODO Edit Cut Copy Paste & Insert etc
	menuBar.AddEx("&View", 0, nil, fltk.SUBMENU)
	menuBar.Add("View/Zoom &In", me.onViewZoomIn)
	menuBar.Add("View/Zoom &Restore", me.onViewZoomRestore)
	menuBar.Add("View/Zoom &Out", me.onViewZoomOut)
	menuBar.AddEx("&Help", 0, nil, fltk.SUBMENU)
	menuBar.Add("Help/&About", me.onHelpAbout)
	vbox.Fixed(menuBar, gui.ButtonHeight)
}

func (me *App) makeToolBar(vbox *fltk.Flex, y, width int) {
	hbox := gui.MakeHBox(0, y, width, gui.ButtonHeight, 0)
	openButton := gui.MakeToolbutton(openSvg)
	openButton.SetCallback(func() { me.onFileOpen() })
	openButton.SetTooltip("Open")
	hbox.Fixed(openButton, gui.ButtonHeight)
	saveButton := gui.MakeToolbutton(saveSvg)
	saveButton.SetCallback(func() { me.onFileSave() })
	saveButton.SetTooltip("Save")
	hbox.Fixed(saveButton, gui.ButtonHeight)
	sep := fltk.NewBox(fltk.THIN_DOWN_BOX, 0, y, gui.Pad, gui.ButtonHeight)
	hbox.Fixed(sep, gui.Pad)
	zoomInButton := gui.MakeToolbutton(zoomInSvg)
	zoomInButton.SetCallback(func() { me.onViewZoomIn() })
	zoomInButton.SetTooltip("Zoom In")
	hbox.Fixed(zoomInButton, gui.ButtonHeight)
	zoomRestoreButton := gui.MakeToolbutton(zoomRestoreSvg)
	zoomRestoreButton.SetCallback(func() { me.onViewZoomRestore() })
	zoomRestoreButton.SetTooltip("Zoom Restore")
	hbox.Fixed(zoomRestoreButton, gui.ButtonHeight)
	zoomOutButton := gui.MakeToolbutton(zoomOutSvg)
	zoomOutButton.SetCallback(func() { me.onViewZoomOut() })
	zoomOutButton.SetTooltip("Zoom Out")
	hbox.Fixed(zoomOutButton, gui.ButtonHeight)
	// TODO other toolbuttons, e.g., Save Cut Copy Paste etc
	hbox.End()
	vbox.Fixed(hbox, gui.ButtonHeight)
}

func (me *App) makePanels(x, y, width, height int) {
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
	me.initializeView()
	me.initializeEditor()
}

func (me *App) initializeView() {
	me.view.SetColor(fltk.ColorFromRgb(0xFF, 0xFF, 0xEB))
	me.view.SetAlign(fltk.ALIGN_TOP | fltk.ALIGN_LEFT | fltk.ALIGN_INSIDE |
		fltk.ALIGN_TEXT_OVER_IMAGE)
}
