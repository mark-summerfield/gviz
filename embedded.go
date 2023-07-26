// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "embed"

//go:embed images/*
//go:embed html/*.html
var embedded embed.FS

func getEmbStr(name string) string {
	raw, err := embedded.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

func getEmbRaw(name string) []byte {
	raw, err := embedded.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return raw
}

const (
	helpHtml = "html/help.html"

	iconSvg            = "images/icon.svg"
	openSvg            = "images/open.svg"
	saveSvg            = "images/save.svg"
	undoSvg            = "images/edit-undo.svg"
	redoSvg            = "images/edit-redo.svg"
	copySvg            = "images/edit-copy.svg"
	cutSvg             = "images/edit-cut.svg"
	pasteSvg           = "images/edit-paste.svg"
	findSvg            = "images/edit-find.svg"
	findAgainSvg       = "images/edit-find-again.svg"
	zoomInSvg          = "images/zoom-in.svg"
	zoomRestoreSvg     = "images/zoom-original.svg"
	zoomOutSvg         = "images/zoom-out.svg"
	dummyPng           = "images/dummy.png"
	boxSvg             = "images/box.svg"
	box3dSvg           = "images/box3d.svg"
	circleSvg          = "images/circle.svg"
	ovalSvg            = "images/oval.svg"
	polygonSvg         = "images/polygon.svg"
	cylinderSvg        = "images/cylinder.svg"
	diamondSvg         = "images/diamond.svg"
	eggSvg             = "images/egg.svg"
	folderSvg          = "images/folder.svg"
	houseSvg           = "images/house.svg"
	noteSvg            = "images/note.svg"
	starSvg            = "images/star.svg"
	tabSvg             = "images/tab.svg"
	trapeziumSvg       = "images/trapezium.svg"
	cdsSvg             = "images/cds.svg"
	componentSvg       = "images/component.svg"
	primersiteSvg      = "images/primersite.svg"
	promoterSvg        = "images/promoter.svg"
	terminatorSvg      = "images/terminator.svg"
	utrSvg             = "images/utr.svg"
	assemblySvg        = "images/assembly.svg"
	fivepoverhangSvg   = "images/fivepoverhang.svg"
	insulatorSvg       = "images/insulator.svg"
	larrowSvg          = "images/larrow.svg"
	lpromoterSvg       = "images/lpromoter.svg"
	noverhangSvg       = "images/noverhang.svg"
	proteasesiteSvg    = "images/proteasesite.svg"
	threepoverhangSvg  = "images/threepoverhang.svg"
	proteinstabSvg     = "images/proteinstab.svg"
	rarrowSvg          = "images/rarrow.svg"
	restrictionsiteSvg = "images/restrictionsite.svg"
	ribositeSvg        = "images/ribosite.svg"
	rnastabSvg         = "images/rnastab.svg"
	rpromoterSvg       = "images/rpromoter.svg"
	signatureSvg       = "images/signature.svg"
)
