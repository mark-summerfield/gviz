// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import _ "embed"

//go:embed Version.dat
var Version string

const (
	appName            = "Dots"
	domain             = "qtrac.eu"
	buttonHeight       = 32
	toolbuttonIconSize = 24
	labelWidth         = 60
	pad                = 5
	border             = 8

	defaultText = `graph graphname {
    Dots [shape=tab]
}`
)

//go:embed images/icon.svg
var iconSvg string

//go:embed images/open.svg
var openSvg string
