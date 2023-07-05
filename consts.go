// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import _ "embed"

//go:embed Version.dat
var Version string

const (
	appName     = "Gviz"
	domain      = "qtrac.eu"
	description = "Edit, view, and export GraphViz diagrams."
	url         = "https://github.com/mark-summerfield/gviz"
	author      = "Mark Summerfield"
	linoWidth   = 32
	defaultText = "graph {\n    Gviz [shape=tab]\n}"
	tinyTimeout = 0.005

	boxShape     = "box"
	polygonShape = "polygon"
	ovalShape    = "oval"
	circleShape  = "circle"
)

//go:embed data/icon.svg
var iconSvg string

//go:embed data/open.svg
var openSvg string

//go:embed data/save.svg
var saveSvg string

//go:embed data/edit-copy.svg
var copySvg string

//go:embed data/edit-cut.svg
var cutSvg string

//go:embed data/edit-paste.svg
var pasteSvg string

//go:embed data/edit-redo.svg
var redoSvg string

//go:embed data/edit-undo.svg
var undoSvg string

//go:embed data/zoom-in.svg
var zoomInSvg string

//go:embed data/zoom-original.svg
var zoomRestoreSvg string

//go:embed data/zoom-out.svg
var zoomOutSvg string

//go:embed data/dummy.png
var dummyPng []byte

//go:embed data/help.html
var helpHtml string

//go:embed data/box.svg
var boxSvg string

//go:embed data/circle.svg
var circleSvg string

//go:embed data/oval.svg
var ovalSvg string

//go:embed data/polygon.svg
var polygonSvg string
