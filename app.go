// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	// config *Config
	editor *fltk.TextEditor
	buffer *fltk.TextBuffer
	view   *fltk.Box
}

func newApp( /*config *Config*/ ) *App {
	app := &App{Window: nil}
	app.Window = fltk.NewWindow(800, 600) // TODO config.Width, config.Height)
	//if config.X > -1 && config.Y > -1 {
	//	app.Window.SetPosition(config.X, config.Y)
	//}
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
	// tile := fltk.NewTile(x, 0, width, height) // TODO if Tile added
	hbox := fltk.NewFlex(x, 0, width, height)
	// tile.resizable(hbox) // TODO if Tile added
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	width /= 2
	me.buffer = fltk.NewTextBuffer()
	me.editor = fltk.NewTextEditor(x, 0, width, height)
	me.editor.SetBuffer(me.buffer)
	x += width
	me.view = fltk.NewBox(fltk.FLAT_BOX, x, 0, width, height)
	hbox.End()
	// tile.end() // TODO if Tile added
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
		case 'q', 'Q':
			if fltk.EventState()&fltk.CTRL != 0 {
				me.onQuit()
			}
		}
	case fltk.CLOSE:
		me.onQuit()
	}
	return false
}

func (me *App) onQuit() {
	//	me.config.X = me.Window.X()
	//	me.config.Y = me.Window.Y()
	//	me.config.Width = me.Window.W()
	//	me.config.Height = me.Window.H()
	//	me.config.LastTab = me.tabs.Value()
	//	me.config.LastCategory = me.categoryChoice.Value()
	//	me.config.LastRegex = me.regexInput.Value()
	//	me.config.LastRegexText = me.regexTextInput.Value()
	//	me.config.LastUnhinted = me.accelTextBuffer.Text()
	//	me.config.LastFromIndex = me.convFromChoice.Value()
	//	me.config.LastToIndex = me.convToChoice.Value()
	//	me.config.LastAmount = me.convAmountSpinner.Value()
	//	me.config.Scale = fltk.ScreenScale(0)
	//	// config.Theme is set in callback
	//	me.config.ShowTooltips = me.showTooltipsCheckButton.Value()
	//	me.config.ShowInitialHelpText = me.showInitialHelpCheckButton.Value()
	//	me.config.CustomTitle = me.customTitleInput.Value()
	//	me.config.CustomHtml = me.customTextBuffer.Text()
	//	me.config.save()
	me.Window.Destroy()
	fmt.Println("onQuit") // TODO delete
}
