// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	gcc "github.com/hailiang/go-gccxml"
	"strings"
)

var (
	apixml = flag.String("apixml", "", "Path of api.xml")
	header = flag.String("header", "", "Path of plplot.h")
)

func main() {
	flag.Parse()
	if *header == "" || *apixml == "" {
		flag.PrintDefaults()
		return
	}

	docFuncs := ParseApiXml(*apixml, "pl")
	h, err := gcc.Xml{*header}.Doc()
	c(err)

	for _, hf := range h.Functions {
		if !strings.HasPrefix(hf.CName(), "c_") {
			continue
		}
		if df := docFuncs.Find(trimPrefix(hf.CName(), "c_")); df != nil {
		} else {
			p(hf.CName(), "is not documented.")
		}
	}

	for _, df := range docFuncs {
		if hf := findHeaderFunc(h, df.Name); hf != nil {
			if len(hf.Arguments) != len(df.Args) {
				p("parameter count mismatch for function", df.Name)
				break
			}
			for i := range df.Args {
				da, ha := df.Args[i], hf.Arguments[i]
				if da.Name != ha.CName() {
					p("number", i, "parameter name mismatch for function",
						df.Name, "(", da.Name, ha.CName(), ")")
				}
			}
		} else {
			p(df.Name, "does not exist but remains in doc.")
		}
	}
}

func findHeaderFunc(h *gcc.XmlDoc, name string) *gcc.Function {
	for _, f := range h.Functions {
		if name == trimPrefix(f.CName(), "c_") {
			return f
		}
	}
	return nil
}

func trimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}
