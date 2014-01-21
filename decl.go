// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Go-plplot is a Go binding to PLPlot scientific plotting library. Besides a thin
wrapper to PLPlot C functions, it also provides a declarative layer.

TODO:
1. Continue plot files one by one.
2. Error handling plsabort, plsexit

*/
package plplot

const (
	// It doesn't really matter what DPI value is because we are using vector
	// graphics, so it is just used to calculate the conversion between points
	// and inches.
	DPI = 72
)

var (
	DefaultFile           = File{"newplot.pdf"}
	DefaultPageConfig     = PageConfig{Width: 300, Height: 300}
	DefaultBackColor      = BackColor{255, 255, 255}
	DefaultFontTable      = FontTable{}
	DefaultForeColor      = ForeColor{0, 0, 0}
	DefaultLineWidth      = LineWidth{1}
	DefaultLineStyle      = LineStyle{1}
	DefaultViewport       = Viewport{XMin: 0.1, XMax: 0.9, YMin: 0.1, YMax: 0.9}
	DefaultViewportNoAxis = Viewport{XMin: 0, XMax: 1, YMin: 0, YMax: 1}
	DefaultFontSize       = FontSize{1}
	DefaultTimeFormat     = "%Y-%m-%d"
	DefaultOffset         = Offset{0, 0, 0, 0}
	DefaultAxis           = Block{
		Axis{Position: LEFT | BOTTOM},
		Axis{Position: TOP | RIGHT, NoTick: true},
	}

	USLetter = PageConfig{Width: round(DPI * 8.5), Height: round(DPI * 11)}
	A4       = PageConfig{Width: round(DPI * 8.3), Height: round(DPI * 11.7)}
)

type Plot Doers

type Subpage Doers

type Block Doers

type Init struct{}

type End struct{}

type File struct {
	Path string
}

type PageConfig struct {
	Width, Height, Left, Top int
	NSubpageX, NSubpageY     int
}

type Color struct {
	R, G, B byte
}

type BackColor Color

type FontTable struct {
	Sans, Serif, Mono, Script, Symbol string
}

type ColorPalette struct {
	Colors []Color
}

type ColorIndex int

type ForeColor Color

type LineWidth struct {
	Value float64
}

type FontSize struct {
	Percent float64
}

type Viewport struct {
	XMin, XMax, YMin, YMax float64
}

/*
type XYRatio struct {
	Value float64
}
*/

type Scale struct {
	XMin, XMax, YMin, YMax float64
}

type Offset struct {
	Left, Right, Top, Bottom float64
}

type AxisLabel struct {
	X, Y, Title string
}

type LineStyle struct {
	Value int
}

type Points struct {
	X, Y  []float64
	Shape string
}

type Text struct {
	X, Y      float64
	Alignment Alignment
	Value     string
}

type PolyLine struct {
	X, Y []float64
}

type Polygon struct {
	X, Y []float64
}

type Histogram struct {
	X    []float64
	NBin int
}

type DensityHistogram struct {
	X    []float64
	NBin int
}
