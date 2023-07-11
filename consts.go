// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import _ "embed"

//go:embed Version.dat
var Version string

const (
	appName        = "Gviz"
	domain         = "qtrac.eu"
	description    = "Edit, view, and export GraphViz diagrams."
	url            = "https://github.com/mark-summerfield/gviz"
	author         = "Mark Summerfield"
	maxRecentFiles = 9
	linoWidth      = 32
	defaultText    = "graph {\n    Gviz [shape=tab]\n}"
	tinyTimeout    = 0.005
	smallTimeout   = 0.1
	dotExe         = "dot"
	rowHeight      = 32
	colWidth       = 60
)

var (
	attributes = []string{"&color=", "&fillcolor=", "&label=", "&style="}
	keywords   = []string{"&bold", "&dashed", "d&otted", "&edge", "f&alse",
		"&filled", "&invis", "&node", "&rounded", "&solid", "s&ubgraph",
		"&true"}

	shapeData = []shapeDatum{
		{"&Box (rectangle)", "box", boxSvg},
		{"Box &3D", "box3d", box3dSvg},
		{"&Circle", "circle", circleSvg},
		{"C&ylinder", "cylinder", cylinderSvg},
		{"&Diamond", "diamond", diamondSvg},
		{"&Egg", "egg", eggSvg},
		{"&Folder", "folder", folderSvg},
		{"&House", "house", houseSvg},
		{"&Note", "note", noteSvg},
		{"&Oval (ellipse)", "oval", ovalSvg},
		{"P&lain", "plain", ""},
		{"&Polygon", "polygon", polygonSvg},
		{"&Star", "star", starSvg},
		{"&Tab", "tab", tabSvg},
		{"&Underline", "underline", ""},
		{"Trape&zium", "trapezium", trapeziumSvg}}
	extraShapeData = []shapeDatum{
		{"&CDS", "cds", cdsSvg},
		{"C&omponent", "component", componentSvg},
		{"&Primersite", "primersite", primersiteSvg},
		// TODO icons & remaining extra chapes
		{"P&romoter", "promoter", ""},
		{"&Terminator", "terminator", ""},
		{"&UTR", "utr", ""}}
)

//go:embed data/icon.svg
var iconSvg string

//go:embed data/open.svg
var openSvg string

//go:embed data/save.svg
var saveSvg string

//go:embed data/edit-undo.svg
var undoSvg string

//go:embed data/edit-redo.svg
var redoSvg string

//go:embed data/edit-copy.svg
var copySvg string

//go:embed data/edit-cut.svg
var cutSvg string

//go:embed data/edit-paste.svg
var pasteSvg string

//go:embed data/edit-find.svg
var findSvg string

//go:embed data/edit-find-again.svg
var findAgainSvg string

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

//go:embed data/box3d.svg
var box3dSvg string

//go:embed data/circle.svg
var circleSvg string

//go:embed data/oval.svg
var ovalSvg string

//go:embed data/polygon.svg
var polygonSvg string

//go:embed data/cylinder.svg
var cylinderSvg string

//go:embed data/diamond.svg
var diamondSvg string

//go:embed data/egg.svg
var eggSvg string

//go:embed data/folder.svg
var folderSvg string

//go:embed data/house.svg
var houseSvg string

//go:embed data/note.svg
var noteSvg string

//go:embed data/star.svg
var starSvg string

//go:embed data/tab.svg
var tabSvg string

//go:embed data/trapezium.svg
var trapeziumSvg string

//go:embed data/cds.svg
var cdsSvg string

//go:embed data/component.svg
var componentSvg string

//go:embed data/primersite.svg
var primersiteSvg string
