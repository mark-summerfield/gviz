// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/pwiecz/go-fltk"
)

func (me *App) initializeEditor() {
	me.buffer = fltk.NewTextBuffer()
	me.buffer.SetText(defaultText)
	me.editor.SetBuffer(me.buffer)
	me.editor.SetTextFont(fltk.COURIER)
	me.editor.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	me.highlightBuffer = fltk.NewTextBuffer()
	me.makeTextStyles()
	me.editor.SetHighlightData(me.highlightBuffer, me.textStyles)
	me.editor.SetCallback(func() {
		me.onTextChanged(true)
	})
	me.onTextChanged(false)
	me.onLinosChange()
}

func (me *App) makeTextStyles() {
	font := fltk.COURIER
	keywordFont := fltk.COURIER_BOLD
	size := 13
	me.textStyles = []fltk.StyleTableEntry{
		{Color: fltk.BLACK, Font: font, Size: size},        // A default
		{Color: fltk.BLACK, Font: keywordFont, Size: size}, // B keywords
		{Color: fltk.BLUE, Font: font, Size: size},         // C compass
	}

}

func (me *App) applySyntaxHighlighting() {
	rx := regexp.MustCompile(
		`(\b(:?strict|graph|digraph|node|edge|subgraph)\b)|` +
			`(\b(:?n|ne|e|se|s|sw|w|nw|c)\b)`)
	raw := []byte(me.buffer.Text())
	highlight := bytes.Repeat([]byte{'A'}, len(raw))
	for i, subIndexes := range rx.FindAllSubmatchIndex(raw, -1) {
		fmt.Println(i, subIndexes)
	}
	me.highlightBuffer.SetText(string(highlight))
}

//func maybeHighlight(highlight []byte, style byte, indexes []int) {
//	fmt.Println(style, indexes)
//}
