// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

type menuItemDatum struct {
	text     string
	shortcut int
	method   func()
	divider  bool
}

type toolbuttonDatum struct {
	svg    string
	method func()
	tip    string
}

type shapeDatum struct {
	display string
	name    string
	svg     string
}
