// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

import (
	"github.com/hailiang/go-plplot/c"
)

var gs globalState

type globalState struct {
}

func (globalState) PageConfig() PageConfig {
	_, _, width, height, left, top := c.Gpage()
	nSubpageX, nSubpageY := c.Gsub()
	return PageConfig{
		Width:     int(width),
		Height:    int(height),
		Left:      int(left),
		Top:       int(top),
		NSubpageX: int(nSubpageX),
		NSubpageY: int(nSubpageY)}
}

func (globalState) Viewport() Viewport {
	xmin, xmax, ymin, ymax := c.Gvpw()
	return Viewport{
		XMin: xmin,
		XMax: xmax,
		YMin: ymin,
		YMax: ymax,
	}
}

func (globalState) ForeColor() ForeColor {
	r, g, b := c.Gcol0(1)
	return ForeColor{byte(r), byte(g), byte(b)}
}

func (globalState) LineWidth() LineWidth {
	return LineWidth{c.Gwidth()}
}

func (globalState) LineStyle() LineStyle {
	return LineStyle{int(c.Glsty())}
}

func (globalState) FontSize() FontSize {
	// value got from gchr is different from the values set via schr
	def, scaledValue := c.Gchr()
	return FontSize{scaledValue/def}
}
