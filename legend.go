// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

import (
	"github.com/hailiang/go-plplot/c"
)

const (
	LEFT Position = 1 << iota
	RIGHT
	TOP
	BOTTOM
	INSIDE
	OUTSIDE
	VIEWPORT
	SUBPAGE
)

type (
	LegendOption int
	LegendType   int
	LegendTypes  []LegendType
)

func (t LegendTypes) ToInt32() []int32 {
	r := make([]int32, len(t))
	for i := range t {
		r[i] = int32(t[i])
	}
	return r
}

const (
	LEGEND_NONE LegendType = 1 << iota
	LEGEND_COLOR_BOX
	LEGEND_LINE
	LEGEND_SYMBOL
)

const (
	LEGEND_TEXT_LEFT LegendOption = 0x10 << iota
	LEGEND_BACKGROUND
	LEGEND_BOUNDING_BOX
	LEGEND_ROW_MAJOR
)

type LegendSize struct {
	Width  float64
	Height float64
}

type LegendOffset struct {
	Left float64
	Top  float64
}

type BoundingBox struct {
	Color int
	Style int
}

type TextLegend struct {
	Offset        float64
	Scale         float64
	Spacing       float64
	Justification float64
	Colors        []int
	Texts         []string
}

type LegendText struct {
	Color int
	Text  string
}

type ColorBoxLegend struct {
	Colors     []int
	Patterns   []int
	Scales     []float64
	LineWidths []float64
}

type LineLegend struct {
	Colors []int
	Styles []int
	Widths []float64
}

type SymbolLegend struct {
	Colors  []int
	Scales  []float64
	Numbers []int
	Symbols []string
}

type Legend struct {
	Option        LegendOption
	Position      Position
	Offset        LegendOffset
	PlotWidth     float64
	BackColor     int
	Box           BoundingBox
	NRow, NColumn int
	OptionArray   LegendTypes
	Text          TextLegend
	ColorBox      ColorBoxLegend
	Line          LineLegend
	Symbol        SymbolLegend
}

func (l Legend) Do() {
	c.Legend(
		int(l.Option),
		int(l.Position),
		l.Offset.Left, l.Offset.Top,
		l.PlotWidth,
		l.BackColor,
		l.Box.Color,
		l.Box.Style,
		l.NRow, l.NColumn,
		len(l.OptionArray),
		l.OptionArray.ToInt32(),
		l.Text.Offset,
		l.Text.Scale,
		l.Text.Spacing,
		l.Text.Justification,
		int32s(l.Text.Colors),
		l.Text.Texts,
		int32s(l.ColorBox.Colors),
		int32s(l.ColorBox.Patterns),
		l.ColorBox.Scales,
		l.ColorBox.LineWidths,
		int32s(l.Line.Colors),
		int32s(l.Line.Styles),
		l.Line.Widths,
		int32s(l.Symbol.Colors),
		l.Symbol.Scales,
		int32s(l.Symbol.Numbers),
		l.Symbol.Symbols)
}

func int32s(s []int) []int32 {
	r := make([]int32, len(s))
	for i := range s {
		r[i] = int32(s[i])
	}
	return r
}
