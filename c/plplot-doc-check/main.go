// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"os"
	"strings"

	gcc "h12.io/go-gccxml"
)

var (
	apixml  = flag.String("apixml", "", "Path of api.xml")
	header  = flag.String("header", "", "Path of plplot.h")
	incdir  = flag.String("incdir", "", "colon seperated include directories")
	castxml = flag.Bool("castxml", false, "whether to use castxml instead of gccxml")
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
	h, err := gcc.Xml{File: *header, Args: incDirs, CastXml: *castxml}.Doc()
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
				docArg, headerArg := df.Args[i], hf.Arguments[i]
				if docArg.Name != headerArg.CName() {
					hasMismatch = true
					p("para", i, "of func", df.Name, "NAME mismatch doc[",
						docArg.Name, "] vs. header[", headerArg.CName(), "]")
				}
				gccType := trimArgType(headerArg.CType().CDecl(""))
				docType := trimArgType(docArg.Type)
				if docType != gccType {
					hasMismatch = true
					p("para", i, "of func", df.Name, "TYPE mismatch",
						"doc[", docType, "] vs. header[", gccType, "]")
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

func trimArgType(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "function ")
	s = strings.TrimPrefix(s, "struct ")
	s = strings.TrimSpace(s)
	return s
}
