// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

//go:build !windows

package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/mark-summerfield/gviz/u"
	"github.com/pwiecz/go-fltk"
)

func aboutHtml() string {
	var year string
	y := time.Now().Year()
	if y == 2023 {
		year = fmt.Sprintf("%d", y)
	} else {
		year = fmt.Sprintf("2023-%d", y-2000)
	}
	distro := ""
	if out, err := exec.Command("lsb_release",
		"-ds").Output(); err == nil {
		distro = strings.TrimSpace(string(out))
	}
	var utsname syscall.Utsname
	_ = syscall.Uname(&utsname)
	if distro == "" {
		distro = u.Int8ToStr(utsname.Sysname[:]) + " " +
			u.Int8ToStr(utsname.Release[:])
	} else {
		distro += " " + u.Int8ToStr(utsname.Release[:])
	}
	return fmt.Sprintf(
		`<center><h3><font color=navy>%s v%s</font></h3></center>
<p><center><font color=navy>An application for editing and viewing GraphViz diagrams.</font></center></p>
<p><center>
<a
href="https://github.com/mark-summerfield/gviz">https://github.com/mark-summerfield/gviz</a>
</center></p>
<p><center>
<font color=green>Copyright © %s Mark Summerfield.<br>
All rights reserved.<br>
License: GPLv3.
</center></p>
<p><center><font color=#222>%s • %s/%s</font></center><br>
<center><font color=#222>go-fltk %s • FLTK
%s</font></center><br>
<center><font color=#222>%s</font></center></p>`,
		appName, Version, year, runtime.Version(), runtime.GOOS,
		runtime.GOARCH, fltk.GoVersion(), fltk.Version(), distro)
}
