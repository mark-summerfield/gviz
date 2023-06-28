// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import _ "embed"

//go:embed Version.dat
var Version string

const (
	appName            = "Gviz"
	domain             = "qtrac.eu"
	buttonHeight       = 32
	toolbuttonIconSize = 24
	labelWidth         = 60
	buttonWidth        = labelWidth + (labelWidth / 2)
	pad                = 5
	border             = 8
	defaultText        = "graph {\n    Gviz [shape=tab]\n}"
)

//go:embed images/icon.svg
var iconSvg string

//go:embed images/open.svg
var openSvg string

//go:embed images/save.svg
var saveSvg string

//go:embed images/close.svg
var closeSvg string

//go:embed images/ok.svg
var okSvg string

//go:embed images/question.svg
var questionSvg string

//go:embed images/dummy.png
var dummyPng []byte
