// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import _ "embed"

//go:embed Version.dat
var Version string

const (
	appName     = "Gviz"
	domain      = "qtrac.eu"
	linoWidth   = 32
	defaultText = "graph {\n    Gviz [shape=tab]\n}"
)

//go:embed images/icon.svg
var iconSvg string

//go:embed images/open.svg
var openSvg string

//go:embed images/save.svg
var saveSvg string

//go:embed images/zoom-in.svg
var zoomInSvg string

//go:embed images/zoom-original.svg
var zoomRestoreSvg string

//go:embed images/zoom-out.svg
var zoomOutSvg string

//go:embed images/dummy.png
var dummyPng []byte
