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
)

//go:embed images/icon.svg
var iconSvg string

//go:embed images/open.svg
var openSvg string

//go:embed images/save.svg
var saveSvg string

//go:embed images/edit-copy.svg
var copySvg string

//go:embed images/edit-cut.svg
var cutSvg string

//go:embed images/edit-paste.svg
var pasteSvg string

//go:embed images/edit-redo.svg
var redoSvg string

//go:embed images/edit-undo.svg
var undoSvg string

//go:embed images/zoom-in.svg
var zoomInSvg string

//go:embed images/zoom-original.svg
var zoomRestoreSvg string

//go:embed images/zoom-out.svg
var zoomOutSvg string

//go:embed images/dummy.png
var dummyPng []byte
