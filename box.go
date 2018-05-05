// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

import (
	"h12.io/go-plplot/c"
)

func init() {
}

type Position int

type Axis struct {
	Position   Position
	NoTick     bool // implies NoLabel
	Tick       Tick
	NoLabel    bool
	Label      TickLabel
	AxisAtZero bool
	NoLine     bool
}

func (s Axis) Do() {
	if s.Label.IsTime() {
		if old := c.Gtimefmt(); old != s.Label.TimeFormat {
			defer c.Timefmt(old)
			s.Label.Do()
		}
	}

	if s.Position&TOP != 0 {
		topAxis(&s)
	}
	if s.Position&LEFT != 0 {
		leftAxis(&s)
	}
	if s.Position&RIGHT != 0 {
		rightAxis(&s)
	}
	if s.Position&BOTTOM != 0 {
		bottomAxis(&s)
	}
}

func (s Axis) Prepare() Doer {
	if s.Label.IsTime() && s.Label.TimeFormat == "" {
		s.Label.TimeFormat = DefaultTimeFormat
	}
	return s
}

/*
func NewAxis(ticks float64, nsub int) *Axis {
	return &Axis{
		Tick: Tick{
			MajorInterval: ticks,
			SubCount:      nsub,
		},
	}
}
*/

func (s *Axis) Option() string {
	opt := ``
	if s.AxisAtZero {
		opt += `a`
	}
	if !s.NoTick {
		opt += s.Tick.Option()
	}
	if !s.NoLabel {
		opt += s.Label.Option()
	}
	return opt
}

type Tick struct {
	MajorInterval float64
	SubCount      int // 0, automatic, -1 disabled
	Outward       bool
	Log           bool
	Grid          Grid
}

func (s *Tick) Option() string {
	opt := ``
	opt += `t`
	if s.SubCount != 1 {
		opt += `s`
	}
	if s.Outward {
		opt += `i`
	}
	if s.Log {
		opt += `l`
	}
	opt += s.Grid.Option()
	return opt
}

type Grid int

const (
	GRID_NONE Grid = iota
	GRID_MAJOR
	GRID_ALL
)

func (s Grid) Option() string {
	switch s {
	case GRID_MAJOR:
		return `g`
	case GRID_ALL:
		return `gh`
	}
	return ``
}

type TickLabel struct {
	Style      TickLabelStyle
	TimeFormat string
	Precision  int
}

type TickLabelStyle int

const (
	NUMERIC TickLabelStyle = iota
	DATETIME
	CUSTOM
	FIXEDPOINT
)

func (s TickLabel) Option() string {
	switch s.Style {
	case NUMERIC:
		return ""
	case DATETIME:
		return "d"
	case CUSTOM:
		return "o"
	case FIXEDPOINT:
		return "f"
	}
	return ""
}

func (s TickLabel) IsTime() bool {
	return s.Style == DATETIME
}

func (s TickLabel) Do() {
	if s.IsTime() {
		c.Timefmt(s.TimeFormat)
	}
}

func topAxis(style *Axis) {
	pre := `w`
	if !style.NoLine {
		pre = `c`
	}
	if !style.NoLabel {
		pre += `m`
	}
	c.Box(pre+style.Option(), style.Tick.MajorInterval, style.Tick.SubCount, ``, 0, 0)
}

func bottomAxis(style *Axis) {
	pre := `u`
	if !style.NoLine {
		pre = `b`
	}
	if !style.NoLabel {
		pre += `n`
	}
	c.Box(pre+style.Option(), style.Tick.MajorInterval, style.Tick.SubCount, ``, 0, 0)
}

func rightAxis(style *Axis) {
	pre := `w`
	if !style.NoLine {
		pre = `c`
	}
	if !style.NoLabel {
		pre += `m`
	}
	c.Box(``, 0, 0, pre+style.Option(), style.Tick.MajorInterval, style.Tick.SubCount)
}

func leftAxis(style *Axis) {
	pre := `u`
	if !style.NoLine {
		pre = `b`
	}
	if !style.NoLabel {
		pre += `n`
	}
	c.Box(``, 0, 0, pre+style.Option(), style.Tick.MajorInterval, style.Tick.SubCount)
}

/*
func Box(xticks, yticks float64, nxsub, nysub int) {
	xStyle := NewAxis(xticks, nxsub)
	yStyle := NewAxis(yticks, nysub)
	LeftAxis(yStyle)
	BottomAxis(xStyle)
	xStyle.Tick.Disabled = true
	yStyle.Tick.Disabled = true
	TopAxis(xStyle)
	RightAxis(yStyle)
}

func LogLogBox(subTick bool) {
	nsub := 1
	if subTick {
		nsub = 0
	}
	xStyle := NewAxis(1, nsub)
	yStyle := NewAxis(1, nsub)
	xStyle.Tick.Log = true
	yStyle.Tick.Log = true
	LeftAxis(yStyle)
	BottomAxis(xStyle)
	xStyle.Tick.Disabled = true
	yStyle.Tick.Disabled = true
	//TopAxis(xStyle)
	//RightAxis(yStyle)
}
*/
