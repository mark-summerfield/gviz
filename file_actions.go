// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark-summerfield/gong"
	"github.com/pwiecz/go-fltk"
)

func (me *App) onFileNew() {
	if !me.maybeSave() {
		return
	}
	me.clear()
}

func (me *App) onFileOpen() {
	if !me.maybeSave() {
		return
	}
	chooser := fltk.NewFileChooser(getPath(me.filename), "*.gv",
		fltk.FileChooser_SINGLE, fmt.Sprintf("Open — %s", appName))
	defer chooser.Destroy()
	chooser.Popup()
	names := chooser.Selection()
	if len(names) == 1 {
		me.loadFile(names[0])
	}
}

func (me *App) onFileSave() {
	me.maybeSave()
}

func (me *App) onFileSaveAs() {
	filename := me.filename
	me.filename = ""
	me.dirty = true
	if me.maybeSave() {
		me.updateTitle()
	} else {
		me.filename = filename
	}
}

func (me *App) onFileExport() {
	fmt.Println("onFileExport") // TODO
}

func (me *App) onFileConfigure() {
	fmt.Println("onFileConfigure") // TODO
}

func (me *App) onFileQuit() {
	if me.dirty && strings.TrimSpace(me.buffer.Text()) != "" &&
		askYesNo("Unsaved Changes", "Save unsaved changes?") == ASK_YES &&
		!me.maybeSave() {
		return
	}
	me.config.X = me.Window.X()
	me.config.Y = me.Window.Y()
	me.config.Width = me.Window.W()
	me.config.Height = me.Window.H()
	// TODO
	// App Scale & Image Zoom & ViewOnLeft are set in config dialog
	me.config.save()
	me.Window.Destroy()
}

func (me *App) maybeSave() bool {
	text := strings.TrimSpace(me.buffer.Text())
	if text == "" || text == strings.TrimSpace(defaultText) {
		return true // don't bother to save empty or default
	}
	if me.dirty {
		if me.filename == "" {
			chooser := fltk.NewFileChooser(getPath(me.filename), "*.gv",
				fltk.FileChooser_CREATE,
				fmt.Sprintf("Save As — %s", appName))
			defer chooser.Destroy()
			chooser.Popup()
			names := chooser.Selection()
			if len(names) == 1 {
				me.filename = names[0]
			} else {
				return false // didn't choose a name
			}
		}
		text += "\n"
		if err := os.WriteFile(me.filename, []byte(text),
			gong.ModeUserRW); err != nil {
			me.onError(err)
			return false
		} else {
			me.dirty = false
		}
	}
	return true
}

func (me *App) loadFile(filename string) {
	raw, err := os.ReadFile(filename)
	if err == nil {
		me.filename = filename
		me.buffer.SetText(string(raw))
		me.onTextChanged()
		me.updateTitle()
		me.dirty = false
		fltk.AddTimeout(0.1, func() { me.scroll.ScrollTo(0, 0) })
	} else {
		me.onError(err)
	}
}

func (me *App) updateTitle() {
	if me.filename != "" {
		me.Window.SetLabel(fmt.Sprintf("%s — %s", appName,
			filepath.Base(me.filename)))
	} else {
		me.Window.SetLabel(appName)
	}
}

func (me *App) clear() {
	me.filename = ""
	me.buffer.SetText(defaultText)
	me.updateTitle()
	if png, err := fltk.NewPngImageFromData(dummyPng); err == nil {
		me.view.SetImage(png)
	}
	me.view.SetLabelColor(fltk.BLUE)
	me.view.SetLabel("Edit graphviz text")
	me.dirty = false
	me.view.Redraw()
}
