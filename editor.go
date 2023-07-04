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
			return false // TODO
			//fallthrough
		case fltk.TAB:
			fallthrough
		case fltk.ENTER_KEY:
			me.onEditorEdit(key)
			return true
		default:
			return false
		}
	}
	return false
}

func (me *App) onEditorEdit(key int) {
	j := me.editor.GetInsertPosition()
	text := me.buffer.Text()
	raw := []byte(text)
	switch key {
	case fltk.BACKSPACE:
		/* // TODO find & delete rune at j-1 but note j is byte count but we
			//  have runes...
		runes := []rune(text)
		if len(runes) > 0 {

		}
		r, size := utf8.DecodeLastRuneInString(text)
		if size > 0 {
			if r == ' ' || r == '\t' {
				fmt.Println("onEditorEdit BS over WS") // maybe unindent
			} // else
			if r != utf8.RuneError {
				me.buffer.SetText(string(raw[:len(raw)-size]))
				me.editor.SetInsertPosition(j - size)
				me.dirty = true
			}
		}
		*/
	case fltk.TAB:
		newRaw := raw[:j]
		newRaw = append(newRaw, '\t')
		newRaw = append(newRaw, raw[j:]...)
		me.buffer.SetText(string(newRaw))
		me.editor.SetInsertPosition(j + 1)
		me.dirty = true
	case fltk.ENTER_KEY:
		newRaw := raw[:j]
		// TODO add copy of prev line's WS (if any)
		newRaw = append(newRaw, '\n')
		newRaw = append(newRaw, raw[j:]...)
		me.buffer.SetText(string(newRaw))
		me.editor.SetInsertPosition(j + 1)
		me.dirty = true
	}
}

/* // TODO on '\n' above
if changed {
	j := me.editor.GetInsertPosition()
	if j > -1 && text[j] == '\n' {
		i := strings.LastIndexByte(text[:j-1], '\n')
		if i > -1 {
			if prevLine := text[i:j]; prevLine != "" {
				// TODO look for ws & if found replace text[i]'s \n with
				// same as first found & add rest
				// for k := i; k < j;
			}
		}
	}
}

*/
