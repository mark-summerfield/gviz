// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
)

func (me *App) onInsertShape(shape string) {
	text := ""
	if shape == "polygon" {
		text = fmt.Sprintf("n%d [shape=%s sides=5]", me.nextNodeId, shape)
	} else {
		text = fmt.Sprintf("n%d [shape=%s]", me.nextNodeId, shape)
	}
	me.editor.InsertText(text)
	me.onTextChanged(true)
}
