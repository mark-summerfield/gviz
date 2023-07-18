// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/awalterschulze/gographviz"
)

func (me *App) onInsertWord(word string) {
	text := strings.ReplaceAll(word, "&", "")
	tag := true
	switch text {
	case "bold tag":
		text = "<b></b>"
	case "font tag":
		text = "<font></font>"
	case "italic tag":
		text = "<i></i>"
	case "table tag":
		text = "<table></table>"
	case "tr (row) tag":
		text = "<tr></tr>"
	case "td (cell) tag":
		text = "<td></td>"
	default:
		tag = false
	}
	offset := 0
	if tag {
		offset = len(text) - strings.IndexByte(text, '/') + 1
	}
	i, j := me.buffer.GetSelectionPosition()
	if i < j {
		me.buffer.ReplaceSelection(text)
		me.editor.SetInsertPosition(i + len(text))
	} else {
		me.editor.InsertText(text)
	}
	if tag {
		me.editor.SetInsertPosition(me.editor.GetInsertPosition() - offset)
	}
	me.onTextChanged(true)
}

func (me *App) onInsertShape(shape string) {
	nodeId := me.getNextNodeId()
	text := ""
	if shape == "polygon" {
		text = fmt.Sprintf("n%d [shape=%s sides=5]", nodeId, shape)
	} else {
		text = fmt.Sprintf("n%d [shape=%s]", nodeId, shape)
	}
	i, j := me.buffer.GetSelectionPosition()
	if i < j {
		me.buffer.ReplaceSelection(text)
		me.editor.SetInsertPosition(i + len(text))
	} else {
		me.editor.InsertText(text)
	}
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

func (me *App) getGraph() (*gographviz.Graph, error) {
	graphAst, err := gographviz.ParseString(me.buffer.Text())
	if err != nil {
		return nil, err
	}
	graph := gographviz.NewGraph()
	if err = gographviz.Analyse(graphAst, graph); err != nil {
		return nil, err
	}
	return graph, nil
}
