// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

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
