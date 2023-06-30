// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

//go:build windows

package gui

import (
	"fmt"
	"runtime"

	"github.com/gonutz/w32/v2"
	"github.com/pwiecz/go-fltk"
)

func DescHtml(appName, version, desc, url, author, year string) string {
	info := w32.RtlGetVersion()
	distro := fmt.Sprintf("Windows %d.%d", info.MajorVersion,
		info.MinorVersion)
	return fmt.Sprintf(
		`<center><h3><font color=navy>%s v%s</font></h3></center>
<p><center><font color=navy>%s</font></center></p>
<p><center> <a href="%s">%s</a></center></p>
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
