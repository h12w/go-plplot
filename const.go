// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignRight
	AlignCenter
)

const (
	DEFAULT_LINE = 1
	DASH_LINE    = 2
)

const (
	LARGE_CIRCLE = -3
	LARGE_POINT  = -2
)

const (
	PT_SQUARE        = iota // '□' +U25A1
	PT_1                    //
	PT_PLUS                 // '+' +U002B
	PT_ASTERISK             // '∗' +U2217
	PT_CIRCLE               // '○' +U25CB
	PT_MULT                 // '×' +U00D7
	PT_6                    // '□' +U25A1
	PT_TRIANGLE             // '△' +U25B3
	PT_EARTH                // '♁' +U2641
	PT_CIRCLED_DOT          // '⊙' +U2299
	PT_10                   //
	PT_DIAMOND              // '♢' +U2662
	PT_STAR                 // '✩' +U2729
	PT_13                   // '□' +U25A1
	PT_14                   //
	PT_STAR_OF_DAVID        // '✡' +U2721
	PT_SOLID_SQUARE         // '■' +U25A0
	PT_BULLET               // '∙' +U2219
	PT_SOLID_STAR           // '⋆' +U22C6
	PT_19                   // '□' +U25A1
	PT_20                   //
	PT_21                   //
	PT_22                   //
	PT_23                   //
	PT_24                   //
	PT_25                   //
	PT_26                   //
	PT_27                   //
	PT_LEFT_ARROW           // '←' +U2190
	PT_RIGHT_ARROW          // '→' +U2192
	PT_UP_ARROW             // '↑' +U2191
	PT_DOWN_ARROW           // '↓' +U2193
)

// const (
//     SCALE_NONE    = -1
//     SCALE_STRETCH = 0
//     SCALE_EQUAL   = 1
//     SCALE_SQUARE  = 2
//     AXIS_NONE      = -2
//     AXIS_BOXONLY   = -1
//     AXIS_DEFAULT   = 0
//     AXIS_ORIGIN    = 1
//     AXIS_MAJORGRID = 2
//     AXIS_MINORGRID = 3
//     AXIS_LOG_XY    = 30
// )

type (
	FontFamily uint32
	FontStyle  uint32
	FontWeight uint32
)

const (
	SANS FontFamily = iota
	SERIF
	MONO
	SCRIPT
	SYMBOL
)

const (
	UPRIGHT FontStyle = iota
	ITALIC
	OBLIQUE
)

const (
	MEDIUM FontWeight = iota
	BOLD
)

