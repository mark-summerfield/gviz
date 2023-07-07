// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"math"
	"strconv"
)

func (me *App) onInsertShape(shape string) {
	nodeId := me.getNextNodeId()
	text := ""
	if shape == "polygon" {
		text = fmt.Sprintf("n%d [shape=%s sides=5]", nodeId, shape)
	} else {
		text = fmt.Sprintf("n%d [shape=%s]", nodeId, shape)
	}
	me.editor.InsertText(text)
	me.onTextChanged(true)
}

func (me *App) getNextNodeId() int {
	if graph, err := me.getGraph(); err == nil {
		for i := 1; i < math.MaxInt; i++ {
			name := "n" + strconv.Itoa(i)
			if !graph.IsNode(name) {
				return i
			}
		}
	}
	return 0
}
