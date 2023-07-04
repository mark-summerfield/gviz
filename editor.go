// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"bytes"
	"regexp"

	"github.com/pwiecz/go-fltk"
)

func (me *App) initializeEditor() {
	me.buffer = fltk.NewTextBuffer()
	me.buffer.SetText(defaultText)
	me.editor.SetBuffer(me.buffer)
	me.editor.SetTextFont(fltk.COURIER)
	me.editor.SetTextSize(me.config.TextSize)
	me.editor.SetLinenumberSize(me.config.TextSize - 1)
	me.editor.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	me.editor.SetEventHandler(me.onEditorEvent)
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
	roman := fltk.COURIER
	bold := fltk.COURIER_BOLD
	italic := fltk.COURIER_ITALIC
	size := me.config.TextSize
	navy := fltk.ColorFromRgb(0, 0, 0x80)
	me.textStyles = []fltk.StyleTableEntry{
		{Color: fltk.BLACK, Font: roman, Size: size},        // A default
		{Color: navy, Font: bold, Size: size},               // B keyword
		{Color: fltk.DARK_RED, Font: roman, Size: size},     // C compass
		{Color: fltk.DARK_GREEN, Font: italic, Size: size},  // D comment
		{Color: fltk.DARK_YELLOW, Font: roman, Size: size},  // E string
		{Color: fltk.DARK_MAGENTA, Font: roman, Size: size}, // F attrib
		{Color: fltk.BLUE, Font: bold, Size: size},          // G arrow/line
		{Color: fltk.DARK_CYAN, Font: roman, Size: size},    // H punct
	}

}

func (me *App) applySyntaxHighlighting() {
	rx := regexp.MustCompile(
		`(\b(:?strict|graph|digraph|node|edge|subgraph)\b)` + // B keyword
			`|(\b(:?n|ne|e|se|s|sw|w|nw|c)\b)` + // C compass
			`|(#.*?)\n` + // D comment
			`|("[^"]+?")` + // E string
			`|(\w+=)` + // F attrib name
			`|(-[->])` + // G arrow or line
			`|([[:punct:]]+)`) // H punct
	raw := []byte(me.buffer.Text())
	highlight := bytes.Repeat([]byte{'A'}, len(raw))
	for _, subIndexes := range rx.FindAllSubmatchIndex(raw, -1) {
		maybeHighlight(highlight, 'H', 18, subIndexes)
		maybeHighlight(highlight, 'G', 16, subIndexes)
		maybeHighlight(highlight, 'F', 14, subIndexes)
		maybeHighlight(highlight, 'C', 6, subIndexes)
		maybeHighlight(highlight, 'B', 2, subIndexes)
		maybeHighlight(highlight, 'E', 12, subIndexes) // must be pre-last
		maybeHighlight(highlight, 'D', 10, subIndexes) // must be last
	}
	me.highlightBuffer.SetText(string(highlight))
}

func maybeHighlight(highlight []byte, style byte, j int, indexes []int) {
	for k := indexes[j]; k < indexes[j+1]; k++ {
		highlight[k] = style
	}
}

func (me *App) onEditorEvent(event fltk.Event) bool {
	key := fltk.EventKey()
	switch fltk.EventType() {
	case fltk.KEY:
		switch key {
		case fltk.BACKSPACE:
			fallthrough
		case fltk.ENTER_KEY:
			return me.onEditorEdit(key)
		default:
			return false
		}
	}
	return false
}

func (me *App) onEditorEdit(key int) bool {
	j := me.editor.GetInsertPosition()
	text := me.buffer.Text()
	if j < 0 || text == "" {
		return false
	}
	raw := []byte(text)
	switch key {
	case fltk.BACKSPACE:
		j--
		for i, r := range text {
			if i == j && r == ' ' {
				k := i
				n := 3
				for ; n > 0 && k > 0 && raw[k] == ' '; k-- {
					n--
				}
				if k < i && n == 0 { // unindent
					newRaw := raw[:k]
					newRaw = append(newRaw, raw[j:]...)
					me.buffer.SetText(string(newRaw))
					me.editor.SetInsertPosition(k)
					me.dirty = true
					return true
				}
			} else if i > j {
				return false
			}
		}
		return false
	case fltk.ENTER_KEY: // TODO
		n := 1
		newRaw := raw[:j]
		if i := bytes.LastIndexByte(raw[:j-1], '\n'); i > -1 {
			prev := raw[i+1 : j]
			for k := 0; k < len(prev); k++ {
				if prev[k] == ' ' || prev[k] == '\t' {
					newRaw = append(newRaw, prev[k])
					n++
				} else {
					break
				}
			}
		}
		newRaw = append(newRaw, '\n')
		newRaw = append(newRaw, raw[j:]...)
		me.buffer.SetText(string(newRaw))
		me.editor.SetInsertPosition(j + n)
		me.dirty = true
	}
	return true
}
