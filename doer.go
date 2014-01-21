// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plplot

import (
	"reflect"
)

type Doer interface {
	Do()
}

type Doers []Doer

type Preparer interface {
	Prepare() Doer
}

func (ds Doers) Do() {
	for _, d := range ds {
		d.Do()
	}
}

// Take and delete a doer of the same type as default_ from Doers, if not
// exists, just return default_.
func (ds *Doers) Take(default_ Doer) Doer {
	for i, d := range *ds {
		if reflect.TypeOf(d) == reflect.TypeOf(default_) {
			*ds = append((*ds)[:i], (*ds)[i+1:]...)
			return d
		}
	}
	return default_
}

func (ds *Doers) Add(vs ...Doer) {
	for _, v := range vs {
		*ds = append(*ds, v)
	}
}

func (ds Doers) Prepare() Doer {
	for i := range ds {
		if d, ok := ds[i].(Preparer); ok {
			ds[i] = d.Prepare()
		}
	}
	return ds
}

func (ds Doers) FindChild(v Doer) Doer {
	for _, d := range ds {
		if reflect.TypeOf(d) == reflect.TypeOf(v) {
			return d
		}
	}
	return nil
}

func (ds Doers) Find(v Doer) (r Doer) {
	dfs(ds, func(d Doer) bool {
		if reflect.TypeOf(d) == reflect.TypeOf(v) {
			r = d
			return false
		}
		return true
	})
	return
}

/*
func (ds Doers) Select(filter DoerFilter) (rs Doers) {
	for _, d := range ds {
		if filter(d) {
			rs = append(rs, d)
		}
	}
	return
}
*/

/*
func (ds Doers) Get(v Doer) Doer {
	for _, d := range ds {
		if reflect.TypeOf(d) == reflect.TypeOf(v) {
			return d
		}
	}
	return nil
}

// If ds does not contain a v type, add v, otherwise, ignore it.
func (ds Doers) Set(v Doer) Doers {
	if !ds.Has(v) {
		return append(ds, v)
	}
	return ds
}
*/

/*


func (ds *Doers) Set(v Doer) {
	if ds.Has(v) {
		ds.Update(v)
	}
	ds.Add(v)
}

func (ds Doers) Update(v Doer) {
	for i, d := range ds {
		if reflect.TypeOf(d) == reflect.TypeOf(v) {
			ds[i] = v
			return
		}
	}
}
*/

/*
func (ds *Doers) Insert(v Doer, i int) {
	*ds = append(*ds, nil)
	copy((*ds)[i+1:], (*ds)[i:])
	(*ds)[i] = v
}
*/

func minMaxScale(root Doer, offset Offset) Scale {
	xmin, xmax, ymin, ymax := 0.0, 1.0, 0.0, 1.0
	firstX, firstY := true, true
	dfs(root, func(d Doer) bool {
		v := reflect.ValueOf(d)
		if v.Kind() == reflect.Struct {
			xv, yv := v.FieldByName("X"), v.FieldByName("Y")
			if xv.IsValid() && xv.Kind() == reflect.Slice {
				xs := xv.Interface().([]float64)
				for _, x := range xs {
					if firstX {
						xmin, xmax = x, x
						firstX = false
					}
					if xmin > x {
						xmin = x
					}
					if xmax < x {
						xmax = x
					}
				}
			}
			if yv.IsValid() && yv.Kind() == reflect.Slice {
				ys := yv.Interface().([]float64)
				for _, y := range ys {
					if firstY {
						ymin, ymax = y, y
						firstY = false
					}
					if ymin > y {
						ymin = y
					}
					if ymax < y {
						ymax = y
					}
				}
			}
		}
		return true
	})
	xrange, yrange := xmax-xmin, ymax-ymin
	xmin -= xrange * offset.Left
	xmax += xrange * offset.Right
	ymin -= yrange * offset.Bottom
	ymax += yrange * offset.Top
	return Scale{XMin: xmin, XMax: xmax, YMin: ymin, YMax: ymax}
}

func dfs(d Doer, visit func(v Doer) bool) {
	if !visit(d) {
		return
	}
	v := reflect.ValueOf(d)
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			dfs(v.Index(i).Interface().(Doer), visit)
		}
	}
}

type Do struct {
	F func() `xml:"-"`
}

func (d Do) Do() {
	d.F()
}
