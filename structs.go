// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

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
