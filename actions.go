// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"math"
	"os"

	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/gviz/gui"
	"github.com/mark-summerfield/gviz/u"
	"github.com/pwiecz/go-fltk"
)

func (me *App) onEvent(event fltk.Event) bool {
	key := fltk.EventKey()
	switch fltk.EventType() {
	case fltk.SHORTCUT:
		if key == fltk.ESCAPE {
			return true // ignore
		}
	case fltk.KEY:
		switch key {
		case fltk.HELP:
			me.onHelpHelp()
			return true
		}
	case fltk.CLOSE:
		me.onFileQuit()
	}
	return false
}

func (me *App) onTextChanged(changed bool) {
	if changed {
		me.dirty = true
	}
	me.applySyntaxHighlighting()
	temppng := fmt.Sprintf("gviz-%d.png", os.Getpid())
	if err := me.exportGraph(temppng); err != nil {
		me.onError(err)
		me.updateView()
		return
	}
	defer os.Remove(temppng)
	png, err := fltk.NewPngImageLoad(temppng)
	if err != nil {
		me.onError(err)
		me.updateView()
		return
	}
	me.updateImage(png)
}

func (me *App) updateImage(png *fltk.PngImage) {
	me.clearView()
	if !gong.IsRealClose(1.0, me.zoomLevel) {
		w := int(math.Round(float64(png.W()) * me.zoomLevel))
		h := int(math.Round(float64(png.H()) * me.zoomLevel))
		png.Scale(w, h, true, true)
	}
	if me.view.W() < u.Max(me.scroll.W(), png.W()) ||
		me.view.H() < u.Max(me.scroll.H(), png.H()) {
		me.view.Resize(0, 0, png.W()+gui.Border, png.H()+gui.Border)
	}
	me.view.SetImage(png)
	me.updateView()
}

func (me *App) updateView() {
	fltk.AddTimeout(tinyTimeout, func() {
		me.view.Redraw()
		me.scroll.ScrollTo(0, 0)
	})
}

func (me *App) onLinosChange() {
	if me.config.Linos {
		me.editor.SetLinenumberWidth(linoWidth)
		me.editor.SetLinenumberAlign(fltk.ALIGN_RIGHT)
		me.editor.SetLinenumberFgcolor(fltk.DARK3)
	} else {
		me.editor.SetLinenumberWidth(0)
	}
}

func (me *App) clearView() {
	me.view.SetLabelColor(fltk.BLACK)
	me.view.SetLabel("")
}

func (me *App) onError(err error) {
	me.view.SetLabelColor(fltk.RED)
	me.view.SetLabel(gong.Wrapped(err.Error(), 40))
}

func (me *App) onInfo(info string) {
	me.view.SetLabelColor(fltk.BLUE)
	me.view.SetLabel(info)
	fltk.AddTimeout(5, func() { me.clearView() })
}
