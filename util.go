// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"os/exec"
	"strings"
)

type UnaryFunc func(float64) float64

func (f UnaryFunc) sample() (x, y []float64) {
	nx, _ := SampleSize()
	xstep, _ := SampleStep()
	x, y = make([]float64, nx), make([]float64, nx)
	for i := 0; i < nx; i++ {
		x[i] = xstep * float64(i)
		y[i] = f(x[i])
	}
	return
}

type BinaryFunc func(float64, float64) float64

func (f BinaryFunc) sample() *BinaryFuncSample {
	vp := gs.Viewport()
	nx, ny := SampleSize()
	xstep, ystep := SampleStep()
	v := make([][]float64, nx)
	for i := 0; i < int(nx); i++ {
		x := vp.XMin + xstep*float64(i)
		v[i] = make([]float64, ny)
		for j := 0; j < int(ny); j++ {
			y := vp.YMin + ystep*float64(j)
			v[i][j] = f(float64(x), float64(y))
		}
	}
	return &BinaryFuncSample{v, xstep, ystep}
}

type BinaryFuncSample struct {
	V            [][]float64
	Xstep, Ystep float64
}

func (s *BinaryFuncSample) min() float64 {
	zmin := math.MaxFloat64
	for _, r := range s.V {
		for _, c := range r {
			if c < zmin {
				zmin = c
			}
		}
	}
	return zmin
}

func (s *BinaryFuncSample) max() float64 {
	zmax := -math.MaxFloat64
	for _, r := range s.V {
		for _, c := range r {
			if c > zmax {
				zmax = c
			}
		}
	}
	return zmax
}

func minMax(data []float64) (float64, float64) {
	min, max := math.MaxFloat64, -math.MaxFloat64
	for _, v := range data {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func p(v ...interface{}) {
	fmt.Println(v...)
}

func pj(v interface{}) {
	buf, err := json.MarshalIndent(v, "", "    ")
	ce(err)
	p(string(buf))
}

func ce(err error) {
	if err != nil {
		panic(err)
	}
}

func px(v interface{}) {
	buf, err := xml.MarshalIndent(v, "", "    ")
	ce(err)
	fmt.Println(string(buf))
}

func pc(v interface{}) {
	s := fmt.Sprintf("%#v", v)
	code := ""
	indent := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i : i+1]; c {
		case "{":
			indent++
			code += c + "\n" + strings.Repeat(" ", indent*4)
		case "}":
			indent--
			code += c
			if i < len(s)-1 && s[i+1] == ',' {
				code += ","
				i++
			}
			for i < len(s)-1 && s[i+1] == ' ' {
				code += " "
				i++
			}
			code += "\n" + strings.Repeat(" ", indent*4)
		default:
			code += string(c)
		}
	}
	fmt.Print(code)
}

func format(file string) {
	cmd := exec.Command("go", "fmt", file)
	err := cmd.Start()
	ce(err)
	err = cmd.Wait()
	ce(err)
}

func fp(w io.Writer, v ...interface{}) {
	fmt.Fprint(w, v...)
	fmt.Fprintln(w)
}

func round(f float64) int {
	return int(f + 0.5)
}

