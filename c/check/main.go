// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	gcc "h12.me/go-gccxml"
	"os"
	"strings"
)

var (
	apixml = flag.String("apixml", "", "Path of api.xml")
	header = flag.String("header", "", "Path of plplot.h")
	incdir = flag.String("incdir", "", "colon seperated include directories")
)

func main() {
	flag.Parse()
	if *header == "" || *apixml == "" {
		flag.PrintDefaults()
		return
	}

	var incDirs []string
	if incdir != nil {
		incDirs = strings.Split(*incdir, ":")
	}
	for i := range incDirs {
		incDirs[i] = "-I" + incDirs[i]
	}

	hasMismatch := false

	docFuncs := ParseApiXml(*apixml, "pl")
	h, err := gcc.Xml{*header, incDirs}.Doc()
	c(err)

	for _, hf := range h.Functions {
		if !strings.HasPrefix(hf.CName(), "c_") {
			continue
		}
		if df := docFuncs.Find(trimPrefix(hf.CName(), "c_")); df != nil {
		} else {
			hasMismatch = true
			p(hf.CName(), "is not documented.")
		}
	}

	for _, df := range docFuncs {
		if hf := findHeaderFunc(h, df.Name); hf != nil {
			if len(hf.Arguments) != len(df.Args) {
				hasMismatch = true
				p("para count mismatch for func", df.Name)
				continue
			}
			for i := range df.Args {
				da, ha := df.Args[i], hf.Arguments[i]
				if da.Name != ha.CName() {
					hasMismatch = true
					p("para", i, "of func", df.Name, "name mismatch [",
						da.Name, "] vs. [", ha.CName(), "]")
				}
				if gccType := strings.TrimSpace(ha.CType().CDecl("")); da.Type != gccType {
					hasMismatch = true
					p("para", i, "of func", df.Name, "type mismatch",
						"[", da.Type, "] vs. [", gccType, "]")
				}
			}
		} else {
			hasMismatch = true
			p(df.Name, "does not exist but remains in doc.")
		}
	}

	if hasMismatch {
		os.Exit(-1)
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
