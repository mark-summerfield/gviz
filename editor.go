// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"bytes"
	"regexp"
	"unicode"

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
	if fltk.EventType() == fltk.KEY {
		ctrl := (fltk.EventState() & fltk.CTRL) != 0
		switch fltk.EventKey() {
		case fltk.ENTER_KEY:
			return me.onEditorEnter()
		case fltk.BACKSPACE:
			return me.onEditorBackspace(ctrl)
		case fltk.DELETE:
			return me.onEditorDelete(ctrl)
		}
	}
	return false
}

// Copy previous line's leading indentation if any.
func (me *App) onEditorEnter() bool {
	j := me.editor.GetInsertPosition()
	text := me.buffer.Text()
	if j < 0 || text == "" {
		return false
	}
	raw := []byte(text)
	insert := "\n"
	if i := bytes.LastIndexByte(raw[:j-1], '\n'); i > -1 {
		prev := raw[i+1 : j]
		for k := 0; k < len(prev); k++ {
			c := prev[k]
			if c == ' ' || c == '\t' {
				insert += string(c)
			} else {
				break
			}
		}
	}
	me.editor.InsertText(insert)
	me.dirty = true
	return true
}

// If backspace on 4+ spaces, unindent by 4 spaces; otherwise,
// if Ctrl+Backspace do delete prev word.
func (me *App) onEditorBackspace(ctrl bool) bool {
	j := me.editor.GetInsertPosition()
	text := me.buffer.Text()
	if j < 0 || text == "" {
		return false
	}
	raw := []byte(text)
	n := 3
	i := j - 1
	for ; n > 0 && i > 0 && raw[i] == ' '; i-- {
		n--
	}
	if n == 0 {
		me.buffer.Select(i, j)
		me.buffer.ReplaceSelection("")
		me.dirty = true
		return true
	}
	if i := bytes.LastIndexFunc(raw[:j-1], func(r rune) bool {
		return r != '_' && (unicode.IsSpace(r) || unicode.IsPunct(r) ||
			unicode.IsSymbol(r))
	}); i > -1 {
		me.buffer.Select(i+1, j)
		me.buffer.ReplaceSelection("")
		me.dirty = true
		return true
	}
	return false
}

// if Ctrl+Delete do delete next word.
func (me *App) onEditorDelete(ctrl bool) bool {
	i := me.editor.GetInsertPosition()
	text := me.buffer.Text()
	if i < 0 || text == "" {
		return false
	}
	raw := []byte(text)
	if j := bytes.IndexFunc(raw[i:], func(r rune) bool {
		return r != '_' && (unicode.IsSpace(r) || unicode.IsPunct(r) ||
			unicode.IsSymbol(r))
	}); j > -1 {
		j += i
		if j <= i {
			return false
		}
		me.buffer.Select(i, j)
		me.buffer.ReplaceSelection("")
		me.dirty = true
		return true
	}
	return false
}
