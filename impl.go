// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

import (
	"github.com/hailiang/go-plplot/c"
	"path/filepath"
)

/*
The following routines must be called before plinit:
setopt
ssub
scolbg
sdev
sfnam
spage
sfile
fontld*
*/

func (plot Plot) Render() {
	plot.Prepare().Do()
}

func (plot Plot) Do() {
	// special initialization for timeformat,
	// may need a better place to put the code
	c.Timefmt(DefaultTimeFormat)

	Doers(plot).Do()
}

func (plot Plot) Prepare() Doer {
	ds := Doers(plot)

	// set default
	init, rest := Doers{}, Doers{}
	init.Add(ds.Take(DefaultFile))
	init.Add(ds.Take(DefaultPageConfig))
	init.Add(ds.Take(DefaultBackColor))
	init.Add(ds.Take(DefaultFontTable))
	init.Add(Init{})

	styles := Doers{}
	for _, d := range ds {
		if DefaultStyles.FindChild(d) != nil {
			styles.Add(d)
		} else {
			break
		}
	}
	ds = ds[len(styles):]
	rest.Add(styles.Take(DefaultForeColor))
	rest.Add(styles.Take(DefaultLineWidth))
	rest.Add(styles.Take(DefaultLineStyle))
	rest.Add(styles.Take(DefaultFontSize))
	if ds.FindChild(Subpage{}) != nil {
		rest.Add(ds...)
	} else {
		rest.Add(Subpage(ds))
	}
	rest.Add(End{})

	ds = append(init, rest...)
	ds.Prepare()
	return Plot(ds)
}

func (Init) Do() {
	c.Init()
}

func (End) Do() {
	c.End()
}

var driverMap = map[string]string{
	".pdf": "pdfcairo",
	".png": "pngcairo",
	".ps":  "pscairo",
	".svg": "svgcairo",
	".eps": "epscairo",
}

func (f File) Do() {
	file := f.Path
	ext := filepath.Ext(file)

	if driver, ok := driverMap[ext]; ok {
		c.Sdev(driver)
		c.Sfnam(file)
	} else {
		panic("driver not supported.")
	}
}

func (p PageConfig) Do() {
	c.Spage(DPI, DPI,
		p.Width, p.Height, p.Left, p.Top)
	if p.NSubpageX > 0 {
		c.Ssub(p.NSubpageX, p.NSubpageY)
	}
}

func (p PageConfig) SubpageWidth() float64 {
	return float64(p.Width) / float64(p.NSubpageX)
}

func (p PageConfig) SubpageHeight() float64 {
	return float64(p.Height) / float64(p.NSubpageY)
}

func (color BackColor) Do() {
	c.Scolbg(int(color.R), int(color.G), int(color.B))
}

func setFont(family, name string) {
	if name != "" {
		c.Setenv(family, name)
	}
}

func (t FontTable) Do() {
	setFont("PLPLOT_FREETYPE_SANS_FAMILY", t.Sans)
	setFont("PLPLOT_FREETYPE_SERIF_FAMILY", t.Serif)
	setFont("PLPLOT_FREETYPE_MONO_FAMILY", t.Mono)
	setFont("PLPLOT_FREETYPE_SCRIPT_FAMILY", t.Script)
	setFont("PLPLOT_FREETYPE_SYMBOL_FAMILY", t.Symbol)
}

func (p ColorPalette) Do() {
	for i, color := range p.Colors {
		c.Scol0(i, int(color.R), int(color.G), int(color.B))
	}
}

func (i ColorIndex) Do() {
	c.Col0(int(i))
}

func (color ForeColor) Do() {
	c.Scol0(1, int(color.R), int(color.G), int(color.B))
	c.Col0(1)
}

func (w LineWidth) Do() {
	c.Width(w.Value)
}

func (s FontSize) Do() {
	c.Schr(0, s.Percent)
}

func (p Subpage) Do() {
	c.Adv(0)
	Doers(p).Do()
}

func (p Subpage) Prepare() Doer {
	ds := Doers(p)
	var v Doer
	if ds.Find(Axis{}) != nil {
		v = ds.Take(DefaultViewport)
	} else {
		v = ds.Take(DefaultViewportNoAxis)
	}
	//x := ds.Take(DefaultXYRatio)
	o := ds.Take(DefaultOffset).(Offset)
	s := ds.Take(minMaxScale(p, o))
	ds = append(Doers{v, o, s}, ds...)
	ds.Prepare()
	return Subpage{Block(ds)}
}

func (b Block) Prepare() Doer {
	Doers(b).Prepare()
	return b
}

func (v Viewport) Do() {
	c.Vpor(v.XMin, v.XMax, v.YMin, v.YMax)
}

/*
func (r XYRatio) Do() {
	// automatic ratio to fill the page
	if r.Value == 0 {
		pageConfig := gs.PageConfig()
		r.Value = pageConfig.SubpageWidth() / pageConfig.SubpageHeight()
	}
	c.Vpor(0, 1, 0, r.Value)
}
*/

func (r Offset) Do() {
}

func (r Scale) Do() {
	c.Wind(r.XMin, r.XMax, r.YMin, r.YMax)
}

func (l AxisLabel) Do() {
	c.Lab(l.X, l.Y, l.Title)
}

func (s LineStyle) Do() {
	c.Lsty(s.Value)
}

func Font(ff FontFamily, fs FontStyle, fw FontWeight) {
	fci := 0x80000000 | uint32(ff) | uint32(fs)<<4 | uint32(fw)<<8
	c.Sfci(int(fci))
}

func (p Points) Do() {
	c.String(len(p.X), p.X, p.Y, p.Shape)
}

/*
func Point(x, y float64, shape int) {
	switch shape {
	case LARGE_POINT:
		Circle(x, y, 0.05, true)
	case LARGE_CIRCLE:
		Circle(x, y, 0.05, false)
	default:
		c.Poin(1, []float64{x}, []float64{y}, shape)
	}
}
*/

func (t Text) Do() {
	c.Ptex(t.X, t.Y, 0, 0, float64(t.Alignment), t.Value)
}

/*
func Line(x0, y0, x1, y1 float64) {
	if x0 == x1 && y0 == y1 {
		panic("Two point equals.")
	} else if x0 == x1 {
		Segment(x0, gs.YMin, x0, gs.YMax)
	} else if y0 == y1 {
		Segment(gs.XMin, y0, gs.XMax, y0)
	} else {
		// y - y0 = m(x-x0)
		m := (y1 - y0) / (x1 - x0)
		sy, ey := m*(gs.XMin-x0)+y0, m*(gs.XMax-x0)+y0
		Segment(gs.XMin, sy, gs.XMax, ey)
	}
}
*/

func (l PolyLine) Do() {
	if len(l.X) != len(l.Y) {
		panic("X, Y length mismatch.")
	}
	c.Line(len(l.X), l.X, l.Y)
}

func TimeFormat(format string) {
	c.Timefmt(format)
}

func (p Polygon) Do() {
	c.Fill(len(p.X), p.X, p.Y)
}

func Circle(x, y, r float64, fill bool) {
	c.Arc(x, y, r, r, 0, 360, 0, fill)
}

// Plot a unary function
func Function(f UnaryFunc) {
	x, y := f.sample()
	c.Line(len(x), x, y)
}

// Contour plot
func Contour(f BinaryFunc, nlevel int) {
	s := f.sample()
	v := s.V
	nx, ny := len(v), len(v[0])
	kx, ky := 1, 1
	lx, ly := nx, ny
	clevels := calcLevels(s.min(), s.max(), nlevel)
	vp := gs.Viewport()
	trans := s.trans(vp.XMin, vp.YMin)
	c.Cont(v, nx, ny, kx, lx, ky, ly, clevels, len(clevels), trans)
}

func (s *BinaryFuncSample) trans(xmin, ymin float64) c.ContourTransform {
	return func(x, y float64) (float64, float64) {
		return xmin + s.Xstep*x, ymin + s.Ystep*y
	}
}

func calcLevels(zmin, zmax float64, nlevel int) []float64 {
	zrange := zmax - zmin
	levels := make([]float64, nlevel)
	levelstep := zrange / float64(nlevel-1)
	for i := range levels {
		levels[i] = zmin + levelstep*float64(i)
	}

	// smooth the edge
	if len(levels) >= 2 {
		levels[0] += zrange * ContourEdgeSmooth
		levels[len(levels)-1] -= zrange * ContourEdgeSmooth
	}
	return levels
}

func (h Histogram) Do() {
	datmin, datmax := minMax(h.X)
	c.Hist(len(h.X), h.X, datmin, datmax, h.NBin, 0)
}

/*
// Histogram
func Histogram(data []float64, nbin int) {
	datmin, datmax := minMax(data)
	c.Hist(len(data), data, datmin, datmax, nbin, 0)
}
*/

func (h DensityHistogram) Do() {
	min, max := minMax(h.X)
	binWidth := (max - min) / float64(h.NBin)
	x, y := make([]float64, h.NBin), make([]float64, h.NBin)
	for i := range x {
		x[i] = min + binWidth*float64(i)
	}
	for _, v := range h.X {
		ibin := int((v - min) / binWidth)
		if ibin == len(y) {
			ibin--
		}
		y[ibin]++
	}
	for i := range y {
		y[i] /= float64(len(h.X)) * binWidth
	}
	c.Bin(len(x), x, y, 0)
}

/*
// Density histogram
func DensityHistogram(data []float64, nbin int) {
	min, max := minMax(data)
	binWidth := (max - min) / float64(nbin)
	x, y := make([]float64, nbin), make([]float64, nbin)
	for i := range x {
		x[i] = min + binWidth*float64(i)
	}
	for _, v := range data {
		ibin := int((v - min) / binWidth)
		if ibin == len(y) {
			ibin--
		}
		y[ibin]++
	}
	for i := range y {
		y[i] /= float64(len(data)) * binWidth
	}
	c.Bin(len(x), x, y, 0)
}
*/

func LogOn() {
	c.Logging = true
}

func LogOff() {
	c.Logging = false
}

var SamplePrecision = float64(0.3)
var ContourEdgeSmooth = float64(0.01)

func SampleSize() (nx, ny int) {
	page := gs.PageConfig()
	return int(float64(page.Width) * SamplePrecision), int(float64(page.Height) * SamplePrecision)
}

func SampleStep() (xstep, ystep float64) {
	nx, ny := SampleSize()
	vp := gs.Viewport()
	xrange, yrange := vp.XMax-vp.XMin, vp.YMax-vp.YMin
	return xrange / float64(nx-1), yrange / float64(ny-1)
}
