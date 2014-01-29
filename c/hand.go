// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c

/*
#include <plplot/plplot.h>
*/
import "C"

import (
	"unsafe"
)

var (
	Logging bool
)

func Setenv(key, value string) {
	env_ := S(key + "=" + value)
	defer freeS(env_)
	C.putenv(env_)
}

type ContourTransform func(x, y float64) (float64, float64)

type globalState struct {
	timefmt_fmt string
	sub_nx      int
	sub_ny      int
	width_width float64
	lsty_n      int
}

var gs = globalState{
	sub_nx:      1,
	sub_ny:      1,
	width_width: 1,
	lsty_n:      1,
}

func Timefmt(fmt string) {
	gs.timefmt_fmt = fmt
	_fmt := C.CString(fmt)
	defer C.free(unsafe.Pointer(_fmt))
	C.c_pltimefmt(_fmt)
}

func Gtimefmt() (fmt string) {
	return gs.timefmt_fmt
}

func Ssub(nx int, ny int) {
	gs.sub_nx = nx
	gs.sub_ny = ny
	_nx := (C.PLINT)(nx)
	_ny := (C.PLINT)(ny)
	C.c_plssub(_nx, _ny)
}

func Gsub() (nx int, ny int) {
	return gs.sub_nx, gs.sub_ny
}

func Width(width float64) {
	gs.width_width = width
	_width := (C.PLFLT)(width)
	C.c_plwidth(_width)
}

func Gwidth() (width float64) {
	return gs.width_width
}

func Lsty(lin int) {
	gs.lsty_n = lin
	_lin := (C.PLINT)(lin)
	C.c_pllsty(_lin)
}

func Glsty() (n int) {
	return gs.lsty_n
}
