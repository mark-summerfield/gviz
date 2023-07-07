// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/gviz/gui"
	"github.com/mark-summerfield/gviz/u"
	"github.com/pwiecz/go-fltk"
)

func (me *App) onFileNew() {
	if !me.maybeSave(false) {
		return
	}
	me.clear()
}

func (me *App) onFileOpen() {
	if !me.maybeSave(false) {
		return
	}
	chooser := fltk.NewFileChooser(u.GetPath(me.filename), "*.gv",
		fltk.FileChooser_SINGLE, fmt.Sprintf("Open — %s", appName))
	defer chooser.Destroy()
	chooser.Popup()
	names := chooser.Selection()
	if len(names) == 1 {
		me.loadFile(names[0])
	}
}

func (me *App) onFileSave() {
	me.maybeSave(false)
}

func (me *App) onFileSaveAs() {
	me.maybeSave(true)
	me.updateTitle()
}

func (me *App) onFileExport() {
	text := strings.TrimSpace(me.buffer.Text())
	if text == "" {
		me.onError(errors.New("nothing to export"))
		return
	}
	chooser := fltk.NewFileChooser(u.GetPath(me.filename),
		"PNG Files (*.png)\tSVG Files (*.svg)", fltk.FileChooser_CREATE,
		fmt.Sprintf("Export — %s", appName))
	defer chooser.Destroy()
	chooser.Popup()
	names := chooser.Selection()
	if len(names) == 1 {
		filename := names[0]
		if err := me.saveGraph(filename); err != nil {
			me.onError(err)
		} else {
			me.onInfo(fmt.Sprintf("Exported to %q.", filename))
		}
	}
}

func (me *App) onFileQuit() {
	if me.dirty && strings.TrimSpace(me.buffer.Text()) != "" &&
		gui.YesNo("Unsaved Changes — "+appName, "Save unsaved changes?",
			iconSvg, me.config.TextSize) == gui.YES &&
		!me.maybeSave(false) {
		return
	}
	me.config.X = me.Window.X()
	me.config.Y = me.Window.Y()
	me.config.Width = me.Window.W()
	me.config.Height = me.Window.H()
	me.config.LastFile = me.filename
	me.config.Scale = fltk.ScreenScale(0)
	me.config.save()
	me.Window.Destroy()
}

func (me *App) maybeSave(saveAs bool) bool {
	text := strings.TrimSpace(me.buffer.Text())
	if text == "" || text == strings.TrimSpace(defaultText) {
		return true // don't bother to save empty or default
	}
	if me.dirty || saveAs {
		if me.filename == "" || saveAs {
			chooser := fltk.NewFileChooser(u.GetPath(me.filename),
				"Graphviz Files (*.gv)", fltk.FileChooser_CREATE,
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
		if me.config.AutoFormat {
			if err := me.saveGraph(me.filename); err == nil {
				me.dirty = false
				return true
			} // else fallthrough and save what text we have
		}
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
		me.onTextChanged(false)
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
	fltk.AddTimeout(0.1, func() { me.onTextChanged(false) })
}

func (me *App) saveGraph(filename string) error {
	tempdot, err := me.getTempGraph()
	if err != nil {
		return err
	}
	defer os.Remove(tempdot.Name())
	format := "canon" // suffix: "gv"
	switch {
	case strings.HasSuffix(filename, "png"):
		format = "png"
	case strings.HasSuffix(filename, "svg"):
		format = "svg"
	}
	cmd := exec.Command(dotExe, "-T"+format, "-o"+filename, tempdot.Name())
	if err = cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (me *App) getTempGraph() (*os.File, error) {
	graph, err := me.getGraph()
	if err != nil {
		return nil, err
	}
	tempdot, err := os.CreateTemp("", "*.gv")
	if err != nil {
		return nil, err
	}
	if _, err = tempdot.Write([]byte(graph.String())); err != nil {
		return nil, err
	}
	if err = tempdot.Close(); err != nil {
		return nil, err
	}
	return tempdot, nil
}

func (me *App) getGraph() (*gographviz.Graph, error) {
	graphAst, err := gographviz.ParseString(me.buffer.Text())
	if err != nil {
		return nil, err
	}
	graph := gographviz.NewGraph()
	if err = gographviz.Analyse(graphAst, graph); err != nil {
		return nil, err
	}
	return graph, nil
}
