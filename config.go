// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/mark-summerfield/gong"
	"github.com/pwiecz/go-fltk"
)

type Config struct {
	filename               string
	X                      int
	Y                      int
	Width                  int
	Height                 int
	Scale                  float32
	TextSize               int
	ViewOnLeft             bool
	Linos                  bool
	LastFile               string
	AutoFormat             bool
	ShowStandardToolbar    bool
	ShowExtraShapesToolbar bool
	RecentFiles            []string
}

func newConfig() *Config {
	filename, found := gong.GetIniFile(domain, appName)
	config := &Config{filename: filename, X: -1, Width: 800, Height: 600,
		Scale: 1.0, TextSize: 14, Linos: true, ShowStandardToolbar: true}
	if found {
		cfg, err := ini.Load(filename)
		if err != nil {
			fmt.Println("newConfig #1", filename, err)
		} else {
			err = cfg.MapTo(config)
			if err != nil {
				fmt.Println("newConfig #2", filename, err)
			} else {
				_, _, width, height := fltk.ScreenWorkArea(0)
				if config.Width < 100 || config.Width > width {
					config.Width = 800
				}
				if config.Height < 100 || config.Height > height {
					config.Height = 600
				}
				if config.Scale < 0.5 || config.Scale > 5 {
					config.Scale = 1
				}
				if config.TextSize < 10 ||
					config.TextSize > 20 {
					config.TextSize = 14
				}
			}
			recentFiles := make([]string, 0, len(config.RecentFiles))
			for _, filename := range config.RecentFiles {
				if gong.FileExists(filename) {
					recentFiles = append(recentFiles, filename)
				}
			}
			config.RecentFiles = recentFiles
		}
	}
	return config
}

func (me *Config) save() {
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, me)
	if err != nil {
		fmt.Println("save #1", me.filename, err)
	} else {
		dir := filepath.Dir(me.filename)
		if dir != "." {
			if !gong.PathExists(dir) {
				_ = os.MkdirAll(dir, fs.ModePerm)
			}
		}
		err := cfg.SaveTo(me.filename)
		if err != nil {
			fmt.Println("save #2", me.filename, err)
		}
	}
}

func (me *Config) maybeAddRecentFile(filename string) {
	for i, name := range me.RecentFiles {
		if name == filename { // cut out since we'll add to front
			me.RecentFiles = append(me.RecentFiles[:i],
				me.RecentFiles[i+1:]...)
			break
		}
	}
	me.RecentFiles = append([]string{filename}, me.RecentFiles...)
	if len(me.RecentFiles) > maxRecentFiles {
		me.RecentFiles = me.RecentFiles[:maxRecentFiles]
	}
}
