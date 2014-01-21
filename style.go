// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

var DefaultStyles = Doers{
	DefaultForeColor,
	DefaultLineWidth,
	DefaultLineStyle,
	DefaultFontSize,
}

type styleValue struct {
	ForeColor
	LineWidth
	LineStyle
	FontSize
}

func loadCurrentStyle() styleValue {
	return styleValue{
		ForeColor: gs.ForeColor(),
		LineWidth: gs.LineWidth(),
		LineStyle: gs.LineStyle(),
		FontSize:  gs.FontSize(),
	}
}

func restore(old styleValue) {
	cur := loadCurrentStyle()
	if old.ForeColor != cur.ForeColor {
		old.ForeColor.Do()
	}
	if old.LineWidth != cur.LineWidth {
		old.LineWidth.Do()
	}
	if old.LineStyle != cur.LineStyle {
		old.LineStyle.Do()
	}
	if old.FontSize != cur.FontSize {
		old.FontSize.Do()
	}
}

var styleStack []styleValue

func (block Block) Do() {
	styleStack = append(styleStack, loadCurrentStyle())
	Doers(block).Do()
	var old styleValue
	styleStack, old = styleStack[:len(styleStack)-1],
		styleStack[len(styleStack)-1]
	restore(old)
}
