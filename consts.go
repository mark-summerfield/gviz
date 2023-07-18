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
	htmlwords = []string{"&bold tag", "b&gcolor=", "&colspan=",
		"&font tag", "&italic tag", "r&owspan=", "&table tag",
		"t&r (row) tag", "t&d (cell) tag"}

	shapeData = []shapeDatum{
		{"&Box (rectangle)", "box", boxSvg},
		{"&Circle", "circle", circleSvg},
		{"C&ylinder", "cylinder", cylinderSvg},
		{"&Diamond", "diamond", diamondSvg},
		{"&Folder", "folder", folderSvg},
		{"Ho&use", "house", houseSvg},
		{"&Note", "note", noteSvg},
		{"&Oval (ellipse)", "oval", ovalSvg},
		{"P&lain", "plain", ""},
		{"&Polygon", "polygon", polygonSvg},
		{"&Tab", "tab", tabSvg},
		{"Trape&zium", "trapezium", trapeziumSvg}}
	extraShapeData = []shapeDatum{
		{"Assembl&y", "assembly", assemblySvg},
		{"&Box 3D", "box3d", box3dSvg},
		{"CDS", "cds", cdsSvg},
		{"&Component", "component", componentSvg},
		{"&Egg", "egg", eggSvg},
		{"&Fivepoverhang", "fivepoverhang", fivepoverhangSvg},
		{"&Insulator", "insulator", insulatorSvg},
		{"L&arrow", "larrow", larrowSvg},
		{"Lpromoter", "lpromoter", lpromoterSvg},
		{"&Noverhang", "noverhang", noverhangSvg},
		{"Primersite", "primersite", primersiteSvg},
		{"Promoter", "promoter", promoterSvg},
		{"Proteasesite", "proteasesite", proteasesiteSvg},
		{"&Proteinstab", "proteinstab", proteinstabSvg},
		{"Rarro&w", "rarrow", rarrowSvg},
		{"Restrictionsite", "restrictionsite", restrictionsiteSvg},
		{"Rib&osite", "ribosite", ribositeSvg},
		{"&Rnastab", "rnastab", rnastabSvg},
		{"Rpro&moter", "rpromoter", rpromoterSvg},
		{"Si&gnature", "signature", signatureSvg},
		{"&Star", "star", starSvg},
		{"&Terminator", "terminator", terminatorSvg},
		{"T&hreepoverhang", "threepoverhang", threepoverhangSvg},
		{"Un&derline", "underline", ""},
		{"&UTR", "utr", utrSvg}}
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

//go:embed data/promoter.svg
var promoterSvg string

//go:embed data/terminator.svg
var terminatorSvg string

//go:embed data/utr.svg
var utrSvg string

//go:embed data/assembly.svg
var assemblySvg string

//go:embed data/fivepoverhang.svg
var fivepoverhangSvg string

//go:embed data/insulator.svg
var insulatorSvg string

//go:embed data/larrow.svg
var larrowSvg string

//go:embed data/lpromoter.svg
var lpromoterSvg string

//go:embed data/noverhang.svg
var noverhangSvg string

//go:embed data/proteasesite.svg
var proteasesiteSvg string

//go:embed data/threepoverhang.svg
var threepoverhangSvg string

//go:embed data/proteinstab.svg
var proteinstabSvg string

//go:embed data/rarrow.svg
var rarrowSvg string

//go:embed data/restrictionsite.svg
var restrictionsiteSvg string

//go:embed data/ribosite.svg
var ribositeSvg string

//go:embed data/rnastab.svg
var rnastabSvg string

//go:embed data/rpromoter.svg
var rpromoterSvg string

//go:embed data/signature.svg
var signatureSvg string
