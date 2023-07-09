// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/gviz/gui"
	"github.com/mark-summerfield/gviz/u"
	"github.com/pwiecz/go-fltk"
)

func (me *App) onEvent(event fltk.Event) bool {
	key := fltk.EventKey()
	switch fltk.EventType() {
	case fltk.SHOW:
		me.onToggleStandardToolbar(false)
		me.onToggleExtraShapesToolbar(false)
		me.layout()
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
	if err := me.saveGraph(temppng); err != nil {
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

func (me *App) onToggleStandardToolbar(refresh bool) {
	if me.config.ShowStandard {
		me.standardToolbar.Show()
	} else {
		me.standardToolbar.Hide()
	}
	if refresh {
		me.layout()
	}
}

func (me *App) onToggleExtraShapesToolbar(refresh bool) {
	if me.config.ShowExtraShapes {
		me.extraShapesToolbar.Show()
	} else {
		me.extraShapesToolbar.Hide()
	}
	if refresh {
		me.layout()
	}
}

func (me *App) layout() {
	me.mainVBox.Resize(me.mainVBox.X(), me.mainVBox.Y(), me.mainVBox.W(),
		me.mainVBox.H())
	me.mainVBox.Redraw()
}

func (me *App) onError(err error) {
	rx := regexp.MustCompile(`Pos\s*[(]\s*offset\s*=\s*\d+,` +
		`\s*line\s*=\s*(\d+),\s*column\s*=(\d+)\s*[)]\s*(.*)$`)
	me.view.SetLabelColor(fltk.RED)
	text := err.Error()
	if matches := rx.FindAllStringSubmatch(text, -1); len(matches) == 1 {
		match := matches[0]
		if len(match) == 4 {
			lino := match[1]
			column := match[2]
			message := strings.TrimLeft(match[3], ", ")
			text = fmt.Sprintf("#%s:%s: %s", lino, column, message)
		}
	}
	text = gong.Wrapped(text, 40)
	if !strings.HasSuffix(text, ".") {
		text += "."
	}
	me.view.SetLabel(text)
}

func (me *App) onInfo(info string) {
	me.view.SetLabelColor(fltk.BLUE)
	me.view.SetLabel(info)
	fltk.AddTimeout(5, func() { me.clearView() })
}
