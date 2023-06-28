// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"os"

	"github.com/mark-summerfield/gong"
	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	config    *Config
	filename  string
	dirty     bool
	buffer    *fltk.TextBuffer
	editor    *fltk.TextEditor
	scroll    *fltk.Scroll
	view      *fltk.Box
	zoomLevel float64
}

func newApp(config *Config) *App {
	app := &App{Window: nil, config: config, zoomLevel: 1}
	app.makeMainWindow()
	app.makeWidgets()
	app.Window.End()
	if len(os.Args) > 1 && gong.FileExists(os.Args[1]) {
		fltk.AddTimeout(0.1, func() { app.loadFile(os.Args[1]) })
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
	addWindowIcon(me.Window, iconSvg)
}

func (me *App) makeWidgets() {
	width := me.Window.W()
	height := me.Window.H()
	var x, y int
	vbox := makeVBox(x, y, width, height, pad)
	me.makeMenuBar(vbox, width)
	y += buttonHeight
	me.makeToolBar(vbox, y, width)
	y += buttonHeight
	height -= 2 * buttonHeight
	tile := fltk.NewTile(x, y, width, height)
	me.makePanels(x, y, width/2, height)
	tile.End()
	vbox.End()
}

func (me *App) makeMenuBar(vbox *fltk.Flex, width int) {
	menuBar := fltk.NewMenuBar(0, 0, width, buttonHeight)
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
	menuBar.AddEx("File/&Configure…", 0, me.onFileConfigure,
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
	vbox.Fixed(menuBar, buttonHeight)
}

func (me *App) makeToolBar(vbox *fltk.Flex, y, width int) {
	hbox := makeHBox(0, y, width, buttonHeight, 0)
	openButton := makeToolbutton(openSvg)
	openButton.SetCallback(func() { me.onFileOpen() })
	openButton.SetTooltip("Open")
	hbox.Fixed(openButton, buttonHeight)
	saveButton := makeToolbutton(saveSvg)
	saveButton.SetCallback(func() { me.onFileSave() })
	saveButton.SetTooltip("Save")
	hbox.Fixed(saveButton, buttonHeight)
	sep := fltk.NewBox(fltk.THIN_DOWN_BOX, 0, y, pad, buttonHeight)
	hbox.Fixed(sep, pad)
	zoomInButton := makeToolbutton(zoomInSvg)
	zoomInButton.SetCallback(func() { me.onViewZoomIn() })
	zoomInButton.SetTooltip("Zoom In")
	hbox.Fixed(zoomInButton, buttonHeight)
	zoomRestoreButton := makeToolbutton(zoomRestoreSvg)
	zoomRestoreButton.SetCallback(func() { me.onViewZoomRestore() })
	zoomRestoreButton.SetTooltip("Zoom Restore")
	hbox.Fixed(zoomRestoreButton, buttonHeight)
	zoomOutButton := makeToolbutton(zoomOutSvg)
	zoomOutButton.SetCallback(func() { me.onViewZoomOut() })
	zoomOutButton.SetTooltip("Zoom Out")
	hbox.Fixed(zoomOutButton, buttonHeight)
	// TODO other toolbuttons, e.g., Save Cut Copy Paste etc
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
	me.initializeView()
	me.initializeEditor()
}

func (me *App) initializeView() {
	me.view.SetColor(fltk.ColorFromRgb(0xFF, 0xFF, 0xEB))
	me.view.SetAlign(fltk.ALIGN_TOP | fltk.ALIGN_LEFT | fltk.ALIGN_INSIDE |
		fltk.ALIGN_TEXT_OVER_IMAGE)
}

func (me *App) initializeEditor() {
	me.buffer.SetText(defaultText)
	me.editor.SetBuffer(me.buffer)
	me.editor.SetTextFont(fltk.COURIER)
	me.editor.SetLinenumberWidth(linoWidth)
	me.editor.SetLinenumberAlign(fltk.ALIGN_RIGHT)
	me.editor.SetLinenumberFgcolor(fltk.DARK3)
	me.editor.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	me.editor.SetCallback(func() { me.onTextChanged(true) })
}
