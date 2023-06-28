// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"bytes"
	"fmt"
	"math"

	"github.com/goccy/go-graphviz"
	"github.com/mark-summerfield/gong"
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
		case fltk.HELP, fltk.F1:
			return true // ignore
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
	text := me.buffer.Text()
	if text == "" {
		me.onError(fmt.Errorf("Need image data, e.g.\n%s", defaultText))
		return
	}
	graph, err := graphviz.ParseBytes([]byte(text))
	if err != nil {
		me.onError(err)
		return
	}
	gv := graphviz.New()
	var raw bytes.Buffer // Tried SVG but text doesn't appear
	if err = gv.Render(graph, graphviz.PNG, &raw); err != nil {
		me.onError(err)
		return
	}
	png, err := fltk.NewPngImageFromData(raw.Bytes())
	if err != nil {
		me.onError(err)
		return
	}
	me.clearView()
	if !gong.IsRealClose(1.0, me.zoomLevel) {
		w := int(math.Round(float64(png.W()) * me.zoomLevel))
		h := int(math.Round(float64(png.H()) * me.zoomLevel))
		png.Scale(w, h, true, true)
	}
	if me.view.W() < max(me.scroll.W(), png.W()) ||
		me.view.H() < max(me.scroll.H(), png.H()) {
		me.view.Resize(0, 0, png.W()+border, png.H()+border)
	}
	me.view.SetImage(png)
}

func (me *App) clearView() {
	me.view.SetLabelColor(fltk.BLACK)
	me.view.SetLabel("")
}

func (me *App) onError(err error) {
	me.view.SetLabelColor(fltk.RED)
	me.view.SetLabel(err.Error())
}

func (me *App) onInfo(info string) {
	me.view.SetLabelColor(fltk.BLUE)
	me.view.SetLabel(info)
	fltk.AddTimeout(5, func() { me.clearView() })
}
