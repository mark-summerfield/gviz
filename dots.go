// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
	"fmt"
)

//go:embed Version.dat
var Version string

func main() {
	fmt.Printf("Hello dots v%s\n", Version)
}
