// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
)

func (me *App) onInsertShape(shape string) {
	me.editor.InsertText(fmt.Sprintf("n%d [shape=%s]", me.nextNodeId,
		shape))
	me.onTextChanged(true)
}
