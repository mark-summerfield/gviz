// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "fmt"

func (me *App) onEditUndo() {
	if me.editor != nil {
		me.editor.Undo()
		me.applySyntaxHighlighting()
	}
}

func (me *App) onEditRedo() {
	if me.editor != nil {
		me.editor.Redo()
		me.applySyntaxHighlighting()
	}
}

func (me *App) onEditCopy() {
	if me.editor != nil {
		me.editor.Copy()
	}
}

func (me *App) onEditCut() {
	if me.editor != nil {
		me.editor.Cut()
	}
}

func (me *App) onEditPaste() {
	if me.editor != nil {
		me.editor.Paste()
	}
}

func (me *App) onEditFind() {
	fmt.Println("onEditFind") // TODO
}

func (me *App) onEditFindAgain() {
	if me.findText == "" {
		me.onEditFind()
	} else {
		i := me.buffer.Search(me.editor.GetInsertPosition()+1, me.findText,
			false, me.findMatchCase)
		if i == -1 {
			me.onInfo(fmt.Sprintf("Didn't find %q searching forward.",
				me.findText))
		} else {
			span := len(me.findText) // byte count is what we want
			me.buffer.Select(i, i+span)
			me.editor.SetInsertPosition(i)
		}
	}
}
