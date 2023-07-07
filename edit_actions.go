// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "fmt"

func (me *App) onEditUndo() {
	if me.editor.Changed() != 0 {
		me.editor.Undo()
	}
}

func (me *App) onEditRedo() {
	if me.editor.Changed() != 0 {
		me.editor.Redo()
	}
}

func (me *App) onEditFind() {
	fmt.Println("onEditFind") // TODO
}

func (me *App) onEditFindAgain() {
	if me.findText == "" {
		me.onEditFind()
	} else {
		i := me.buffer.Search(me.editor.GetInsertPosition(), me.findText,
			false, me.findMatchCase)
		if i == -1 {
			me.onInfo(fmt.Sprintf("Didn't find %q searching forward.",
				me.findText))
		} else {
			span := len([]byte(me.findText))
			me.buffer.Select(i, i+span)
			me.editor.SetInsertPosition(i)
		}
	}
}
