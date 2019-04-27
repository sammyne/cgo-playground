package main

import "fmt"

// the -rpath option set the path to load dynamic libraries
// when running the program

// #cgo CFLAGS: -I./
// #cgo LDFLAGS: -L${SRCDIR}/ -lhi -Wl,-rpath=${SRCDIR}
// #include "hi.h" //非标准c头文件，所以用引号
import "C"

func main() {
	C.hi()
	fmt.Println("Hi, vim-go")
}
