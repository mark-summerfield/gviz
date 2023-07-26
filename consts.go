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
