// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
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
	me.applySyntaxHighlighting()
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
	me.updateNextNodeId(graph)
	png, err := fltk.NewPngImageFromData(raw.Bytes())
	if err != nil {
		me.onError(err)
		return
	}
	me.updateView(png)
}

func (me *App) updateNextNodeId(graph *cgraph.Graph) {
	node := graph.FirstNode()
	for node != nil {
		name := node.Name()
		if strings.HasPrefix(name, "n") {
			if n, err := strconv.Atoi(name[1:]); err == nil {
				me.nextNodeId = u.Max(n+1, me.nextNodeId)
			}
		}
		node = graph.NextNode(node)
	}
}

func (me *App) updateView(png *fltk.PngImage) {
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
	fltk.AddTimeout(tinyTimeout, func() { me.view.Redraw() })
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
	me.view.SetLabel(err.Error())
}

func (me *App) onInfo(info string) {
	me.view.SetLabelColor(fltk.BLUE)
	me.view.SetLabel(info)
	fltk.AddTimeout(5, func() { me.clearView() })
}
