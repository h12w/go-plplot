// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_quick(t *testing.T) {
	LogOn()
	unit := 32
	plot := Plot{
		File{"out.pdf"},
		PageConfig{
			Width:     unit * 4,
			Height:    unit * 3,
			NSubpageX: 1,
			NSubpageY: 2,
		},
		LineWidth{0.2},
		Subpage{
			//Viewport{XMin: 0, XMax: 0.5, YMin: 0, YMax: 1},
			Block{
				PolyLine{[]float64{0, 3}, []float64{0, 3}},
			},
			Axis{Position: TOP | RIGHT | LEFT | BOTTOM},
			Legend{
				Option:      LEGEND_BACKGROUND | LEGEND_BOUNDING_BOX,
				Position:    TOP | RIGHT | INSIDE | VIEWPORT,
				Offset:      LegendOffset{0.02, 0.02},
				PlotWidth:   0.05,
				BackColor:   0,
				Box:         BoundingBox{1, 1},
				OptionArray: []LegendType{LEGEND_LINE, LEGEND_LINE, LEGEND_LINE, LEGEND_LINE},
				Text:        TextLegend{Offset: 0.5, Scale: 0.8, Spacing: 1.5, Justification: 1.0, Colors: []int{1, 1, 1, 1}, Texts: []string{"k=1", "k=2", "k=3", "k=4"}},
				Line:        LineLegend{Colors: []int{3, 5, 8, 11}, Styles: []int{1, 1, 1, 1}, Widths: []float64{0.2, 0.2, 0.2, 0.2}},
			},
		},
		Subpage{},
	}
	pp := spew.ConfigState{Indent: "    "}

	prepared := plot.Prepare()
	pp.Dump(prepared)
	prepared.Do()
}
