// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c

/*
#include <stdlib.h>
*/
import "C"

import (
    "fmt"
    "unsafe"
)

type P unsafe.Pointer

func S(s string) *C.char {
    return C.CString(s)
}

func freeS(s *C.char) {
    C.free(unsafe.Pointer(s))
}

func fromDoubleArray(p *C.double, i int) float64 {
    return float64(*(*C.double)(P(uintptr(P(p)) + unsafe.Sizeof(p)*uintptr(i))))
}

func fromIntArray(p *C.int, i int) int {
    return int(*(*C.int)(P(uintptr(P(p)) + unsafe.Sizeof(p)*uintptr(i))))
}

func p(v ...interface{}) {
    fmt.Println(v...)
}
