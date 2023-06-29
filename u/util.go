// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package u

import (
	"os"
	"path/filepath"
)

func GetPath(filename string) string {
	if filename != "" {
		return filepath.Dir(filename)
	} else {
		path, err := os.UserHomeDir()
		if err == nil {
			return path
		}
		return "./"
	}
}

func Int8ToStr(raw []int8) string {
	data := make([]byte, 0, len(raw))
	for _, i := range raw {
		if i == 0 {
			break
		}
		data = append(data, byte(i))
	}
	return string(data)
}

// TODO drop once go 1.21 in use
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
