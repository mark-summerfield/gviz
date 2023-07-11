// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

func (me *App) onViewZoomIn() {
	me.zoomView(1.1)
}

func (me *App) onViewZoomOut() {
	me.zoomView(0.9)
}

func (me *App) zoomView(amount float64) {
	me.zoomLevel *= amount
	me.mainVBox.Layout()
	me.onTextChanged(false)
}

func (me *App) onViewZoomRestore() {
	me.zoomLevel = 1
	me.onTextChanged(false)
}
