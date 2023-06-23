// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

const (
	domain       = "qtrac.eu"
	appName      = "Dots"
	buttonHeight = 32
	labelWidth   = 60
	pad          = 5
)

//go:embed images/icon.svg
var iconSvg string

//go:embed images/open.svg
var openSvg string
