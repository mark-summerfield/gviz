// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package gui

import _ "embed"

const (
	ButtonHeight       = 32
	ToolbuttonIconSize = 24
	LabelWidth         = 60
	ButtonWidth        = LabelWidth + (LabelWidth / 2)
	Pad                = 5
	Border             = 8
)

//go:embed question.svg
var questionSvg string
