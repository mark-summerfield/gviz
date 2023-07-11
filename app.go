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
	me.makeFileMenu(menuBar)
	me.makeEditMenu(menuBar)
	me.makeInsertMenu(menuBar)
	me.makeZoomMenu(menuBar)
	me.makeHelpMenu(menuBar)
	vbox.Fixed(menuBar, gui.ButtonHeight)
}

func (me *App) makeFileMenu(menuBar *fltk.MenuBar) {
	menuBar.AddEx("&File", 0, nil, fltk.SUBMENU)
	for _, item := range []gui.MenuItem{
		gui.NewMenuItem("File/&New", fltk.CTRL+'n', me.onFileNew, false),
		gui.NewMenuItem("File/&Open", fltk.CTRL+'o', me.onFileOpen, false),
		gui.NewMenuItem(recentFilesMenuText, 0, nil, false),
		gui.NewMenuItem("File/&Save", fltk.CTRL+'s', me.onFileSave, false),
		gui.NewMenuItem("File/Save &As…", 0, me.onFileSaveAs, false),
		gui.NewMenuItem("File/&Export…", 0, me.onFileExport, true),
		gui.NewMenuItem("File/&Configure…", 0, me.onConfigure, true),
		gui.NewMenuItem("File/&Quit", fltk.CTRL+'q', me.onFileQuit,
			false)} {
		gui.MakeMenuItem(menuBar, item)
	}
}

func (me *App) makeEditMenu(menuBar *fltk.MenuBar) {
	menuBar.AddEx("&Edit", 0, nil, fltk.SUBMENU)
	for _, item := range []gui.MenuItem{
		gui.NewMenuItem("Edit/&Undo", fltk.CTRL+'z', me.onEditUndo, false),
		gui.NewMenuItem("Edit/&Redo", fltk.CTRL+fltk.SHIFT+'z',
			me.onEditRedo, true),
		gui.NewMenuItem("Edit/&Copy", fltk.CTRL+'c', me.onEditCopy, false),
		gui.NewMenuItem("Edit/Cu&t", fltk.CTRL+'x', me.onEditCut, false),
		gui.NewMenuItem("Edit/&Paste", fltk.CTRL+'v', me.onEditPaste, true),
		gui.NewMenuItem("Edit/&Find…", fltk.CTRL+'f', me.onEditFind, false),
		gui.NewMenuItem("Edit/Find &Again", fltk.F3, me.onEditFindAgain,
			false)} {
		gui.MakeMenuItem(menuBar, item)
	}
}

func (me *App) makeInsertMenu(menuBar *fltk.MenuBar) {
	menuBar.AddEx("&Insert", 0, nil, fltk.SUBMENU)
	menuBar.AddEx("Insert/&Attribute", 0, nil, fltk.SUBMENU)
	me.makeSubmenuTextItems(menuBar, "Insert/Attribute/", attributes)
	menuBar.AddEx("Insert/&Keyword", 0, nil, fltk.SUBMENU|fltk.MENU_DIVIDER)
	me.makeSubmenuTextItems(menuBar, "Insert/Keyword/", keywords)
	me.makeSubmenuShapeItems(menuBar, "Insert/", shapeData)
	menuBar.AddEx("Insert/E&xtra", 0, nil, fltk.SUBMENU)
	me.makeSubmenuShapeItems(menuBar, "Insert/Extra/", extraShapeData)
}

func (me *App) makeZoomMenu(menuBar *fltk.MenuBar) {
	menuBar.AddEx("&Zoom", 0, nil, fltk.SUBMENU)
	for _, item := range []gui.MenuItem{
		gui.NewMenuItem("Zoom/&In", fltk.CTRL+'+', me.onViewZoomIn, false),
		gui.NewMenuItem("Zoom/&Restore", fltk.CTRL+'=',
			me.onViewZoomRestore, false),
		gui.NewMenuItem("Zoom/&Out", fltk.CTRL+'-', me.onViewZoomOut,
			false)} {
		gui.MakeMenuItem(menuBar, item)
	}
}

func (me *App) makeHelpMenu(menuBar *fltk.MenuBar) {
	menuBar.AddEx("&Help", 0, nil, fltk.SUBMENU)
	menuBar.Add("Help/&About", me.onHelpAbout)
	menuBar.AddEx("Help/&Help", fltk.F1, me.onHelpHelp, fltk.MENU_VALUE)
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
	shapes []shapeDatum) {
	for _, shape := range shapes {
		shape := shape
		menuBar.Add(fmt.Sprintf(submenu+shape.display),
			func() { me.onInsertShape(shape.name) })
	}
}

func (me *App) makeStandardToolBar(vbox *fltk.Flex, y,
	width int) *fltk.Flex {
	hbox := gui.MakeHBox(0, y, width, gui.ButtonHeight, 0)
	me.makeStandardToolbuttons(y, hbox)
	me.makeStandardShapeToolbuttons(hbox)
	hbox.End()
	vbox.Fixed(hbox, gui.ButtonHeight)
	return hbox
}

func (me *App) makeStandardToolbuttons(y int, hbox *fltk.Flex) {
	sep := toolbuttonDatum{"", nil, ""}
	for _, toolbutton := range []toolbuttonDatum{
		{openSvg, me.onFileOpen, "Open"},
		{saveSvg, me.onFileSave, "Save"},
		sep,
		{undoSvg, me.onEditUndo, "Undo"},
		{redoSvg, me.onEditRedo, "Redo"},
		sep,
		{copySvg, me.onEditCopy, "Copy"},
		{cutSvg, me.onEditCut, "Cut"},
		{pasteSvg, me.onEditPaste, "Paste"},
		sep,
		{findSvg, me.onEditFind, "Find"},
		{findAgainSvg, me.onEditFindAgain, "Find"},
		sep,
		{zoomInSvg, me.onViewZoomIn, "Zoom In"},
		{zoomRestoreSvg, me.onViewZoomRestore, "Zoom Restore"},
		{zoomOutSvg, me.onViewZoomOut, "Zoom Out"},
		sep} {
		if toolbutton.svg == "" {
			gui.MakeSep(y, hbox)
		} else {
			button := gui.MakeToolbutton(toolbutton.svg)
			button.SetCallback(toolbutton.method)
			button.SetTooltip(toolbutton.tip)
			hbox.Fixed(button, gui.ButtonHeight)
		}
	}
}

func (me *App) makeStandardShapeToolbuttons(hbox *fltk.Flex) {
	for _, shape := range shapeData {
		shape := shape
		button := me.makeShapeToolbutton(shape)
		hbox.Fixed(button, gui.ButtonHeight)
	}
}

func (me *App) makeExtraShapesToolBar(vbox *fltk.Flex, y,
	width int) *fltk.Flex {
	hbox := gui.MakeHBox(0, y, width, gui.ButtonHeight, 0)
	for _, shape := range extraShapeData {
		if shape.svg == "" { // TODO delete once all icons done
			continue
		}
		shape := shape
		button := me.makeShapeToolbutton(shape)
		hbox.Fixed(button, gui.ButtonHeight)
	}
	hbox.End()
	vbox.Fixed(hbox, gui.ButtonHeight)
	return hbox
}

func (me *App) makeShapeToolbutton(shape shapeDatum) *fltk.Button {
	button := gui.MakeToolbutton(shape.svg)
	button.SetCallback(func() { me.onInsertShape(shape.name) })
	button.SetTooltip("Insert " +
		strings.ReplaceAll(shape.display, "&", ""))
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
