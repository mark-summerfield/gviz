// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"math"
	"strconv"
)

func (me *App) onInsertShape(shape string) {
	me.updateNextNodeId()
	text := ""
	if shape == "polygon" {
		text = fmt.Sprintf("n%d [shape=%s sides=5]", me.nextNodeId, shape)
	} else {
		text = fmt.Sprintf("n%d [shape=%s]", me.nextNodeId, shape)
	}
	me.editor.InsertText(text)
	me.onTextChanged(true)
}

func (me *App) updateNextNodeId() {
	if graph, err := me.getGraph(); err == nil {
		for i := me.nextNodeId; i < math.MaxInt; i++ {
			name := "n" + strconv.Itoa(i)
			if !graph.IsNode(name) {
				me.nextNodeId = i
				break
			}
		}
	}
}
