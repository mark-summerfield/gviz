// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/gviz/gui"
	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	config             *Config
	filename           string
	dirty              bool
	mainVBox           *fltk.Flex
	standardToolbar    *fltk.Flex
	extraShapesToolbar *fltk.Flex
	editor             *fltk.TextEditor
	buffer             *fltk.TextBuffer
	highlightBuffer    *fltk.TextBuffer
	textStyles         []fltk.StyleTableEntry
	scroll             *fltk.Scroll
	view               *fltk.Box
	zoomLevel          float64
	findText           string
	findMatchCase      bool
}

func newApp(config *Config) *App {
	app := &App{Window: nil, config: config, zoomLevel: 1}
	app.makeMainWindow()
	app.makeWidgets()
	app.Window.End()
	if len(os.Args) > 1 && gong.FileExists(os.Args[1]) {
		fltk.AddTimeout(smallTimeout, func() { app.loadFile(os.Args[1]) })
	} else if config.LastFile != "" && gong.FileExists(config.LastFile) {
		fltk.AddTimeout(smallTimeout,
			func() { app.loadFile(config.LastFile) })
	} else {
		fltk.AddTimeout(smallTimeout, func() {
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
	me.standardToolbar = me.makeStandardToolBar(vbox, y, width)
	y += gui.ButtonHeight
	me.extraShapesToolbar = me.makeExtraShapesToolBar(vbox, y, width)
	y += gui.ButtonHeight
	height -= 2 * gui.ButtonHeight
	tile := fltk.NewTile(x, y, width, height)
	me.makePanels(x, y, width/2, height)
	tile.End()
	vbox.End()
	me.mainVBox = vbox
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
	menuBar.AddEx("Edit/&Undo", fltk.CTRL+'z', me.onEditUndo,
		fltk.MENU_VALUE)
	menuBar.AddEx("Edit/&Redo", fltk.CTRL+fltk.SHIFT+'z', me.onEditRedo,
		fltk.MENU_VALUE|fltk.MENU_DIVIDER)
	menuBar.AddEx("Edit/&Copy", fltk.CTRL+'c', me.onEditCopy,
		fltk.MENU_VALUE)
	menuBar.AddEx("Edit/Cu&t", fltk.CTRL+'x', me.onEditCut,
		fltk.MENU_VALUE)
	menuBar.AddEx("Edit/&Paste", fltk.CTRL+'v', me.onEditPaste,
		fltk.MENU_VALUE|fltk.MENU_DIVIDER)
	menuBar.AddEx("Edit/&Find…", fltk.CTRL+'f', me.onEditFind,
		fltk.MENU_VALUE)
	menuBar.AddEx("Edit/Find &Again", fltk.F3, me.onEditFindAgain,
		fltk.MENU_VALUE)
	menuBar.AddEx("&Insert", 0, nil, fltk.SUBMENU)
	menuBar.AddEx("Insert/&Attribute", 0, nil, fltk.SUBMENU)
	me.makeSubmenuTextItems(menuBar, "Insert/Attribute/", attributes)
	menuBar.AddEx("Insert/&Keyword", 0, nil, fltk.SUBMENU|fltk.MENU_DIVIDER)
	me.makeSubmenuTextItems(menuBar, "Insert/Keyword/", keywords)
	me.makeSubmenuShapeItems(menuBar, "Insert/", shapeData)
	menuBar.AddEx("Insert/E&xtra", 0, nil, fltk.SUBMENU)
	me.makeSubmenuShapeItems(menuBar, "Insert/Extra/", extraShapeData)
	menuBar.AddEx("&View", 0, nil, fltk.SUBMENU)
	menuBar.Add("View/Zoom &In", me.onViewZoomIn)
	menuBar.Add("View/Zoom &Restore", me.onViewZoomRestore)
	menuBar.Add("View/Zoom &Out", me.onViewZoomOut)
	menuBar.AddEx("&Help", 0, nil, fltk.SUBMENU)
	menuBar.Add("Help/&About", me.onHelpAbout)
	menuBar.AddEx("Help/&Help", fltk.F1, me.onHelpHelp, fltk.MENU_VALUE)
	vbox.Fixed(menuBar, gui.ButtonHeight)
}

func (me *App) makeSubmenuTextItems(menuBar *fltk.MenuBar, submenu string,
	words []string) {
	for _, word := range words {
		word := word
		menuBar.Add(fmt.Sprintf(submenu+word),
			func() { me.onInsertWord(word) })
	}
}

func (me *App) makeSubmenuShapeItems(menuBar *fltk.MenuBar, submenu string,
	data []shapeDatum) {
	for _, datum := range data {
		datum := datum
		menuBar.Add(fmt.Sprintf(submenu+datum.display),
			func() { me.onInsertShape(datum.name) })
	}
}

func (me *App) makeStandardToolBar(vbox *fltk.Flex, y,
	width int) *fltk.Flex {
	hbox := gui.MakeHBox(0, y, width, gui.ButtonHeight, 0)
	openButton := gui.MakeToolbutton(openSvg)
	openButton.SetCallback(me.onFileOpen)
	openButton.SetTooltip("Open")
	saveButton := gui.MakeToolbutton(saveSvg)
	saveButton.SetCallback(me.onFileSave)
	saveButton.SetTooltip("Save")
	gui.MakeSep(y, hbox)
	undoButton := gui.MakeToolbutton(undoSvg)
	undoButton.SetCallback(me.onEditUndo)
	undoButton.SetTooltip("Undo")
	redoButton := gui.MakeToolbutton(redoSvg)
	redoButton.SetCallback(me.onEditRedo)
	redoButton.SetTooltip("Redo")
	gui.MakeSep(y, hbox)
	copyButton := gui.MakeToolbutton(copySvg)
	copyButton.SetCallback(me.onEditCopy)
	copyButton.SetTooltip("Copy")
	cutButton := gui.MakeToolbutton(cutSvg)
	cutButton.SetCallback(me.onEditCut)
	cutButton.SetTooltip("Cut")
	pasteButton := gui.MakeToolbutton(pasteSvg)
	pasteButton.SetCallback(me.onEditPaste)
	pasteButton.SetTooltip("Paste")
	// TODO sep + find & find again
	gui.MakeSep(y, hbox)
	zoomInButton := gui.MakeToolbutton(zoomInSvg)
	zoomInButton.SetCallback(me.onViewZoomIn)
	zoomInButton.SetTooltip("Zoom In")
	zoomRestoreButton := gui.MakeToolbutton(zoomRestoreSvg)
	zoomRestoreButton.SetCallback(me.onViewZoomRestore)
	zoomRestoreButton.SetTooltip("Zoom Restore")
	zoomOutButton := gui.MakeToolbutton(zoomOutSvg)
	zoomOutButton.SetCallback(me.onViewZoomOut)
	zoomOutButton.SetTooltip("Zoom Out")
	gui.MakeSep(y, hbox)
	for _, button := range []*fltk.Button{openButton, saveButton,
		undoButton, redoButton, copyButton, cutButton, pasteButton,
		zoomInButton, zoomRestoreButton, zoomOutButton} {
		hbox.Fixed(button, gui.ButtonHeight)
	}
	for _, datum := range shapeData {
		datum := datum
		button := me.makeShapeToolbutton(datum)
		hbox.Fixed(button, gui.ButtonHeight)
	}
	hbox.End()
	vbox.Fixed(hbox, gui.ButtonHeight)
	return hbox
}

func (me *App) makeExtraShapesToolBar(vbox *fltk.Flex, y,
	width int) *fltk.Flex {
	hbox := gui.MakeHBox(0, y, width, gui.ButtonHeight, 0)
	for _, datum := range extraShapeData {
		if datum.svg == "" { // TODO delete once all icons done
			continue
		}
		datum := datum
		button := me.makeShapeToolbutton(datum)
		hbox.Fixed(button, gui.ButtonHeight)
	}
	hbox.End()
	vbox.Fixed(hbox, gui.ButtonHeight)
	return hbox
}

func (me *App) makeShapeToolbutton(datum shapeDatum) *fltk.Button {
	button := gui.MakeToolbutton(datum.svg)
	button.SetCallback(func() { me.onInsertShape(datum.name) })
	button.SetTooltip("Insert " +
		strings.ReplaceAll(datum.display, "&", ""))
	return button
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
