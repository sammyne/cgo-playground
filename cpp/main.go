package main

//go:generate g++ -o c.o -c c.c
//go:generate ar r libhi.so c.o
//go:generate rm c.o

// #cgo LDFLAGS: -L${SRCDIR} -lhi -lstdc++
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}
// #cgo CFLAGS: -I${SRCDIR}
// #include "c.h"
import "C"

func main() {
	C.hi()
}
