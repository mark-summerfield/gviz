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
	me.editor.InsertText(strings.ReplaceAll(word, "&", ""))
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
