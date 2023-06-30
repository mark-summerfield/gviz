// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

//go:build !windows

package gui

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/mark-summerfield/gviz/u"
	"github.com/pwiecz/go-fltk"
)

func DescHtml(appName, version, desc, url, author, year string) string {
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
<p><center><font color=navy>%s</font></center></p>
<p><center><a href="%s">%s</a></center></p>
<p><center>
<font color=green>Copyright © %s %s.<br>
All rights reserved.<br>
License: GPLv3.
</center></p>
<p><center><font color=#222>%s • %s/%s</font></center><br>
<center><font color=#222>go-fltk %s • FLTK
%s</font></center><br>
<center><font color=#222>%s</font></center></p>`,
		appName, version, desc, url, url, author, year, runtime.Version(),
		runtime.GOOS, runtime.GOARCH, fltk.GoVersion(), fltk.Version(),
		distro)
}
